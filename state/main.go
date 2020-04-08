package main

import (
	"fmt"
	"log"
	"test/jsontest"
)

func main() {
	//jsonMap := make(map[string]interface{})
	resp, err := jsontest.GetTest()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Application)
	fmt.Println(resp.AvailableProcessors)
	fmt.Println(resp.FreeMemory)
	fmt.Println(resp.GrpcAPIEnabled)
	fmt.Println(resp.GrpcAPIPort)
	fmt.Println(resp.IsScanning)
	fmt.Println(resp.LastBlock)
	fmt.Println(resp.LastBlockchainFeeder)
	fmt.Println(resp.LastBlockchainFeederHeight)
}
