package main

import "fmt"

// go supports anonymous functions

func adder() func(int) int{
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// So I created a function that has an anonymous function within it allowing me to assign other things, like variables, to that function.
// then it returns another anonymous function which takes the sum and += and returns the sum
// THIS IS FUNCTIONAL PROGRAMMING!!
func main() {
	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}