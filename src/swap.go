package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/go-web3"
	"github.com/umbracle/go-web3/abi"
	"github.com/umbracle/go-web3/contract"
	"github.com/umbracle/go-web3/jsonrpc"
)

//Web3 HTTP client to send calls to the RPC enpoint
var web3Client = initWeb3Client()

var uniswapRouterV2 = initUniswapRouterV2()

//initialize web3 http client
func initWeb3Client() *jsonrpc.Client {
	//TODO: need to pass in node url
	client, err := jsonrpc.NewClient("")
	if err != nil {
		fmt.Println("Failed to connect", err)
		return nil
	}
	return client
}

//initialize uniswap router v2 contract instance
func initUniswapRouterV2() *contract.Contract {
	//initialize a web3 address with the uniswap router hex address
	uniswapRouterV2Address := web3.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	//TODO: need to read in abi from file
	//read in the uniswapRouterV2 abi
	abi, err := abi.NewABI("abi from file")
	if err != nil {
		fmt.Println("Error when reading Uniswap ABI", err)
		return nil
	}
	return contract.NewContract(uniswapRouterV2Address, abi, web3Client)
}

//Uniswap Router02 swap functions--------------------------

func SwapExactTokensForTokens(amountIn uint, amountOutMin uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouterV2.Txn("swapExactTokensForTokens", web3.Latest, amountIn, amountOutMin, path, to, deadline)
	return txn
}

func SwapTokensForExactTokens(amountOut uint, amountInMax uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouterV2.Txn("swapTokensForExactTokens", web3.Latest, amountOut, amountInMax, path, to, deadline)
	return txn
}

func SwapExactETHForTokens(amountOutMin uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouterV2.Txn("swapExactETHForTokens", web3.Latest, amountOutMin, path, to, deadline)
	return txn
}

func SwapTokensForExactETH(amountOut uint, amountInMax uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouterV2.Txn("swapTokensForExactETH", web3.Latest, amountOut, amountInMax, path, to, deadline)
	return txn
}

func SwapExactTokensForETH(amountIn uint, amountOutMin uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouterV2.Txn("swapExactTokensForETH", web3.Latest, amountIn, amountOutMin, path, to, deadline)
	return txn
}

func SwapETHForExactTokens(amountOut uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouterV2.Txn("swapETHForExactTokens", web3.Latest, amountOut, path, to, deadline)
	return txn
}

func SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn uint, amountOutMin uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouterV2.Txn("swapExactTokensForTokensSupportingFeeOnTransferTokens", web3.Latest, amountIn, amountOutMin, path, to, deadline)
	return txn
}

func SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouterV2.Txn("swapExactETHForTokensSupportingFeeOnTransferTokens", web3.Latest, amountOutMin, path, to, deadline)
	return txn
}

func SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn uint, amountOutMin uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouterV2.Txn("swapExactTokensForETHSupportingFeeOnTransferTokens", web3.Latest, amountIn, amountOutMin, path, to, deadline)
	return txn
}
