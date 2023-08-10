package erc20

//go:generate abigen --sol contract/USDT.sol --pkg contract --out contract/usdt.go

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/go-bamboo/contrib/contracts/erc20/contract"
	"github.com/go-kratos/kratos/v2/errors"
)

// Erc20 is a Go wrapper around an on-chain checkpoint oracle contract.
type Erc20 struct {
	address  common.Address
	contract *contract.ERC20
}

// NewErc20 binds checkpoint contract and returns a registrar instance.
func NewErc20(contractAddr common.Address, backend bind.ContractBackend) (ctrt *Erc20, err error) {
	c, err := contract.NewERC20(contractAddr, backend)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	ctrt = &Erc20{address: contractAddr, contract: c}
	return
}

// ContractAddr returns the address of contract.
func (c *Erc20) ContractAddr() common.Address {
	return c.address
}

// Contract returns the underlying contract instance.
func (c *Erc20) Contract() *contract.ERC20 {
	return c.contract
}

func (c *Erc20) BalanceOf(ctx context.Context, from common.Address) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	in, err := c.contract.BalanceOf(opts, from)
	if err != nil {
		return in, errors.FromError(err)
	}
	return in, nil
}

func (c *Erc20) Decimals(ctx context.Context) (uint8, error) {
	opts := &bind.CallOpts{
		Context: ctx,
	}
	in, err := c.contract.Decimals(opts)
	if err != nil {
		return in, errors.FromError(err)
	}
	return in, nil
}

func (c *Erc20) TransferFrom(ctx context.Context, chainID *big.Int, pk *ecdsa.PrivateKey, nonce *big.Int, from common.Address, to common.Address, value *big.Int) (txHash, rawTx string, err error) {
	opts, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		return
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.Value = big.NewInt(0)
	tx, err := c.contract.TransferFrom(opts, from, to, value)
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
