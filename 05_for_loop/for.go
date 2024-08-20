package main

import "fmt"

//for => only construct in go for looping ; no while , dowhile loop available
func main() {

	//while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	//infinite loop
	// for {
	// 	println(1)
	// }

	//classic for loop
	// for i := 0; i <= 3; i++ {
	// 	if i == 2 {
	// 		continue // can work with continue and break as well
	// 		// break
	// 	}
	// 	fmt.Println(i)
	// }

	// 1.22 feature => range
	for i := range 10 { // this will print from 0 to 9
		fmt.Println(i)
	}

}
