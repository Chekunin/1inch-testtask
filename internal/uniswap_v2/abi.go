package uniswap_v2

// UniswapV2PairABI is the ABI for Uniswap V2 Pair contract
const UniswapV2PairABI = `[
	{
		"constant": true,
		"inputs": [],
		"name": "getReserves",
		"outputs": [
			{"name": "_reserve0", "type": "uint112"},
			{"name": "_reserve1", "type": "uint112"},
			{"name": "_blockTimestampLast", "type": "uint32"}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "token0",
		"outputs": [{"name": "", "type": "address"}],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "token1",
		"outputs": [{"name": "", "type": "address"}],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]`
