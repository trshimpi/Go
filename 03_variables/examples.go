package main

import "fmt"

func main() {
	// Print a simple greeting
	fmt.Println("Hello World")

	// Declare and initialize a string variable in one line
	var name string = "John"
	fmt.Println(name)

	// Declare a string variable without initializing
	var name2 string
	// Assign a value to the variable later
	name2 = "Tanmay"
	fmt.Println(name2)

	// Declare and initialize an integer variable
	var age int = 20
	fmt.Println(age)

	// Declare and initialize a boolean variable
	var isCool bool = true
	fmt.Println(isCool)

	// Note: Go supports type inference, so we could also write:
	// name := "John"
	// age := 20
	// isCool := true

	// However, using var is recommended for variable declarations
	// and := is preferred for assignments.
}
