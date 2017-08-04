package main

import "github.com/cossacksman/nest/api"
import "fmt"

func main() {
	// api.TogglePower(1, api.Stop)
	server := api.GetServer(11)
	fmt.Printf("%+v", server.Data.Attributes)
}
