package main

import (
	"fmt"

	"github.com/trshimpi/myNiceProgram/helper"
)

func main() {
	helper.PrintHello()

	var myVar helper.SomeType
	myVar.TypeName = "Some Type"
	myVar.TypeNumber = 10
	fmt.Println(myVar)
}
