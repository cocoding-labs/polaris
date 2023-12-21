package jsonstore

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"math/big"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/valyala/fastjson"
	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/jsonstore"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/eth/crypto"
)

// Contract is the main struct for the HypNative contract
type Contract struct {
	ethprecompile.BaseContract
}

// NewPrecompileContract creates a new instance of the HypNative contract
func NewPrecompileContract() *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.JsonStoreMetaData.ABI,
			common.HexToAddress("0x666F726d61000000000000000000000000000002"),
		),
	}
}

// Solidity: function get(address key, uint256 slot, string path) view returns(string)
func (c *Contract) Get(ctx context.Context, key common.Address, slot *big.Int, path string) (string, error) {
	storageVal := retrieveStorage(ctx, key, slot)
	filteredVal := gjson.Get(storageVal, path)
	return filteredVal.String(), nil
}

// Solidity: function get(address key, uint256 slot) view returns(string)
func (c *Contract) Get0(ctx context.Context, key common.Address, slot *big.Int) (string, error) {
	storageVal := retrieveStorage(ctx, key, slot)
	if storageVal == "" {
		storageVal = `{}`
	}
	return storageVal, nil
}

// Solidity: function dataURI(address key, uint256 slot) view returns(string)
func (c *Contract) DataURI(ctx context.Context, key common.Address, slot *big.Int) (string, error) {
	storageVal := retrieveStorage(ctx, key, slot)
	if storageVal == "" {
		storageVal = `{}`
	}
	dataURI := "data:application/json;base64," + base64.StdEncoding.EncodeToString([]byte(storageVal))
	return dataURI, nil
}

// Solidity: function set(uint256 slot, string jsonBlob) returns(bool)
func (c *Contract) Set(ctx context.Context, slot *big.Int, jsonBlob string) (bool, error) {
	pCtx := vm.UnwrapPolarContext(ctx)
	key := pCtx.MsgSender()
	return updateStorage(ctx, key, slot, jsonBlob)
}

// Solidity: function set(uint256 slot, string path, string value) returns(bool)
func (c *Contract) Set0(ctx context.Context, slot *big.Int, path string, value string) (bool, error) {
	return c.Set1(ctx, slot, []string{path}, []string{value})
}

// Solidity: function set(uint256 slot, string[] path, string[] value) returns(bool)
func (c *Contract) Set1(ctx context.Context, slot *big.Int, paths []string, values []string) (bool, error) {
	pCtx := vm.UnwrapPolarContext(ctx)
	key := pCtx.MsgSender()
	storageVal := retrieveStorage(ctx, key, slot)
	if storageVal == "" {
		storageVal = `{}`
	}
	updatedJson := storageVal
	for i, path := range paths {
		updatedJson, _ = sjson.Set(updatedJson, path, values[i])
	}
	return updateStorage(ctx, key, slot, updatedJson)
}

// Solidity: function subReplace(uint256 slot, string searchPath, string[] replacePaths, string[] values) returns(bool)
func (c *Contract) SubReplace(ctx context.Context, slot *big.Int, searchPath string, replacePaths []string, values []string) (bool, error) {
	pCtx := vm.UnwrapPolarContext(ctx)
	key := pCtx.MsgSender()
	storageVal := retrieveStorage(ctx, key, slot)
	result := gjson.Get(storageVal, searchPath)
	updatedJson := storageVal
	offset := 0
	for _, r := range result.Array() {
		replacement := r.Raw
		for i, replacePath := range replacePaths {
			replacement, _ = sjson.Set(replacement, replacePath, values[i])
		}
		updatedJson = updatedJson[:r.Index+offset] + replacement + updatedJson[r.Index+offset+len(r.Raw):]
		offset += len(replacement) - len(r.Raw)
	}
	return updateStorage(ctx, key, slot, updatedJson)
}

