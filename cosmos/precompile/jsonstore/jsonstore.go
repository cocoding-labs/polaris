package jsonstore

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"io"
	"math/big"

	generated "github.com/berachain/polaris/contracts/bindings/cosmos/precompile/jsonstore"
	"github.com/berachain/polaris/cosmos/precompile/jsonutil"
	ethprecompile "github.com/berachain/polaris/eth/core/precompile"
	pvm "github.com/berachain/polaris/eth/core/vm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/valyala/fastjson"
)

// Contract is the main struct for the HypNative contract
type Contract struct {
	ethprecompile.BaseContract
	jsonutil *jsonutil.Contract
}

// NewPrecompileContract creates a new instance of the HypNative contract
func NewPrecompileContract() *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.JsonStoreMetaData.ABI,
			common.HexToAddress("0x666F726d61000000000000000000000000000002"),
		),
		jsonutil: jsonutil.NewPrecompileContract(),
	}
}

// Solidity: function exists(address key, uint256 slot, string path) view returns(bool)
func (c *Contract) Exists(ctx context.Context, key common.Address, slot *big.Int, path string) (bool, error) {
	storageVal := retrieveStorage(ctx, key, slot)
	return c.jsonutil.Exists(ctx, storageVal, path)
}

// Solidity: function get(address key, uint256 slot, string path) view returns(string)
func (c *Contract) Get(ctx context.Context, key common.Address, slot *big.Int, path string) (string, error) {
	storageVal := retrieveStorage(ctx, key, slot)
	return c.jsonutil.Get(ctx, storageVal, path)
}

// Solidity: function get(address key, uint256 slot) view returns(string)
func (c *Contract) Get0(ctx context.Context, key common.Address, slot *big.Int) (string, error) {
	return retrieveStorage(ctx, key, slot), nil
}

// Solidity: function getBool(address key, uint256 slot, string path) view returns(bool)
func (c *Contract) GetBool(ctx context.Context, key common.Address, slot *big.Int, path string) (bool, error) {
	storageVal := retrieveStorage(ctx, key, slot)
	return c.jsonutil.GetBool(ctx, storageVal, path)
}

// Solidity: function getInt(address key, uint256 slot, string path) view returns(int256)
func (c *Contract) GetInt(ctx context.Context, key common.Address, slot *big.Int, path string) (*big.Int, error) {
	storageVal := retrieveStorage(ctx, key, slot)
	return c.jsonutil.GetInt(ctx, storageVal, path)
}

// Solidity: function getUint(address key, uint256 slot, string path) view returns(uint256)
func (c *Contract) GetUint(ctx context.Context, key common.Address, slot *big.Int, path string) (*big.Int, error) {
	return c.GetInt(ctx, key, slot, path)
}

// Solidity: function dataURI(address key, uint256 slot) view returns(string)
func (c *Contract) DataURI(ctx context.Context, key common.Address, slot *big.Int) (string, error) {
	storageVal := retrieveStorage(ctx, key, slot)
	return c.jsonutil.DataURI(ctx, storageVal)
}

// Solidity: function set(uint256 slot, string jsonBlob) returns(bool)
func (c *Contract) Set(ctx context.Context, slot *big.Int, jsonBlob string) (bool, error) {
	return updateMsgSenderStorage(ctx, slot, jsonBlob)
}

// Solidity: function set(uint256 slot, string path, string value) returns(bool)
func (c *Contract) Set0(ctx context.Context, slot *big.Int, path string, value string) (bool, error) {
	return c.Set1(ctx, slot, []string{path}, []string{value})
}

// Solidity: function set(uint256 slot, string[] paths, string[] values) returns(bool)
func (c *Contract) Set1(ctx context.Context, slot *big.Int, paths []string, values []string) (bool, error) {
	storageVal := retrieveMsgSenderStorage(ctx, slot)
	updatedJson, err := c.jsonutil.Set(ctx, storageVal, paths, values)
	if err != nil {
		return false, err
	}
	return updateMsgSenderStorage(ctx, slot, updatedJson)
}

