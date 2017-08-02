package network

import "net/http"
import "io/ioutil"

// Get - A method to perform GET requests.
func Get(uri string, args []string, headers map[string]string) string {
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
		return string(output)
	}
}

// Post - A method to perform POST requests.
func Post() {

}
