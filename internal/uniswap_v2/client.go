package uniswap_v2

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IUniswapV2 interface {
	GetReserves(ctx context.Context, poolAddress common.Address) (*big.Int, *big.Int, error)
	GetToken0(ctx context.Context, poolAddress common.Address) (common.Address, error)
	GetToken1(ctx context.Context, poolAddress common.Address) (common.Address, error)
	Close()
}

// Client wraps the Ethereum client
type Client struct {
	client    *ethclient.Client
	parsedABI abi.ABI
}

// NewClient creates a new Ethereum client
func NewClient(infuraURL string) (*Client, error) {
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(UniswapV2PairABI))
	if err != nil {
		return nil, err
	}

	return &Client{
		client:    client,
		parsedABI: parsedABI,
	}, nil
}

// Close closes the Ethereum client connection
func (c *Client) Close() {
	c.client.Close()
}

// GetReserves gets the reserves from a Uniswap V2 pair
func (c *Client) GetReserves(ctx context.Context, poolAddress common.Address) (*big.Int, *big.Int, error) {
	contract := bind.NewBoundContract(poolAddress, c.parsedABI, c.client, c.client, c.client)

	var out []interface{}
	if err := contract.Call(&bind.CallOpts{Context: ctx}, &out, "getReserves"); err != nil {
		return nil, nil, err
	}

	return out[0].(*big.Int), out[1].(*big.Int), nil
}

// GetToken0 gets token0 address from the pair
func (c *Client) GetToken0(ctx context.Context, poolAddress common.Address) (common.Address, error) {
	contract := bind.NewBoundContract(poolAddress, c.parsedABI, c.client, c.client, c.client)

	var out []interface{}
	if err := contract.Call(&bind.CallOpts{Context: ctx}, &out, "token0"); err != nil {
		return common.Address{}, err
	}

	return out[0].(common.Address), nil
}

// GetToken1 gets token1 address from the pair
func (c *Client) GetToken1(ctx context.Context, poolAddress common.Address) (common.Address, error) {
	contract := bind.NewBoundContract(poolAddress, c.parsedABI, c.client, c.client, c.client)

	var out []interface{}
	if err := contract.Call(&bind.CallOpts{Context: ctx}, &out, "token1"); err != nil {
		return common.Address{}, err
	}

	return out[0].(common.Address), nil
}
