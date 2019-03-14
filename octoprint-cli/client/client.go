package client

import (
	"bytes"
	"fmt"
	"net/http"
)

func Get(address string) *http.Response {
	response, err := http.Get(address)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	return response
}

func Post(address string, jsonValue []uint8) *http.Response {
	request, _ := http.NewRequest("POST", address, bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	return response
}

func Client() {
}
