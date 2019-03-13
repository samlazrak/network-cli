package main

import "fmt"

func main() {
	// long method
	i := 1
	for i <= 10 {
		//fmt.Println(i)
		i++ // or i = i+1
	}

	// Short method
	for i := 1; i <= 10; i++ {
		//fmt.Printf("" + "number: %d\n", i)
	}

	// FizzBuzz - print numbers 1 to 100 but for every multiples of three print Fizz of the number and every multiples of five print Buzz. For numbers that are multiples of both print FizzBuzz.
	for i := 1; i <= 100; i++ {
		// or if i%15
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
