package main

import "fmt"

// name:= "golang" => can't declare like this out of function def other ways are possible

func main() {

	// once declared you have to use variables in golang
	// var name string = "golang"

	// infer => Golang can infer type of variable based on its value
	// var firstName = "Tanmay"
	// var age int = 30

	// fmt.Println(name)
	// fmt.Println(firstName)
	// fmt.Println(age)

	// shorthand syntax
	lastName := "golang" // used when we are declaring and assigning a variable at the same time
	fmt.Println(lastName)

	// to declare variable without assigning values you must use this syntax
	// var name string
	// name = "Tanmay"

	// var price float32
	// price = 50.5
	// var price = 50.5
	price := 50.5

	fmt.Println(price)

}
