package bridge

//go:generate abigen --sol contract/Bridge.sol --pkg contract --out contract/bridge.go
// abigen v1.10.1
// solc 0.8.1

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/go-bamboo/contrib/contracts/bridge/contract"
	"github.com/go-kratos/kratos/v2/errors"
)

// Bridge is a Go wrapper around an on-chain checkpoint oracle contract.
type Bridge struct {
	address  common.Address
	contract *contract.Bridge
}

// NewBridge binds checkpoint contract and returns a registrar instance.
func NewBridge(contractAddr common.Address, backend bind.ContractBackend) (ctrt *Bridge, err error) {
	c, err := contract.NewBridge(contractAddr, backend)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	ctrt = &Bridge{address: contractAddr, contract: c}
	return
}

// ContractAddr returns the address of contract.
func (bridge *Bridge) ContractAddr() common.Address {
	return bridge.address
}

// Contract returns the underlying contract instance.
func (bridge *Bridge) Contract() *contract.Bridge {
	return bridge.contract
}

func (bridge *Bridge) GetIn(ctx context.Context, from common.Address, offset *big.Int) (contract.BridgeTokenIn, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	in, err := bridge.contract.GetIn(opts, offset)
	if err != nil {
		return in, errors.FromError(err)
	}
	return in, nil
}

func (bridge *Bridge) GetOut(ctx context.Context, from common.Address, offset *big.Int) (contract.BridgeTokenOut, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	out, err := bridge.contract.GetOut(opts, offset)
	if err != nil {
		return out, errors.FromError(err)
	}
	return out, nil
}

func (bridge *Bridge) GetInLength(ctx context.Context, from common.Address) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	out, err := bridge.contract.GetInLength(opts)
	if err != nil {
		return nil, errors.FromError(err)
	}
	return out, err
}

// GetBalance 获取主币余额
func (bridge *Bridge) GetBalance(ctx context.Context, from common.Address) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	out, err := bridge.contract.GetBalance(opts)
	if err != nil {
		return nil, errors.FromError(err)
	}
	return out, nil
}

// Owner 获取主币余额
func (bridge *Bridge) Owner(ctx context.Context, from common.Address) (common.Address, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	out, err := bridge.contract.Owner(opts)
	if err != nil {
		return out, errors.FromError(err)
	}
	return out, nil
}

func (bridge *Bridge) Operator1(ctx context.Context, from common.Address) (common.Address, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	out, err := bridge.contract.Operator1(opts)
	if err != nil {
		return out, errors.FromError(err)
	}
	return out, nil
}

func (bridge *Bridge) Operator2(ctx context.Context, from common.Address) (common.Address, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	out, err := bridge.contract.Operator2(opts)
	if err != nil {
		return out, errors.FromError(err)
	}
	return out, nil
}

func (bridge *Bridge) DepositToken(ctx context.Context, chainID *big.Int, fromPriv *ecdsa.PrivateKey, nonce *big.Int, token common.Address, value *big.Int) (txHash, rawTx string, err error) {
	opts, err := bind.NewKeyedTransactorWithChainID(fromPriv, chainID)
	if err != nil {
		return
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.GasLimit = 100000
	tx, err := bridge.contract.DepositToken(opts, token, value)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	txHash = tx.Hash().Hex()
	rawTxBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	rawTx = hexutil.Encode(rawTxBytes)
	return
}

func (bridge *Bridge) Withdraw(ctx context.Context, chainID *big.Int, fromPriv *ecdsa.PrivateKey, nonce *big.Int, token common.Address) (txHash, rawTx string, err error) {
	opts, err := bind.NewKeyedTransactorWithChainID(fromPriv, chainID)
	if err != nil {
		return
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.Value = big.NewInt(0)
	opts.GasLimit = 100000

	tx, err := bridge.contract.Withdraw(opts, token)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	txHash = tx.Hash().Hex()
	rawTxBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	rawTx = hexutil.Encode(rawTxBytes)
	return
}

func (bridge *Bridge) SendToken(ctx context.Context, chainID *big.Int, fromPriv *ecdsa.PrivateKey, nonce *big.Int, fromIndex *big.Int, token common.Address, to common.Address, value *big.Int) (txHash, rawTx string, err error) {
	opts, err := bind.NewKeyedTransactorWithChainID(fromPriv, chainID)
	if err != nil {
		return
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.Value = big.NewInt(0)
	opts.GasLimit = 400000

	tx, err := bridge.contract.SendToken(opts, fromIndex, token, to, value)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	txHash = tx.Hash().Hex()
	rawTxBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	rawTx = hexutil.Encode(rawTxBytes)
	return
}

func (bridge *Bridge) SetOperator(ctx context.Context, chainID *big.Int, fromPriv *ecdsa.PrivateKey, nonce *big.Int, operator common.Address, index *big.Int) (txHash, rawTx string, err error) {
	opts, err := bind.NewKeyedTransactorWithChainID(fromPriv, chainID)
	if err != nil {
		return
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.Value = big.NewInt(0)
	opts.GasLimit = 100000

	tx, err := bridge.contract.SetOperator(opts, operator, index)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	txHash = tx.Hash().Hex()
	rawTxBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	rawTx = hexutil.Encode(rawTxBytes)
	return
}

func (bridge *Bridge) PauseToken(ctx context.Context, chainID *big.Int, fromPriv *ecdsa.PrivateKey, nonce *big.Int, token common.Address) (txHash, rawTx string, err error) {
	opts, err := bind.NewKeyedTransactorWithChainID(fromPriv, chainID)
	if err != nil {
		return
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.Value = big.NewInt(0)
	opts.GasLimit = 100000

	tx, err := bridge.contract.PauseToken(opts, token)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	txHash = tx.Hash().Hex()
	rawTxBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	rawTx = hexutil.Encode(rawTxBytes)
	return
}

func (bridge *Bridge) UnpauseToken(ctx context.Context, chainID *big.Int, fromPriv *ecdsa.PrivateKey, nonce *big.Int, token common.Address) (txHash, rawTx string, err error) {
	opts, err := bind.NewKeyedTransactorWithChainID(fromPriv, chainID)
	if err != nil {
		return
	}
	opts.Nonce = nonce
	opts.Value = big.NewInt(0)
	opts.GasLimit = 100000

	tx, err := bridge.contract.UnpauseToken(opts, token)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	txHash = tx.Hash().Hex()
	rawTxBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	rawTx = hexutil.Encode(rawTxBytes)
	return
}
