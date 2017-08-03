package config

import (
	"encoding/json"
	"io/ioutil"
)

var configuration Config

// Config - Application config.
type Config struct {
	PrivateKey string `json:"private"`
	PublicKey  string `json:"public"`
	Domain     string `json:"domain"`
}

// GetConfig - Sets, if necessary, and gets
func GetConfig() Config {
	if configuration == (Config{}) {
		data, err := ioutil.ReadFile("config.json")
		if err != nil {
			panic(err)
		}

		json.Unmarshal(data, &configuration)
	}

	return configuration
}
