package main

import "fmt"

func main() {
	// Arrays - must be a fixed length and must declare type
	// Slices - an Array without a fixed type
	var fruitArr[2]string

	// Assign Values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// [] to access specific indices
	fmt.Println(fruitArr)

	// Declare and assign
	vegetableArr := [2]string{
		"Brocoli",
		"Asparagus",
	}
	fmt.Println(vegetableArr)

	// Slice
	fruitSlice := []string{
		"Apple",
		"Orange",
		"Grape",
	}
	fmt.Println(fruitSlice)

	// appending to slice
	fruitSlice = append(fruitSlice, "Melon")
	fmt.Println(fruitSlice)
	// print length
	fmt.Println(len(fruitSlice))
	// print a range
	fmt.Println(fruitSlice[1:3])
}
