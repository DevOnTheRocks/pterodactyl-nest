package network

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cossacksman/nest/config"
)

// Get - A method to perform GET requests.
func Get(endpoint string, args []string, headers map[string]string) []byte {
	uri := config.GetConfig().Domain + endpoint
	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

}

// GetAuthHeaders - Create an auth string from a URI.
func GetAuthHeaders(endpoint string) map[string]string {
	cfg := config.GetConfig()
	configuration := config.GetConfig()
	uri := configuration.Domain + endpoint
	hash := hmac.New(sha256.New, []byte(cfg.PrivateKey))
	hash.Write([]byte(uri))
	based := b64.StdEncoding.EncodeToString(hash.Sum(nil))
	output := cfg.PublicKey + "." + based

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", output),
	}

	return headers
}
