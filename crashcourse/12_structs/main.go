package main

import "fmt"

// Structs
// Structs are sort of like classes
// different kinds of methods:
// pointer reciever - for methods that change things
// value reciever - methods that do calculations and do not change things

// define person struct
type Person struct {
	firstName string
	lastName string
	age int
}

// struct methods (value reciever)
// p, person, or anything else
func(person Person) greet() string {
	// p/person is similar to this.
	return "hello my name is " +person.firstName
	// how to call
	// Person.greet(person)
	// or
	// person.greet()
	// similar to a class/object
}
// hasBirthday method ( pointer reciever ) because we are changing something
// * means the pointer
func(person *Person) hasBirthday() {
	// no return, only changes
	person.age++
}

func (person *Person) getMarried(spouseLastName string) {
	person.lastName = spouseLastName
}

func main() {
	// Init person using struct
	person := Person{
		firstName:"Sam",
		lastName:"Lazrak",
		age:22,
	}
	// Similar to javascript, can use person as the temp holder of the value in a function/method

	// Though perhaps pointers change this?
	fmt.Println(person)
	fmt.Println(person.firstName)
	// calling struct method
	fmt.Println(Person.greet(person))
	// or?
	fmt.Println(person.greet())
	person.hasBirthday()
	fmt.Println(person.age)
	person.getMarried("lastname")
	fmt.Println(person.lastName)
}
