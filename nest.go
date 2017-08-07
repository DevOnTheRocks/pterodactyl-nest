package main

import (
	"fmt"

	"github.com/cossacksman/nest/api"
)

func main() {
	response := api.TogglePower(1, api.Stop)
	// server := api.GetServer(1)
	fmt.Printf("\n%+v", response)
}
