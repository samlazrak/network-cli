package main

import "fmt"

func main() {
	x := 15
	y := 10

	// if - common practice is to not use parenthesis
	// if (x<y) {}
	// can add && for and and || for or
	if x < y {
		// %d is a placeholder -- make sure to use Printf
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "green"

	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color ain't either of em")
	}

	// Switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("It's not blue or red")
	}
}
