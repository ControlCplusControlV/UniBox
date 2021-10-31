package main

import (
	"fmt"
	"github.com/senseyeio/roger"
)

func Getdata() {
	fmt.Println("cool")
}

func getTokenBalance() {
	rClient, err := roger.NewRClient("127.0.0.1", 6311)
	if err != nil {
		fmt.Println("Failed to connect")
		return
	}

	value, err := rClient.Eval("pi")
	if err != nil {
		fmt.Println("Command failed: " + err.Error())
	} else {
		fmt.Println(value) // 3.141592653589793
	}

	helloWorld, _ := rClient.Eval("as.character('Hello World')")
	fmt.Println(helloWorld) // Hello World

	arrChan := rClient.Evaluate("Sys.sleep(5); c(1,1)")
	arrResponse := <-arrChan
	arr, _ := arrResponse.GetResultObject()
	fmt.Println(arr) // [1, 1]
}
