package main

import "fmt"

func add(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("nums:", nums, "sum:", sum)
	return sum
}

func main() {

	/*
		Variadic functions can be called with any number of trailing arguments.
		For example, fmt.Println is a common variadic function.
	*/
	add(1, 2, 3)
	add(1, 2, 3, 4)

	nums2 := []int{1, 2, 3, 4, 5}

	//If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	add(nums2...)

}
