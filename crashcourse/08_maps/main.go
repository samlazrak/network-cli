package main

import "fmt"

// map is a key value pair
func main() {
	// define map - map of firt names for the key and email for the value
	// map[dataType for key] dataType for value
	emails := make(map[string]string)

	// cleaner declaration
	emails2 := map[string]string {
		"Bob": "bob@gmail.com",
		"Sharon": "bob@gmail.com",
		"Waffle": "bob@gmail.com",
	}

	// Assign key values
	emails["Jimbo"] = "Jimbo@gmail.com"
	emails["Bob"] = "Bob@gmail.com"
	emails["Sharon"] = "Sharon@gmail.com"

	// output - map[Bob:Bob@gmail.com Jimbo:Jimbo@gmail.com Sharon:Sharon@gmail.com]
	fmt.Println(emails)
	// by key
	fmt.Println(emails["Bob"])
	// length
	fmt.Println(len(emails))
	// delete - (map, key)
	delete(emails, "Bob")
	// output
	fmt.Println(emails)
	fmt.Println(emails2)
}
