package handlers

import (
	"1inch_testtask/internal/models"
	"1inch_testtask/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler handles the /estimate endpoint
type Handler struct {
	uniswapService *usecase.Usecase
}

// NewHandler creates a new Handler
func NewHandler(uniswapService *usecase.Usecase) *Handler {
	return &Handler{
		uniswapService: uniswapService,
	}
}

// Estimate calculates the estimated output amount for a Uniswap V2 swap
// @Summary Calculate swap estimation
// @Description Estimates the output amount for a Uniswap V2 token swap based on current pool reserves
// @Tags estimate
// @Accept json
// @Produce json
// @Param pool query string true "Uniswap V2 pool address" example(0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852)
// @Param src query string true "Source token address" example(0xdAC17F958D2ee523a2206206994597C13D831ec7)
// @Param dst query string true "Destination token address" example(0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2)
// @Param src_amount query string true "Source amount to swap (integer with respect to decimals)" example(10000000)
// @Success 200 {object} models.EstimateResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /estimate [get]
func (h *Handler) Estimate(c echo.Context) error {
	var req models.EstimateRequest

	// Bind query parameters
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid_request",
			Message: "Failed to parse query parameters: " + err.Error(),
		})
	}

	// Validate request
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
	}

	// Calculate estimation
	outputAmount, err := h.uniswapService.EstimateSwap(
		c.Request().Context(),
		req.Pool,
		req.Src,
		req.Dst,
		req.SrcAmount,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "calculation_error",
			Message: "Failed to calculate swap estimation: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.EstimateResponse{
		DstAmount: outputAmount.String(),
	})
}
