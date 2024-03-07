package nativeminter

import (
	"bytes"
	"context"
	"errors"
	"math/big"

	generated "github.com/berachain/polaris/contracts/bindings/forma/precompile/nativeminter"
	ethprecompile "github.com/berachain/polaris/eth/core/precompile"
	pvm "github.com/berachain/polaris/eth/core/vm"
	"github.com/ethereum/go-ethereum/common"
)

// Contract is the main struct for the HypNative contract
type Contract struct {
	ethprecompile.BaseContract
}

// NewPrecompileContract creates a new instance of the HypNative contract
func NewPrecompileContract() *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.NativeMinterABI,
			common.HexToAddress("0x666F726d61000000000000000000000000000001"),
		),
	}
}

// StorageSlots is a struct that holds the storage slots for the HypNative contract
type StorageSlots struct {
	Owner  common.Hash
	Minter common.Hash
}

// Slots is a global variable that holds the storage slots for the HypNative contract
var Slots = StorageSlots{
	Owner:  common.BytesToHash([]byte{0x00}), // slot 0
	Minter: common.BytesToHash([]byte{0x01}), // slot 1
}

var ZeroAddress = common.Address{}

// Owner returns the owner of the HypNative contract
func (c *Contract) Owner(ctx context.Context) (common.Address, error) {
	pCtx := pvm.UnwrapPolarContext(ctx)
	return common.BytesToAddress(pCtx.GetState(Slots.Owner).Bytes()), nil
}

// Minter returns the minter of the HypNative contract
func (c *Contract) Minter(ctx context.Context) (common.Address, error) {
	pCtx := pvm.UnwrapPolarContext(ctx)
	return common.BytesToAddress(pCtx.GetState(Slots.Minter).Bytes()), nil
}

// Mint mints new tokens to the specified address
func (c *Contract) Mint(ctx context.Context, addr common.Address, amount *big.Int) (bool, error) {
	err := senderIsMinter(ctx)
	if err != nil {
		return false, err
	}

	pCtx := pvm.UnwrapPolarContext(ctx)
	pCtx.AddBalance(addr, amount)

	return true, nil
}

// Burn burns tokens from the specified address
func (c *Contract) Burn(ctx context.Context, addr common.Address, amount *big.Int) (bool, error) {
	err := senderIsMinter(ctx)
	if err != nil {
		return false, err
	}

	pCtx := pvm.UnwrapPolarContext(ctx)

	// check balance is valid
	balance := pCtx.GetBalance(addr)
	if balance.Cmp(amount) < 0 {
		return false, errors.New("insufficient balance for burn")
	}

	pCtx.SubBalance(addr, amount)

	return true, nil
}

// SetMinter sets a new minter for the HypNative contract
func (c *Contract) SetMinter(ctx context.Context, newMinter common.Address) (bool, error) {
	err := senderIsOwner(ctx)
	if err != nil {
		return false, err
	}

	pCtx := pvm.UnwrapPolarContext(ctx)
	pCtx.SetState(Slots.Minter, common.BytesToHash(newMinter.Bytes()))

	return true, nil
}

// TransferOwnership transfers the ownership of the HypNative contract to a new owner
func (c *Contract) TransferOwnership(ctx context.Context, newOwner common.Address) (bool, error) {
	if bytes.Equal(newOwner.Bytes(), ZeroAddress.Bytes()) {
		return false, errors.New("new owner is the zero address")
	}

	return c.internalTransferOwnership(ctx, newOwner)
}

// RenounceOwnership renounces the ownership of the HypNative contract
func (c *Contract) RenounceOwnership(ctx context.Context) (bool, error) {
	return c.internalTransferOwnership(ctx, ZeroAddress)
}

func (c *Contract) internalTransferOwnership(ctx context.Context, newOwner common.Address) (bool, error) {
	err := senderIsOwner(ctx)
	if err != nil {
		return false, err
	}

	pCtx := pvm.UnwrapPolarContext(ctx)
	pCtx.SetState(Slots.Owner, common.BytesToHash(newOwner.Bytes()))

	return true, nil
}

// Access control
func senderIsOwner(ctx context.Context) error {
	pCtx := pvm.UnwrapPolarContext(ctx)
	owner := common.BytesToAddress(pCtx.GetState(Slots.Owner).Bytes())
	if owner.Cmp(pCtx.MsgSender()) != 0 {
		return errors.New("caller is not the owner")
	}
	return nil
}

func senderIsMinter(ctx context.Context) error {
	pCtx := pvm.UnwrapPolarContext(ctx)
	allowedMinter := common.BytesToAddress(pCtx.GetState(Slots.Minter).Bytes())
	if allowedMinter.Cmp(pCtx.MsgSender()) != 0 {
		return errors.New("caller is not the minter")
	}
	return nil
}
