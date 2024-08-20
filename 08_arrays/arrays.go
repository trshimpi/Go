package main

import "fmt"

func main() {

	// go stores zeroed values by default if we create empty array
	// int => 0 , string => "" , bool => false
	var nums [4]int
	var strings [4]string
	var bools [4]bool

	fmt.Println(nums)    //[0 0 0 0]
	fmt.Println(strings) //[   ]
	fmt.Println(bools)   //[false false false false]

	//length of an array
	fmt.Println(len(nums))

	//assigning and getting values from array
	nums[0] = 1
	fmt.Println(nums[0])

	//to declare array in single line
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)

	//2-d array
	arr2d := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(arr2d)
	fmt.Println(arr2d[1])
	fmt.Println(arr2d[1][1])

	// - use array when size is already known (fixed size)
	// - memory optmization
	// - constant time access
}
