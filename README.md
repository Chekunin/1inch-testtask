# 1inch test task (implemented by Aleksei Chekunin)

A REST API backend for Uniswap V2 swap estimation that calculates output amounts based on current pool reserves.

## How to run
You can run with docker:
```bash
docker compose up --build
```

Or you can run it locally:
```bash
make run
```

The server will start on `http://localhost:8080`

## Features

- **Single endpoint** `/estimate` for Uniswap V2 swap calculations
- **Real-time data** from Ethereum mainnet via Infura
- **Accurate calculations** using Uniswap V2 formula with 0.3% fee
- **Input validation** for addresses and amounts
- **Swagger documentation** available at `/swagger/`
- **Health check** endpoint at `/health`
- **Comprehensive testing** with unit tests

## Tech Stack

- **Go 1.23+** - Programming language
- **Echo v4** - Web framework
- **Geth** - Ethereum client library
- **Swaggo** - Swagger documentation generation
- **Testify** - Testing framework

## API Documentation

### Estimate Endpoint

**GET** `/estimate`

Calculates the estimated output amount for a Uniswap V2 token swap.

#### Query Parameters

| Parameter | Type | Required | Description | Example |
|-----------|------|----------|-------------|---------|
| `pool` | string | Yes | Uniswap V2 pool address | `0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852` |
| `src` | string | Yes | Source token address | `0xdAC17F958D2ee523a2206206994597C13D831ec7` |
| `dst` | string | Yes | Destination token address | `0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2` |
| `src_amount` | string | Yes | Source amount (integer with respect to decimals) | `10000000` |

#### Example Request

```bash
curl "http://localhost:8080/estimate?pool=0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852&src=0xdAC17F958D2ee523a2206206994597C13D831ec7&dst=0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2&src_amount=10000000"
```

#### Example Response

```json
{
  "dst_amount": "3978866028279530"
}
```

#### Error Responses

**400 Bad Request** - Invalid parameters:
```json
{
  "error": "validation_error",
  "message": "invalid pool address: address must be 40 hex characters"
}
```

**500 Internal Server Error** - Calculation error:
```json
{
  "error": "calculation_error",
  "message": "Failed to calculate swap estimation: failed to get reserves"
}
```

### Health Check

**GET** `/health`

Returns the service health status.

```json
{
  "status": "ok"
}
```

### Swagger Documentation

Interactive API documentation is available at: `http://localhost:8080/swagger/`

## Uniswap V2 Formula

The service uses the standard Uniswap V2 constant product formula with fee:

```
amountOut = (amountIn * 997 * reserveOut) / (reserveIn * 1000 + amountIn * 997)
```

Where:
- `997/1000 = 0.997` accounts for the 0.3% trading fee
- `reserveIn` and `reserveOut` are the current pool reserves
- `amountIn` is the input token amount
