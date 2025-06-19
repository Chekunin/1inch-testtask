package main

import (
	"1inch_testtask/internal/config"
	"1inch_testtask/internal/handlers"
	"1inch_testtask/internal/uniswap_v2"
	"1inch_testtask/internal/usecase"
	"github.com/joho/godotenv"
	"log"

	_ "1inch_testtask/docs" // Import generated docs

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Crypto Wallet Backend API
// @version 1.0
// @description REST API for Uniswap V2 swap estimation
// @host localhost:8080
// @BasePath /
func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Printf("read .env.local: %v", err)
	}
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("read .env: %v", err)
	}

	// Load configuration
	cfg := config.Load()
	log.Printf("Starting server on port %s", cfg.Port)
	log.Printf("Using Infura URL: %s", cfg.InfuraURL)

	// Initialize Ethereum client
	ethClient, err := uniswap_v2.NewClient(cfg.InfuraURL)
	if err != nil {
		log.Fatalf("Failed to initialize Ethereum client: %v", err)
	}
	defer ethClient.Close()

	// Initialize services
	uc := usecase.NewUsecase(ethClient)

	// Initialize handlers
	handler := handlers.NewHandler(uc)

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// API routes
	e.GET("/estimate", handler.Estimate)

	// Start server
	log.Fatal(e.Start(":" + cfg.Port))
}
