package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"main.go/swap"
)

func main() {
	// pk: b6d9645fe31847ea39342d9f9003650888fcf9e4aa9deab1886bdfd7a1bdd227

	swap.ExactInputSingle(common.HexToAddress("0x110a13fc3efe6a245b50102d2d79b3e76125ae83"), common.HexToAddress("0xc778417e063141139fce010982780140aa0cd5ab"), uint(3000), common.HexToAddress("0x1d3C8F8083b96cF4085Cc69f1fD4B471e0ce0f5d"), *big.NewInt(100000000000000), *big.NewInt(100), *big.NewInt(0), *big.NewInt(0))

}
