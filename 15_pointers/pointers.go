package main

import "fmt"

func changeNumbyVal(num int) {
	num = 2
}

func changeNumbyRef(num *int) {
	*num = 5
}

func ch1(nums []int) {
	nums = append(nums, 4)
}

func ch2(nums *[]int) {
	*nums = append(*nums, 4)
}

func main() {

	num := 10
	fmt.Println("value of num before change", num)
	changeNumbyVal(num)
	fmt.Println("value of num after changebyval", num)
	changeNumbyRef(&num)
	fmt.Println("value of num after changebyref", num)

	nums := []int{1, 2, 3}
	fmt.Println("value of nums before change", nums)
	ch1(nums)
	fmt.Println("value of nums after changebyval", nums)
	ch2(&nums)
	fmt.Println("value of nums after changebyref", nums)

}