// Solidity: function subReplace(uint256 slot, string searchPath, string replacePath, string value) returns(bool)
func (c *Contract) SubReplace0(ctx context.Context, slot *big.Int, searchPath string, replacePath string, value string) (bool, error) {
	return c.SubReplace(ctx, slot, searchPath, []string{replacePath}, []string{value})
}

// Solidity: function remove(uint256 slot, string path) returns(bool)
func (c *Contract) Remove(ctx context.Context, slot *big.Int, path string) (bool, error) {
	pCtx := vm.UnwrapPolarContext(ctx)
	key := pCtx.MsgSender()
	storageVal := retrieveStorage(ctx, key, slot)
	if storageVal == "" {
		storageVal = `{}`
	}
	updatedJson, _ := sjson.Delete(storageVal, path)
	return updateStorage(ctx, key, slot, updatedJson)
}

func retrieveStorage(ctx context.Context, key common.Address, slot *big.Int) string {
	pCtx := vm.UnwrapPolarContext(ctx)
	storageKey := crypto.Keccak256Hash(
		common.BytesToHash(key.Bytes()).Bytes(),
		common.BigToHash(slot).Bytes(),
	)
	storageVal := pCtx.GetState(storageKey).Bytes()

	if storageVal[0] != 0 {
		length := int(storageVal[len(storageVal)-1] / 2)
		res, _ := uncompress(storageVal[:length])
		return string(res)
	} else {
		length := int(new(big.Int).SetBytes(storageVal).Int64() / 2)
		res := make([]byte, length)
		chunkKeyInt := crypto.Keccak256Hash(storageKey.Bytes()).Big()
		for i := 0; i < length; i += 32 {
			end := i + 32
			if end > length {
				end = length
			}

			chunkKey := common.BigToHash(chunkKeyInt)
			chunkVal := pCtx.GetState(chunkKey).Bytes()
			copy(res[i:end], chunkVal[:])
			chunkKeyInt.Add(chunkKeyInt, big.NewInt(1))
		}
		res, _ = uncompress(res)
		return string(res)
	}
}

func updateStorage(ctx context.Context, key common.Address, slot *big.Int, jsonBlob string) (bool, error) {
	validationErr := fastjson.Validate(jsonBlob)
	if validationErr != nil {
		return false, errors.New("invalid JSON")
	}

	pCtx := vm.UnwrapPolarContext(ctx)
	storageKey := crypto.Keccak256Hash(
		common.BytesToHash(key.Bytes()).Bytes(),
		common.BigToHash(slot).Bytes(),
	)

	// compact json
	var buf bytes.Buffer
	err := json.Compact(&buf, []byte(jsonBlob))
	if err != nil {
		return false, errors.New("unknown compact error")
	}
	jsonBytes := buf.Bytes()

	// compress for storage
	jsonBytes, _ = compress(jsonBytes)

	var storageVal common.Hash
	length := len(jsonBytes)
	if length < 32 {
		store := make([]byte, 32)
		copy(store[0:length], jsonBytes[:])
		store[31] = byte(length * 2)
		storageVal = common.BytesToHash(store)
	} else {
		storageVal = common.BigToHash(big.NewInt(int64(length * 2)))
		chunkKeyInt := crypto.Keccak256Hash(storageKey.Bytes()).Big()
		for i := 0; i < length; i += 32 {
			end := i + 32
			if end > length {
				end = length
			}

			chunk := make([]byte, 32)
			copy(chunk, jsonBytes[i:end])

			chunkKey := common.BigToHash(chunkKeyInt)
			pCtx.SetState(chunkKey, common.BytesToHash(chunk))
			chunkKeyInt.Add(chunkKeyInt, big.NewInt(1))
		}
	}

	pCtx.SetState(storageKey, storageVal)

	return true, nil
}

func compress(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err := gz.Write(input)
	if err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func uncompress(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz, err := gzip.NewReader(bytes.NewBuffer(input))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(&buf, gz)
	if err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
