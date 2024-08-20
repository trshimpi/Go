package main

import (
	"fmt"
	"time"
)

func main() {
	i := 7

	// No need to explicitely write break in each case end , go applies it automatically
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Default")
	}

	// multiple condition switch => , separated values can be used in this case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Holiday")
	default:
		fmt.Println("Workday")
	}

	//type switch => .(type)
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("Other ")
		}
	}

	whoAmI(10.0)
}
