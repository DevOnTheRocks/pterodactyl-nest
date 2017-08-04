package network

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cossacksman/nest/config"
)

// Get - A method to perform GET requests.
func Get(endpoint string, args []string, headers map[string]string) []byte {
	uri := config.GetConfig().Domain + endpoint
	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

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
func Post() {
	// uri := config.GetConfig().Domain + endpoint
	// req, err := http.NewRequest("GET", uri, nil)

}

// Put - A method to perform PUT requests.
func Put(endpoint string, body interface{}, headers map[string]string) []byte {
	uri := config.GetConfig().Domain + endpoint
	data, err := json.Marshal(body)
	client := &http.Client{}

	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", uri, strings.NewReader(string(data)))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

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

// GetAuthHeaders - Create an auth string from a URI.
func GetAuthHeaders(endpoint string) map[string]string {
	cfg := config.GetConfig()
	configuration := config.GetConfig()
	uri := configuration.Domain + endpoint

	fmt.Printf("Contacting endpoint: %s \n", uri)

	hash := hmac.New(sha256.New, []byte(cfg.PrivateKey))
	hash.Write([]byte(uri))
	based := b64.StdEncoding.EncodeToString(hash.Sum(nil))
	output := cfg.PublicKey + "." + based

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", output),
	}

	return headers
}
