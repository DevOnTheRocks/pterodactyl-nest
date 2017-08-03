package main

import (
	"fmt"

	"github.com/cossacksman/nest/api"
)

func main() {
	response := api.GetServers()
	for _, server := range response.Data {
		fmt.Printf("%+v", server)
	}
}
