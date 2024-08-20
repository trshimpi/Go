package main

import "fmt"

func main() {
	// age := 10
	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is teenager")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	// var role = "admin"
	// var hasPermissions = true

	// // logical operators || , && ,etc.
	// if role == "admin" || hasPermissions {
	// 	fmt.Println("Yes")
	// }

	// we can declare a variable inside if construct
	if age := 19; age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 {
		fmt.Println("Person is teenager")
	} else {
		fmt.Println("Person is not an adult")
	}

	// go does not have ternary , you will have to use normal if else
}
