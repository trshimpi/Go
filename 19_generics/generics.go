package main

import "fmt"

// generics are introduced in go 1.18
// generics are used when we want to generalize a function to use multiple data types as params
// we can use any , interface{} , comparable or with pipe symbol multiple data types as generic type

// func printSlice[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printSlice[T int | string | bool](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type Stack[T any] struct {
	elements []T
}

func main() {

	// nums := []int{1, 2, 3}
	// names := []string{"js", "ts", "go"}
	bools := []bool{true, false, true}
	printSlice(bools)

	myStack := Stack[string]{
		elements: []string{"js", "go"},
	}
	myStack2 := Stack[int]{
		elements: []int{1, 2, 3},
	}
	fmt.Println(myStack)
	fmt.Println(myStack2)

}