// Solidity: function setRaw(uint256 slot, string path, string rawBlob) returns(bool)
func (c *Contract) SetRaw(ctx context.Context, slot *big.Int, path string, rawBlob string) (bool, error) {
	return c.SetRaw0(ctx, slot, []string{path}, []string{rawBlob})
}

// Solidity: function setRaw(uint256 slot, string[] paths, string[] rawBlobs) returns(bool)
func (c *Contract) SetRaw0(ctx context.Context, slot *big.Int, paths []string, rawBlobs []string) (bool, error) {
	storageVal := retrieveMsgSenderStorage(ctx, slot)
	updatedJson, err := c.jsonutil.SetRaw0(ctx, storageVal, paths, rawBlobs)
	if err != nil {
		return false, err
	}
	return updateMsgSenderStorage(ctx, slot, updatedJson)
}

// Solidity: function setBool(uint256 slot, string path, bool value) returns(bool)
func (c *Contract) SetBool(ctx context.Context, slot *big.Int, path string, value bool) (bool, error) {
	return c.SetBool0(ctx, slot, []string{path}, []bool{value})
}

// Solidity: function setBool(uint256 slot, string[] paths, bool[] values) returns(bool)
func (c *Contract) SetBool0(ctx context.Context, slot *big.Int, paths []string, values []bool) (bool, error) {
	storageVal := retrieveMsgSenderStorage(ctx, slot)
	updatedJson, err := c.jsonutil.SetBool(ctx, storageVal, paths, values)
	if err != nil {
		return false, err
	}
	return updateMsgSenderStorage(ctx, slot, updatedJson)
}

// Solidity: function setInt(uint256 slot, string[] paths, int256[] values) returns(bool)
func (c *Contract) SetInt(ctx context.Context, slot *big.Int, paths []string, values []*big.Int) (bool, error) {
	storageVal := retrieveMsgSenderStorage(ctx, slot)
	updatedJson, err := c.jsonutil.SetInt(ctx, storageVal, paths, values)
	if err != nil {
		return false, err
	}
	return updateMsgSenderStorage(ctx, slot, updatedJson)
}

// Solidity: function setInt(uint256 slot, string path, int256 value) returns(bool)
func (c *Contract) SetInt0(ctx context.Context, slot *big.Int, path string, value *big.Int) (bool, error) {
	return c.SetInt(ctx, slot, []string{path}, []*big.Int{value})
}

// Solidity: function setUint(uint256 slot, string[] paths, uint256[] values) returns(bool)
func (c *Contract) SetUint(ctx context.Context, slot *big.Int, paths []string, values []*big.Int) (bool, error) {
	return c.SetInt(ctx, slot, paths, values)
}

// Solidity: function setUint(uint256 slot, string path, uint256 value) returns(bool)
func (c *Contract) SetUint0(ctx context.Context, slot *big.Int, path string, value *big.Int) (bool, error) {
	return c.SetInt0(ctx, slot, path, value)
}

// Solidity: function subReplace(uint256 slot, string searchPath, string[] replacePaths, string[] values) returns(bool)
func (c *Contract) SubReplace(ctx context.Context, slot *big.Int, searchPath string, replacePaths []string, values []string) (bool, error) {
	storageVal := retrieveMsgSenderStorage(ctx, slot)
	updatedJson, err := c.jsonutil.SubReplace(ctx, storageVal, searchPath, replacePaths, values)
	if err != nil {
		return false, err
	}
	return updateMsgSenderStorage(ctx, slot, updatedJson)
}

// Solidity: function subReplace(uint256 slot, string searchPath, string replacePath, string value) returns(bool)
func (c *Contract) SubReplace0(ctx context.Context, slot *big.Int, searchPath string, replacePath string, value string) (bool, error) {
	return c.SubReplace(ctx, slot, searchPath, []string{replacePath}, []string{value})
}

// Solidity: function subReplaceBool(uint256 slot, string searchPath, string replacePath, bool value) returns(bool)
func (c *Contract) SubReplaceBool(ctx context.Context, slot *big.Int, searchPath string, replacePath string, value bool) (bool, error) {
	return c.SubReplaceBool0(ctx, slot, searchPath, []string{replacePath}, []bool{value})
}

