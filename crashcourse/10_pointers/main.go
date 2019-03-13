package main

import "fmt"

func main() {
	// pointer allows to point to the memory address of a value thats in a variable
	a := 5
	// b is a pointer of a
	b := &a
	// b = 0xc000088000 - memory address
	// where the value is stored in memory
	fmt.Println(a, b)
	// type
	fmt.Printf("%T\n", b)
	// b type is *int -- * means pointer
	// int and *int are not the same
	// use * to read val from addr
	fmt.Println(*b) // prints 5
	fmt.Println(*&a) // same
	// change val with pointer
	*b = 10 // Oh
	// b is equal to the MEMORY ADDRESS of a
	// therefore if you change the value that the memory address of a holds to 10 the variable a also equals 10
	fmt.Println(a)
	// why use?
	// might have to pass a lot of data stored in an address which means you can pass the address and not the actual data to increase performance
	// everything in go is passed by value not reference
}
