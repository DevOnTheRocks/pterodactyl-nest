package main

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	// "encoding/json"
	"fmt"

	"github.com/cossacksman/nest/network"
)

// Keys - lorem ipsum
type Keys struct {
	Private string `json:"private"`
	Public  string `json:"public"`
}

func main() {
	keys := getKeys()
	uri := "https://panel.gameonthe.rocks/api/admin/users"
	token := auth(uri, keys.Private, keys.Public)
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	response := network.Get(uri, nil, headers)

	fmt.Println(response)

	var users UsersResponse
	json.Unmarshal([]byte(response), &users)

	for _, user := range users.Data {
		fmt.Println(user.Attributes.NameFirst, user.Attributes.NameLast)
	}
}

func getKeys() Keys {
	data, err := ioutil.ReadFile("keys.json")
	if err != nil {
		panic(err)
	}

	var keys Keys
	json.Unmarshal(data, &keys)

	return keys
}

// Auth - Create an auth string from a URI.
func auth(uri string, private string, public string) string {
	hash := hmac.New(sha256.New, []byte(private))
	hash.Write([]byte(uri))
	based := b64.StdEncoding.EncodeToString(hash.Sum(nil))
	output := public + "." + based

	return output
}
