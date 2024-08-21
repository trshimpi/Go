package main

import "fmt"

func addition(a int, b int) int {
	return a + b
}

func add3(a, b, c int) int {
	return a + b + c
}

//in  go function  can return multiple values
func multireturn(a int, b int) (int, int) {
	return a, b
}

// we can pass function as param to another function
func processIt(fn func(a int) int) {
	fn(2)
}

// function can return another function as well
func anotherFunc() func(a int) int {
	return func(a int) int {
		return a
	}
}

func main() {

	fmt.Println(addition(3, 4))
	fmt.Println(add3(12, 3, 4))

	a, b := multireturn(2, 3)
	fmt.Println(a, b)

	fn := func(a int) int {
		fmt.Println(a)
		return a
	}

	processIt(fn)

	fn1 := anotherFunc()

	fmt.Println(fn1(10))
}
