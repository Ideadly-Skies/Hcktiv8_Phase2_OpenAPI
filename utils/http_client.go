package utils

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"io"
	// "log"
	"net/http"
	// "bytes"
)

var Client = &http.Client{}

func RequestGET(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Check for non-2xx status codes
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("failed to get data, status code: %d", res.StatusCode)
	}	

	return body, nil
}

func RequestPOST(url string, headers map[string]string, dataPayload io.Reader) ([]byte, error) {
	client := &http.Client{}

	// Create a new POST request
	req, err := http.NewRequest("POST", url, dataPayload)
	if err != nil {
		return nil, err
	}

	// Add headers to the request
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Execute the request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Read and return the response body
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