// Solidity: function subReplaceBool(uint256 slot, string searchPath, string[] replacePaths, bool[] values) returns(bool)
func (c *Contract) SubReplaceBool0(ctx context.Context, slot *big.Int, searchPath string, replacePaths []string, values []bool) (bool, error) {
	storageVal := retrieveMsgSenderStorage(ctx, slot)
	updatedJson, err := c.jsonutil.SubReplaceBool(ctx, storageVal, searchPath, replacePaths, values)
	if err != nil {
		return false, err
	}
	return updateMsgSenderStorage(ctx, slot, updatedJson)
}

// Solidity: function subReplaceInt(uint256 slot, string searchPath, string replacePath, int256 value) returns(bool)
func (c *Contract) SubReplaceInt(ctx context.Context, slot *big.Int, searchPath string, replacePath string, value *big.Int) (bool, error) {
	return c.SubReplaceInt0(ctx, slot, searchPath, []string{replacePath}, []*big.Int{value})
}

// Solidity: function subReplaceInt(uint256 slot, string searchPath, string[] replacePaths, int256[] values) returns(bool)
func (c *Contract) SubReplaceInt0(ctx context.Context, slot *big.Int, searchPath string, replacePaths []string, values []*big.Int) (bool, error) {
	storageVal := retrieveMsgSenderStorage(ctx, slot)
	updatedJson, err := c.jsonutil.SubReplaceInt0(ctx, storageVal, searchPath, replacePaths, values)
	if err != nil {
		return false, err
	}
	return updateMsgSenderStorage(ctx, slot, updatedJson)
}

// Solidity: function subReplaceUint(uint256 slot, string searchPath, string replacePath, uint256 value) returns(bool)
func (c *Contract) SubReplaceUint(ctx context.Context, slot *big.Int, searchPath string, replacePath string, value *big.Int) (bool, error) {
	return c.SubReplaceInt(ctx, slot, searchPath, replacePath, value)
}

// Solidity: function subReplaceUint(uint256 slot, string searchPath, string[] replacePaths, uint256[] values) returns(bool)
func (c *Contract) SubReplaceUint0(ctx context.Context, slot *big.Int, searchPath string, replacePaths []string, values []*big.Int) (bool, error) {
	return c.SubReplaceInt0(ctx, slot, searchPath, replacePaths, values)
}

// Solidity: function remove(uint256 slot, string path) returns(bool)
func (c *Contract) Remove(ctx context.Context, slot *big.Int, path string) (bool, error) {
	storageVal := retrieveMsgSenderStorage(ctx, slot)
	updatedJson, err := c.jsonutil.Remove(ctx, storageVal, path)
	if err != nil {
		return false, err
	}
	return updateMsgSenderStorage(ctx, slot, updatedJson)
}

func retrieveMsgSenderStorage(ctx context.Context, slot *big.Int) string {
	pCtx := pvm.UnwrapPolarContext(ctx)
	key := pCtx.MsgSender()
	return retrieveStorage(ctx, key, slot)
}

func updateMsgSenderStorage(ctx context.Context, slot *big.Int, jsonBlob string) (bool, error) {
	pCtx := pvm.UnwrapPolarContext(ctx)
	key := pCtx.MsgSender()
	return updateStorage(ctx, key, slot, jsonBlob)
}

func retrieveStorage(ctx context.Context, key common.Address, slot *big.Int) string {
	strRes := ""
	pCtx := pvm.UnwrapPolarContext(ctx)
	storageKey := crypto.Keccak256Hash(
		common.BytesToHash(key.Bytes()).Bytes(),
		common.BigToHash(slot).Bytes(),
	)
	storageVal := pCtx.GetState(storageKey).Bytes()
	if storageVal[0] != 0 {
		length := int(storageVal[len(storageVal)-1] / 2)
		res, _ := uncompress(storageVal[:length])
		strRes = string(res)
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
		strRes = string(res)
	}
	if strRes == "" {
		strRes = `{}`
	}
	return strRes
}

func updateStorage(ctx context.Context, key common.Address, slot *big.Int, jsonBlob string) (bool, error) {
	validationErr := fastjson.Validate(jsonBlob)
	if validationErr != nil {
		return false, errors.New("invalid JSON")
	}

	pCtx := pvm.UnwrapPolarContext(ctx)
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
