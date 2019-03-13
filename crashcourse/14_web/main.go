package main

import (
	"fmt"
	"net/http"
)

// two special parameters: w & http.ResponseWriter, and r *http.Request pointer
func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Hello world!</h1>")
}

func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)


	fmt.Println("Starting server on port 3000")
	http.ListenAndServe(":3000", nil)
}

// golang great for the web