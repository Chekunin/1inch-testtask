package models

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// EstimateRequest represents the request parameters for the /estimate endpoint
type EstimateRequest struct {
	Pool      string `query:"pool" validate:"required" example:"0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852"`
	Src       string `query:"src" validate:"required" example:"0xdAC17F958D2ee523a2206206994597C13D831ec7"`
	Dst       string `query:"dst" validate:"required" example:"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"`
	SrcAmount string `query:"src_amount" validate:"required" example:"10000000"`
}

// EstimateResponse represents the response for the /estimate endpoint
type EstimateResponse struct {
	DstAmount string `json:"dst_amount" example:"6241000000000000"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Validate validates the EstimateRequest
func (r *EstimateRequest) Validate() error {
	if err := validateAddress(r.Pool); err != nil {
		return errors.New("invalid pool address: " + err.Error())
	}
	if err := validateAddress(r.Src); err != nil {
		return errors.New("invalid src address: " + err.Error())
	}
	if err := validateAddress(r.Dst); err != nil {
		return errors.New("invalid dst address: " + err.Error())
	}

	amount, err := strconv.ParseUint(r.SrcAmount, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid amount format: %s", r.SrcAmount)
	}
	if amount == 0 {
		return fmt.Errorf("amount must be greater than 0")
	}

	return nil
}

// validateAddress validates Ethereum address format
func validateAddress(address string) error {
	if address == "" {
		return errors.New("address cannot be empty")
	}

	// Remove 0x prefix if present
	addr := strings.TrimPrefix(address, "0x")

	// Check length (40 hex characters)
	if len(addr) != 40 {
		return errors.New("address must be 40 hex characters")
	}

	// Check if it's valid hex
	matched, err := regexp.MatchString("^[0-9a-fA-F]+$", addr)
	if err != nil || !matched {
		return errors.New("address must contain only hex characters")
	}

	return nil
}
