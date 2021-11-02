package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/go-web3"
	"github.com/umbracle/go-web3/abi"
	"github.com/umbracle/go-web3/contract"
	"github.com/umbracle/go-web3/jsonrpc"
)

//Web3 HTTP client to send calls to the RPC enpoint
var web3Client = initWeb3Client()

var uniswapRouter02 = initUniswapRouter02()

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
func initUniswapRouter02() *contract.Contract {
	//initialize a web3 address with the uniswap router hex address
	uniswapRouter02Address := web3.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	//read in the uniswapRouter02 abi from file
	abiBytes, err := ioutil.ReadFile("uniswapRouter02ABI.json")
	if err != nil {
		fmt.Println("Error when reading UniswapRouter02 ABI from file.")
	}

	//create a new web3 abi
	abi, err := abi.NewABI(string(abiBytes))
	if err != nil {
		fmt.Println("Error when creating UniswapRouter02ABI", err)
		return nil
	}

	return contract.NewContract(uniswapRouter02Address, abi, web3Client)
}

//Uniswap Router02 swap functions--------------------------

func SwapExactTokensForTokens(amountIn uint, amountOutMin uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouter02.Txn("swapExactTokensForTokens", web3.Latest, amountIn, amountOutMin, path, to, deadline)
	return txn
}

func SwapTokensForExactTokens(amountOut uint, amountInMax uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouter02.Txn("swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
	return txn
}

func SwapExactETHForTokens(amountOutMin uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouter02.Txn("swapExactETHForTokens", amountOutMin, path, to, deadline)
	return txn
}

func SwapTokensForExactETH(amountOut uint, amountInMax uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouter02.Txn("swapTokensForExactETH", amountOut, amountInMax, path, to, deadline)
	return txn
}

func SwapExactTokensForETH(amountIn uint, amountOutMin uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouter02.Txn("swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
	return txn
}

func SwapETHForExactTokens(amountOut uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouter02.Txn("swapETHForExactTokens", amountOut, path, to, deadline)
	return txn
}

func SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn uint, amountOutMin uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouter02.Txn("swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
	return txn
}

func SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouter02.Txn("swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
	return txn
}

func SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn uint, amountOutMin uint, path []common.Address, to common.Address, deadline uint) *contract.Txn {
	txn := uniswapRouter02.Txn("swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
	return txn
}
