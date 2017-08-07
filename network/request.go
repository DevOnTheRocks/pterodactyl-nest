package network

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cossacksman/nest/config"
)

// Get - A method to perform GET requests.
func Get(endpoint string) []byte {
	key, value := GetAuthHeader(endpoint, "")
	uri := config.GetConfig().Domain + endpoint
	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(key, value)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if output, err := ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
	} else {
		return output
	}
}

// Post - A method to perform POST requests.
func Post(endpoint string, body string) []byte {
	uri := config.GetConfig().Domain + endpoint
	key, value := GetAuthHeader(endpoint, body)
	client := &http.Client{}

	req, err := http.NewRequest("POST", uri, strings.NewReader(body))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set(key, value)

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if output, err := ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
	} else {
		return output
	}
}

// Put - A method to perform PUT requests.
func Put(endpoint string, body string, headers map[string]string) {
	// TODO:- Stub.
}

// GetAuthHeader - Create an auth string from a URI.
func GetAuthHeader(endpoint string, body string) (string, string) {
	cfg := config.GetConfig()
	uri := cfg.Domain + endpoint

	hash := hmac.New(sha256.New, []byte(cfg.PrivateKey))
	hash.Write([]byte(uri + body))
	based := b64.StdEncoding.EncodeToString(hash.Sum(nil))
	output := cfg.PublicKey + "." + based

	key := "Authorization"
	value := fmt.Sprintf("Bearer %s", output)

	return key, value
}
