package main

import (
	"fmt"

	"github.com/cossacksman/nest/api"
)

func main() {
	response := api.GetNodes()
	fmt.Printf("%+v", response)

	for _, nodes := range response {
		fmt.Printf("%+v", nodes)
	}
}
