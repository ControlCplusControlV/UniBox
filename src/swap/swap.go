package swap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/D-Cous/go-web3"
	"github.com/D-Cous/go-web3/abi"
	"github.com/D-Cous/go-web3/contract"
	"github.com/D-Cous/go-web3/jsonrpc"
	"github.com/D-Cous/go-web3/wallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//

//import from json
var WalletAddress string

var walletKey *wallet.Key

//Web3 HTTP client to send calls to the RPC enpoint
//import from json
var chainID = uint64(3)

var uniswapRouterV2 = initializeUniswapRouterV2()
var uniswapRouterV3 = initializeUniswapRouterV3()

var web3Client *jsonrpc.Client

func loadConfig() {
	content, err := ioutil.ReadFile("config.json")

	if err != nil {
		fmt.Println(err)
	}

	var config map[string]interface{}

	json.Unmarshal(content, &config)

	WalletAddress = config["publicKey"].(string)

	walletKey = NewWalletKey(config["privateKey"].(string))

	web3Client = initializeWeb3Client(config["alchemyURL"].(string))

}

func Init() {
	loadConfig()
}

//Returns *wallet.Key which is used to sign transactions
func NewWalletKey(privateKey string) *wallet.Key {
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	key := wallet.NewKey(ecdsaPrivateKey)

	return key
}

//initialize web3 http client
func initializeWeb3Client(nodeURL string) *jsonrpc.Client {
	client, err := jsonrpc.NewClient(nodeURL)
	if err != nil {
		fmt.Println("Failed to connect to node", err)
		os.Exit(1)
	}
	return client
}

//initialize uniswap router v2 contract instance
func initializeUniswapRouterV2() *contract.Contract {
	//initialize a web3 address with the uniswap router hex address
	uniswapRouterV2Address := web3.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	//read in the uniswapRouterV2 abi from file
	abiBytes, err := ioutil.ReadFile("abi/uniswapRouterV2ABI.json")
	if err != nil {
		fmt.Println("Error when reading uniswapRouterV2 ABI from file.")
	}
	//create a new web3 abi
	abi, err := abi.NewABI(string(abiBytes))
	if err != nil {
		fmt.Println("Error when creating uniswapRouterV2ABI", err)
		os.Exit(1)
	}
	contractInstance := contract.NewContract(uniswapRouterV2Address, abi, web3Client)
	//set the from address to the user wallet address
	contractInstance.SetFrom(web3.HexToAddress(WalletAddress))
	return contractInstance
}

//initialize uniswap router v3 contract instance
func initializeUniswapRouterV3() *contract.Contract {
	//initialize a web3 address with the uniswap router hex address
	uniswapRouterV2Address := web3.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564")
	//read in the uniswapRouterV2 abi from file
	abiBytes, err := ioutil.ReadFile("abi/uniswapRouterV3ABI.json")
	if err != nil {
		fmt.Println("Error when reading uniswapRouterV2 ABI from file.")
	}

	//create a new web3 abi
	abi, err := abi.NewABI(string(abiBytes))
	if err != nil {
		fmt.Println("Error when creating uniswapRouterV2ABI", err)
		os.Exit(1)
	}
	contractInstance := contract.NewContract(uniswapRouterV2Address, abi, web3Client)
	//set the from address to the user wallet address
	contractInstance.SetFrom(web3.HexToAddress(WalletAddress))
	return contractInstance
}

func GetCurrentBlock() (uint64, error) {
	blockNumber, err := web3Client.Eth().BlockNumber()
	if err != nil {
		return 0, err
	} else {
		return blockNumber, nil
	}
}

