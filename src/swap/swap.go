package swap

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/D-Cous/go-web3"
	"github.com/D-Cous/go-web3/abi"
	"github.com/D-Cous/go-web3/contract"
	"github.com/D-Cous/go-web3/jsonrpc"
	"github.com/D-Cous/go-web3/wallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

//TODO: add a node url
//Web3 HTTP client to send calls to the RPC enpoint
var web3Client *jsonrpc.Client

//TODO: add a sender wallet address
var uniswapRouter02 *contract.Contract

//Returns *wallet.Key which is used to sign transactions
func NewWalletKey(privateKey string) *wallet.Key {
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)

	key := wallet.NewKey(ecdsaPrivateKey)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return key
}

//initialize web3 http client
func InitializeWeb3Client(nodeURL string) {
	client, err := jsonrpc.NewClient(nodeURL)
	if err != nil {
		fmt.Println("Failed to connect to node", err)
		os.Exit(1)
	}
	web3Client = client
}

//initialize uniswap router v2 contract instance
func InitializeUniswapRouter02(senderWalletAddress string) {
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
		os.Exit(1)
	}
	contractInstance := contract.NewContract(uniswapRouter02Address, abi, web3Client)
	//set the from addresss
	contractInstance.SetFrom(web3.HexToAddress(senderWalletAddress))
	uniswapRouter02 = contractInstance
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

func ToChecksumAddress(address string) (string, error) {
	//check that the address is a valid Ethereum address
	re1 := regexp.MustCompile("^(0x)?[0-9a-f]{40}$")
	if !re1.MatchString(address) {
		return "", fmt.Errorf("given address '%s' is not a valid Ethereum Address", address)
	}
	//convert the address to lowercase
	re2 := regexp.MustCompile("/^0x/i")
	address = re2.ReplaceAllString(address, "")

	//convert address to sha3 hash
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(address))
	sum := hasher.Sum(nil)
	addressHash := fmt.Sprintf("%x", sum)
	addressHash = re2.ReplaceAllString(addressHash, "")

	//compile checksum address
	checksumAddress := "0x"

	for i := 0; i < len(address); i++ {
		indexedValue, err := (strconv.ParseInt(string(rune(addressHash[i])), 16, 32))
		if err != nil {
			fmt.Println("Error when parsing addressHash during checksum conversion", err)
			return "", err
		}
		if indexedValue > 7 {
			checksumAddress += strings.ToUpper(string(address[i]))
		} else {
			checksumAddress += string(address[i])
		}
	}
	return checksumAddress, nil
}
