package main

import (
	"encoding/json"
	"fmt"
	"github.com/samlazrak/octoprint-cli/client"
	"io/ioutil"
)

func main(){
	// Get
	//var resGet = client.Get("https://httpbin.org/get")
	//data, _ := ioutil.ReadAll(response.Body)
	//fmt.Printf("data: %T\n", data)

	jsonData := map[string]string{"firstname": "Sam", "lastname": "Raboy"}
	jsonValue, _ := json.Marshal(jsonData)
	var resPost = client.Post("https://httpbin.org/post", jsonValue)
	fmt.Printf("respost: %T\n", resPost)
	data, _ := ioutil.ReadAll(resPost.Body)
	fmt.Printf("data: %T\n", data)
	fmt.Println(string(data))
}