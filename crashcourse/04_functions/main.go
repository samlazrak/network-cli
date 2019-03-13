package main

import "fmt"

// Similar to other languages, going over it for the syntax
// basic functions

// func funcName(argument type) returnType
func greeting(name string) string {
	// concatination, just like javascript
	return "Hello " + name
}

// can remove first int since both variables are ints (num1, num2 int)
func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	//fmt.Println(function()) - this calls the function
	fmt.Println(greeting("Sam"))
	fmt.Println(getSum(3, 4))
}
