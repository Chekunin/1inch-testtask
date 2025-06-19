package usecase

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to create big.Int from string
func mustBigInt(s string) *big.Int {
	val, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("invalid big int string: " + s)
	}
	return val
}

func TestService_calculateOutputAmount(t *testing.T) {
	service := &Usecase{}

	tests := []struct {
		name        string
		amountIn    *big.Int
		reserveIn   *big.Int
		reserveOut  *big.Int
		expected    *big.Int
		description string
	}{
		{
			name:        "basic calculation",
			amountIn:    big.NewInt(1000000),                 // 1 USDT (6 decimals)
			reserveIn:   big.NewInt(1000000000000),           // 1M USDT
			reserveOut:  mustBigInt("500000000000000000000"), // 500 ETH
			expected:    mustBigInt("498499502995995"),       // Expected output
			description: "1 USDT in, should get ~0.498 ETH out",
		},
		{
			name:        "zero amount in",
			amountIn:    big.NewInt(0),
			reserveIn:   big.NewInt(1000000000000),
			reserveOut:  mustBigInt("500000000000000000000"),
			expected:    big.NewInt(0),
			description: "Zero input should return zero output",
		},
		{
			name:        "zero reserve in",
			amountIn:    big.NewInt(1000000),
			reserveIn:   big.NewInt(0),
			reserveOut:  mustBigInt("500000000000000000000"),
			expected:    big.NewInt(0),
			description: "Zero reserve in should return zero output",
		},
		{
			name:        "zero reserve out",
			amountIn:    big.NewInt(1000000),
			reserveIn:   big.NewInt(1000000000000),
			reserveOut:  big.NewInt(0),
			expected:    big.NewInt(0),
			description: "Zero reserve out should return zero output",
		},
		{
			name:        "large amount calculation",
			amountIn:    big.NewInt(100000000000),             // 100k USDT
			reserveIn:   big.NewInt(10000000000000),           // 10M USDT
			reserveOut:  mustBigInt("5000000000000000000000"), // 5k ETH
			expected:    mustBigInt("49357901719853064942"),   // Expected output
			description: "Large amount swap calculation",
		},
		{
			name:        "small amount calculation",
			amountIn:    big.NewInt(1),                     // 1 wei
			reserveIn:   mustBigInt("1000000000000000000"), // 1 ETH
			reserveOut:  big.NewInt(2000000000),            // 2000 USDT
			expected:    big.NewInt(0),                     // Expected output (rounds down)
			description: "Very small amount swap",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.calculateOutputAmount(tt.amountIn, tt.reserveIn, tt.reserveOut)
			assert.Equal(t, tt.expected, result, tt.description)
		})
	}
}

// TestUniswapV2Formula verifies the Uniswap V2 formula implementation
func TestUniswapV2Formula(t *testing.T) {
	service := &Usecase{}

	// Test case based on real Uniswap V2 pair data
	// This test verifies that our formula matches the expected Uniswap V2 calculation
	amountIn := big.NewInt(1000000)                     // 1 USDT (6 decimals)
	reserveIn := big.NewInt(50000000000000)             // 50M USDT reserve
	reserveOut := mustBigInt("20000000000000000000000") // 20k ETH reserve

	result := service.calculateOutputAmount(amountIn, reserveIn, reserveOut)

	// Manual calculation:
	// amountInWithFee = 1000000 * 0.97 = 997000
	// numerator = 997000 * 20000000000000000000000 = 1.994 * 10^28
	// denominator = 50000000000000 + 997000 = 50000000997000
	// result = 1.994 * 10^28 / 50000000997000 ~ 398799992047928,1585643125

	expected := mustBigInt("398799992047928")
	assert.Equal(t, expected, result, "Uniswap V2 formula calculation should match expected result")
}
