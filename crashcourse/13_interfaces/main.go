package main

import (
	"fmt"
	"math"
)

// interfaces are data types that represent a set of method signatures for structs

// define interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// value reciever
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

// pass the shape defined into the interface then the interface finds a associated func that implements said shape and then does its area function..?
func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{
		x: 0,
		y: 0,
		radius: 5,
	}
	rectangle := Rectangle{
		width: 10,
		height: 5,
	}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))

}
