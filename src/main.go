package main

import (
	"fmt"

	"main.go/swap"
)

func main() {

	address := "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
	addr, err := swap.ToChecksumAddress(address)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(addr)

}
