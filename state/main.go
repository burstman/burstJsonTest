package main

import (
	"fmt"
	"log"
	"test/jsontest"
)

func main() {
	resp, err := jsontest.GetTest()
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range resp {
		fmt.Printf("%v : %v\n", k, v)
	}

}
