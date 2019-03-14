package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// GET request -- functional
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		// _ is ignoring the error
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	// Creating a map of the json. This is like creating a json object in js.
	jsonData := map[string]string{"firstname": "Sam", "lastname": "Raboy"}
	// Marshalling map into json. _ is ignoring the error
	jsonValue, _ := json.Marshal(jsonData)

	// POST request -- 3rd argument is buffering the json into bytes.
	//response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	// Alternative method of POST request -- formulating request with header info
	// Creates request object ( is that the right term? )

	request, _ := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	// & means I am referencing the memory addr of http, not actually passing the data
	// Created client object, and had it Do the request
	client := &http.Client{}
	response, err = client.Do(request)
	// Handling response
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		// _ is ignoring the error
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
