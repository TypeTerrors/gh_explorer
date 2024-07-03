package tools

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var Client = &http.Client{
	Timeout: time.Second * 5,
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxConnsPerHost:     100,
		MaxIdleConnsPerHost: 100,
	},
}

func Get(url string, headers map[string]string) ([]byte, error) {
	maxRetries := 3
	var res []byte
	for i := 0; i < maxRetries; i++ {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("GET request", "Error creating GET request 😭", err)
			return nil, err
		}

		for k, v := range headers {
			req.Header.Add(k, v)
		}

		resp, err := Client.Do(req)
		if err != nil {
			fmt.Println("GET request", "Error making GET request 😭", err)
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("GET request", "Error reading GET request 😭", err)
			return nil, err
		}
		res = body
		break
	}
	return res, nil
}

func Post[T any](url string, headers map[string]string, data interface{}) (*T, error) {
	maxRetries := 3
	var res T
	var clientResponse *http.Response

	for i := 0; i < maxRetries; i++ {

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("POST request", "Error marshalling POST request 😭", err)
			return nil, err
		}

		req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
		if err != nil {
			fmt.Println("POST request", "Error creating POST request 😭", err)
			return nil, err
		}

		for k, v := range headers {
			req.Header.Add(k, v)
		}

		clientResponse, err = Client.Do(req)
		if err != nil {
			fmt.Println("POST request", "Error making POST request 😭", err)
			return nil, err
		}
		defer clientResponse.Body.Close()

		var body []byte
		if clientResponse.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(clientResponse.Body)
			if err != nil {
				fmt.Println("POST request", "Error creating gzip reader 😭", err)
				return nil, err
			}
			defer clientResponse.Body.Close()

			body, err = io.ReadAll(reader)
			if err != nil {
				fmt.Println("POST request", "Error reading POST request 😭", err)
				return nil, err
			}

			err = json.Unmarshal(body, &res)
			if err != nil {
				fmt.Println("POST request", "Error unmarshalling POST request 😭", err)
				return nil, err
			}
		} else {
			body, err = io.ReadAll(clientResponse.Body)
			if err != nil {
				fmt.Println("POST request", "Error reading POST request 😭", err)
				return nil, err
			}

			// Parse the data
			err = json.Unmarshal(body, &res)
			if err != nil {
				fmt.Println("POST request", "Error unmarshalling POST request 😭", err)
				return nil, err
			}
		}
		break
	}
	return &res, nil
}
