package main

import (
	"fmt"

	"github.com/senseyeio/roger"
)

func getFactoryStats() {
	rClient, err := roger.NewRClient("127.0.0.1", 6311)
	if err != nil {
		fmt.Println("Failed to connect")
		return
	}
	sesh, _ := rClient.GetSession()

	var rCmd string = "factory_stats_v2()"

	UserlpPositions, err := sesh.Eval(rCmd)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(UserlpPositions)
	}
	sesh.Close() // Cleanup when done
}

func main() {
	getFactoryStats()
}
