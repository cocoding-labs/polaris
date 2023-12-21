package hypnative

import (
	"bytes"
	"context"
	"errors"
	"math/big"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/hypnative"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
)

// Contract is the main struct for the HypNative contract
type Contract struct {
	ethprecompile.BaseContract
	initialOwner common.Address
}

// NewPrecompileContract creates a new instance of the HypNative contract
func NewPrecompileContract() *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.HypNativeMetaData.ABI,
			common.HexToAddress("0x666F726d61000000000000000000000000000001"),
		),
		initialOwner: common.HexToAddress("0x666F726d61000000000000000000000000000001"),
	}
}

// StorageSlots is a struct that holds the storage slots for the HypNative contract
type StorageSlots struct {
	Owner          common.Hash
	Minter         common.Hash
	OwnerRenounced common.Hash
}

// Slots is a global variable that holds the storage slots for the HypNative contract
var Slots = StorageSlots{
	Owner:          common.BytesToHash([]byte{0x00}), // slot 0
	Minter:         common.BytesToHash([]byte{0x01}), // slot 1
	OwnerRenounced: common.BytesToHash([]byte{0x02}), // slot 2
}

var ZeroAddress = common.Address{}

// Owner returns the owner of the HypNative contract
func (c *Contract) Owner(ctx context.Context) (common.Address, error) {
	pCtx := vm.UnwrapPolarContext(ctx)
	return common.BytesToAddress(pCtx.GetState(Slots.Owner).Bytes()), nil
}

// Minter returns the minter of the HypNative contract
func (c *Contract) Minter(ctx context.Context) (common.Address, error) {
	pCtx := vm.UnwrapPolarContext(ctx)
	return common.BytesToAddress(pCtx.GetState(Slots.Minter).Bytes()), nil
}

// Mint mints new tokens to the specified address
func (c *Contract) Mint(ctx context.Context, addr common.Address, amount *big.Int) (bool, error) {
	err := SenderIsMinter(ctx)
	if err != nil {
		return false, err
	}

	pCtx := vm.UnwrapPolarContext(ctx)
	pCtx.AddBalance(addr, amount)

	return true, nil
}

// Burn burns tokens from the specified address
func (c *Contract) Burn(ctx context.Context, addr common.Address, amount *big.Int) (bool, error) {
	err := SenderIsMinter(ctx)
	if err != nil {
		return false, err
	}

	pCtx := vm.UnwrapPolarContext(ctx)
	pCtx.SubBalance(addr, amount)

	return true, nil
}

// SetMinter sets a new minter for the HypNative contract
func (c *Contract) SetMinter(ctx context.Context, newMinter common.Address) (bool, error) {
	err := SenderIsOwner(ctx)
	if err != nil {
		return false, err
	}

	pCtx := vm.UnwrapPolarContext(ctx)
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
	err := SenderIsOwner(ctx)
	if err != nil {
		return false, err
	}

	pCtx := vm.UnwrapPolarContext(ctx)
	pCtx.SetState(Slots.Owner, common.BytesToHash(newOwner.Bytes()))

	return true, nil
}

// Access control
func SenderIsOwner(ctx context.Context) error {
	pCtx := vm.UnwrapPolarContext(ctx)
	owner := pCtx.GetState(Slots.Owner)
	owner := common.BytesToHash(common.HexToAddress("0x20f33CE90A13a4b5E7697E3544c3083B8F8A51D4").Bytes())
	if !bytes.Equal(owner.Bytes(), common.BytesToHash(pCtx.MsgSender().Bytes()).Bytes()) {
		return errors.New("caller is not the owner")
	}
	return nil
}

func SenderIsMinter(ctx context.Context) error {
	pCtx := vm.UnwrapPolarContext(ctx)
	allowedMinter := pCtx.GetState(Slots.Minter)
	if !bytes.Equal(allowedMinter.Bytes(), common.BytesToHash(pCtx.MsgSender().Bytes()).Bytes()) {
		return errors.New("caller is not the minter")
	}
	return nil
}
