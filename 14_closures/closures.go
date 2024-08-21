package main

import "fmt"

/*

Go supports anonymous functions, which can form closures.
Anonymous functions are useful when you want to define a function inline without having to name it.
*/
func closureFun() func() {
	i := 0
	return func() {
		i++
		fmt.Println(i)
	}
}

func add(a int) func(b int) func(c int) int {
	return func(b int) func(c int) int {
		return func(c int) int {
			return a + b + c
		}
	}
}
func main() {
	b := closureFun()
	b()
	b()
	b()

	c := closureFun()
	c()

	sum := add(1)(2)(3)

	fmt.Println(sum)

}
