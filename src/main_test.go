package main_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"main.go/swap"
)

var nodeURL = ""
var walletAddress = ""
var privateKeyHex = ""

func TestSwapFunctions(t *testing.T) {

	//intialize web3client
	swap.InitializeWeb3Client(nodeURL)
	//initialize uniswapRouter02 with wallet address
	swap.InitializeUniswapRouter02(walletAddress)
	//test new wallet key for signing
	key := swap.NewWalletKey(privateKeyHex)
	if key == nil {
		t.Fatal("NewWalletKey Error")
	}

	// //test approve function
	// err := swap.Approve("0xc778417E063141139Fce010982780140Aa0cD5Ab", walletAddress, key, 3)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	//set lp path
	path := []common.Address{common.HexToAddress("0xc778417E063141139Fce010982780140Aa0cD5Ab"), common.HexToAddress("0x110a13FC3efE6A245B50102D2d79B3E76125Ae83")}
	//reciever wallet address
	to := common.HexToAddress(walletAddress)
	//minimum amount out
	amountOutMin := uint(0)
	//exact amount in
	amountIn := uint(1000000)
	//timeout deadline
	deadline := uint(10000000000000000000)

	//SwapExactTokensForTokens
	txn0 := swap.SwapExactTokensForTokens(amountIn, amountOutMin, path, to, deadline)
	if txn0 == nil {
		t.Fatal("SwapExactTokensForTokens Error")
	}

	err0 := txn0.SignAndSend(key, 3)
	if err0 != nil {
		t.Fatal("Sign and send Error", err0)
	}

}
