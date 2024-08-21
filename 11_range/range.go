package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4}

	sum := 0

	// first param is index , if you don't want to use it , use _ instead
	for _, num := range nums {
		sum += num
	}

	// range on slices or arrays will return index and actual number as two param
	for i, num := range nums {
		fmt.Println("i", i, "num:", num)
	}

	fmt.Println("Sum :", sum)

	m := map[string]int{"key1": 1, "key2": 2}

	// range on map returns key and value as two params
	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}

	// key will returned if u r using only one param
	for k := range m {
		fmt.Println(k)
	}

	//range on strings iterates over Unicode code points.
	//The first value is the starting byte index of the rune and the second the rune itself
	for i, c := range "Golang" {
		fmt.Println("i:", i, "c", c)
	}

	for i, c := range "Golang" {
		fmt.Println("i:", i, "c", string(c))
	}
}
