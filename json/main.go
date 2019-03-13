package main

import "fmt"

type Book struct {
	Title string
	Author string
}

func main() {
	book := Book{Title: "Title", Author: "Author"}

	fmt.Printf("%+v\n", book)
}
