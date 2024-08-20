package main

import "fmt"

// const age = 30

// name:= "golang" => can't declare like this out of function def

func main() {
	// const name string ="golang"
	// const name = "golang"

	// fmt.Println(age)

	// grouping of constants
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
