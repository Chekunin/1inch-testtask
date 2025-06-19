package usecase

import (
	"1inch_testtask/internal/uniswap_v2"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Usecase handles Uniswap V2 calculations
type Usecase struct {
	uniswapV2Client uniswap_v2.IUniswapV2
}

// NewUsecase creates a new Uniswap service
func NewUsecase(uniswapV2Client uniswap_v2.IUniswapV2) *Usecase {
	return &Usecase{
		uniswapV2Client: uniswapV2Client,
	}
}

// EstimateSwap calculates the output amount for a Uniswap V2 swap
func (s *Usecase) EstimateSwap(ctx context.Context, poolAddr, srcAddr, dstAddr, srcAmountStr string) (*big.Int, error) {
	// Parse source amount
	srcAmount, ok := new(big.Int).SetString(srcAmountStr, 10)
	if !ok {
		return nil, fmt.Errorf("invalid src_amount: %s", srcAmountStr)
	}

	// Convert addresses
	poolAddress := common.HexToAddress(poolAddr)
	srcAddress := common.HexToAddress(srcAddr)
	dstAddress := common.HexToAddress(dstAddr)

	// Get token addresses from the pool
	token0, err := s.uniswapV2Client.GetToken0(ctx, poolAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get token0: %w", err)
	}

	token1, err := s.uniswapV2Client.GetToken1(ctx, poolAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get token1: %w", err)
	}

	// Get reserves
	reserve0, reserve1, err := s.uniswapV2Client.GetReserves(ctx, poolAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get reserves: %w", err)
	}

	// Determine which token is which and get the appropriate reserves
	var reserveIn, reserveOut *big.Int
	if srcAddress == token0 && dstAddress == token1 {
		reserveIn = reserve0
		reserveOut = reserve1
	} else if srcAddress == token1 && dstAddress == token0 {
		reserveIn = reserve1
		reserveOut = reserve0
	} else {
		return nil, fmt.Errorf("token pair mismatch: src=%s, dst=%s, token0=%s, token1=%s",
			srcAddr, dstAddr, token0.Hex(), token1.Hex())
	}

	// Calculate output amount using Uniswap V2 formula
	outputAmount := s.calculateOutputAmount(srcAmount, reserveIn, reserveOut)

	return outputAmount, nil
}

// calculateOutputAmount implements the Uniswap V2 swap formula
// amountOut = (amountIn * 997 * reserveOut) / (reserveIn * 1000 + amountIn * 997)
// This accounts for the 0.3% fee (997/1000 = 0.997)
func (s *Usecase) calculateOutputAmount(amountIn, reserveIn, reserveOut *big.Int) *big.Int {
	if amountIn.Cmp(big.NewInt(0)) <= 0 {
		return big.NewInt(0)
	}
	if reserveIn.Cmp(big.NewInt(0)) <= 0 || reserveOut.Cmp(big.NewInt(0)) <= 0 {
		return big.NewInt(0)
	}

	// amountInWithFee = amountIn * 997
	amountInWithFee := new(big.Int).Mul(amountIn, big.NewInt(997))

	// numerator = amountInWithFee * reserveOut
	numerator := new(big.Int).Mul(amountInWithFee, reserveOut)

	// denominator = reserveIn * 1000 + amountInWithFee
	denominator := new(big.Int).Mul(reserveIn, big.NewInt(1000))
	denominator.Add(denominator, amountInWithFee)

	// amountOut = numerator / denominator
	amountOut := new(big.Int).Div(numerator, denominator)

	return amountOut
}