func Approve(contractAddress string) error {
	walletAddress := WalletAddress
	key := walletKey
	chainID = chainID
	//initialize a web3 address with the uniswap router hex address
	web3ContractAddress := web3.HexToAddress(contractAddress)
	//read in the uniswapRouterV2 abi from file
	abiBytes, err := ioutil.ReadFile("abi/erc20ABI.json")
	if err != nil {
		fmt.Println("Error when reading ERC20 ABI from file.")
	}

	//create a new web3 abi
	abi, err := abi.NewABI(string(abiBytes))
	if err != nil {
		fmt.Println("Error when creating ERC20ABI", err)
		return nil
	}
	contractInstance := contract.NewContract(web3ContractAddress, abi, web3Client)
	//set the from addresss
	contractInstance.SetFrom(web3.HexToAddress(walletAddress))
	//approve the token to interact with the uniswap router
	approveTxn := contractInstance.Txn("approve", web3.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"), 9000000000000000000)

	//send the transaction
	err = approveTxn.SignAndSend(key, chainID)
	if err != nil {
		return err
	}

	return nil
}

//Uniswap V3 functions--------------------------
type ExactInputSingleData struct {
	tokenIn           common.Address
	tokenOut          common.Address
	fee               uint
	recipient         common.Address
	deadline          big.Int
	amountIn          big.Int
	amountOutMinimum  big.Int
	sqrtPriceLimitX96 big.Int
}

func ExactInputSingle(tokenIn common.Address, tokenOut common.Address, fee uint, recipient common.Address, deadline big.Int, amountIn big.Int, amountOutMinimum big.Int, sqrtPriceLimitX96 big.Int) web3.Hash {

	exactInputSingleData := ExactInputSingleData{
		tokenIn:           tokenIn,
		tokenOut:          tokenOut,
		fee:               fee,
		recipient:         recipient,
		deadline:          deadline,
		amountIn:          amountIn,
		amountOutMinimum:  amountOutMinimum,
		sqrtPriceLimitX96: sqrtPriceLimitX96,
	}
	txn := uniswapRouterV3.Txn("exactInputSingle", exactInputSingleData)
	err := txn.SignAndSend(walletKey, chainID)
	if err != nil {
		panic(err)
	}
	return txn.Hash
}

//Uniswap V2 functions--------------------------

func SwapExactTokensForTokens(amountIn uint, amountOutMin uint, path []common.Address, to common.Address, deadline uint) web3.Hash {
	txn := uniswapRouterV2.Txn("swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
	err := txn.SignAndSend(walletKey, chainID)
	if err != nil {
		panic(err)
	}
	return txn.Hash
}

func SwapTokensForExactTokens(amountOut uint, amountInMax uint, path []common.Address, to common.Address, deadline uint) web3.Hash {
	txn := uniswapRouterV2.Txn("swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
	err := txn.SignAndSend(walletKey, chainID)
	if err != nil {
		panic(err)
	}
	return txn.Hash
}

func SwapExactETHForTokens(amountOutMin uint, path []common.Address, to common.Address, deadline uint) web3.Hash {
	txn := uniswapRouterV2.Txn("swapExactETHForTokens", amountOutMin, path, to, deadline)
	err := txn.SignAndSend(walletKey, chainID)
	if err != nil {
		panic(err)
	}
	return txn.Hash
}

func SwapTokensForExactETH(amountOut uint, amountInMax uint, path []common.Address, to common.Address, deadline uint) web3.Hash {
	txn := uniswapRouterV2.Txn("swapTokensForExactETH", amountOut, amountInMax, path, to, deadline)
	err := txn.SignAndSend(walletKey, chainID)
	if err != nil {
		panic(err)
	}
	return txn.Hash
}

func SwapExactTokensForETH(amountIn uint, amountOutMin uint, path []common.Address, to common.Address, deadline uint) web3.Hash {
	txn := uniswapRouterV2.Txn("swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
	err := txn.SignAndSend(walletKey, chainID)
	if err != nil {
		panic(err)
	}
	return txn.Hash
}

func SwapETHForExactTokens(amountOut uint, path []common.Address, to common.Address, deadline uint) web3.Hash {
	txn := uniswapRouterV2.Txn("swapETHForExactTokens", amountOut, path, to, deadline)
	err := txn.SignAndSend(walletKey, chainID)
	if err != nil {
		panic(err)
	}
	return txn.Hash
}

func SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn uint, amountOutMin uint, path []common.Address, to common.Address, deadline uint) web3.Hash {
	txn := uniswapRouterV2.Txn("swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
	err := txn.SignAndSend(walletKey, chainID)
	if err != nil {
		panic(err)
	}
	return txn.Hash
}

func SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin uint, path []common.Address, to common.Address, deadline uint) web3.Hash {
	txn := uniswapRouterV2.Txn("swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
	err := txn.SignAndSend(walletKey, chainID)
	if err != nil {
		panic(err)
	}
	return txn.Hash
}

func SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn uint, amountOutMin uint, path []common.Address, to common.Address, deadline uint) web3.Hash {
	txn := uniswapRouterV2.Txn("swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
	err := txn.SignAndSend(walletKey, chainID)
	if err != nil {
		panic(err)
	}
	return txn.Hash
}
