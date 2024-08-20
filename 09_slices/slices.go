package main

import (
	"fmt"
	"slices"
)

// slices => dynamic (no fixed length)
func main() {
	//An uninitialized slice equals to nil and has length 0.
	var nums []int
	fmt.Println("nums ", nums, nums == nil, len(nums))

	// cap => capacity - maximum number of elements can fit , third param -> initial capacity
	// capacity and length autoincrements based on number of elements present in slice
	var a = make([]int, 3, 10)
	fmt.Println("a:", a, "len:", len(a), "cap:", cap(a))

	a = append(a, 1, 2, 3)
	fmt.Println("a after append:", a, "len:", len(a), "cap:", cap(a))

	a = append(a, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("a after second append:", a, "len:", len(a), "cap:", cap(a))

	// to have non-zero values
	var b = make([]int, 0, 5)
	fmt.Println("b :", b, "len:", len(b), "cap:", cap(b))
	b = append(b, 1, 2)
	fmt.Println("b after append ", b)

	c := []int{}
	fmt.Println("C ", c)

	var d = make([]int, len(a))
	copy(d, a) // copy function used to copy whole array
	copy(c, a) // destination slice length should be non zero
	copy(b, a) // gave [0 0] as output , not copied properly just made all present elements zero ,when length of slices are not same
	fmt.Println("d after copy", d)
	fmt.Println("c after copy", c)
	fmt.Println("b after copy", b)

	// slice operator
	var e = []int{1, 2, 3, 4, 5}
	fmt.Println("slice ", e[2:4])
	fmt.Println("slice2", e[:2])
	fmt.Println("Slice3", e[2:])

	t1 := []string{"a", "b", "c"}

	t2 := []string{"a", "b", "c"}

	if slices.Equal(t1, t2) {
		fmt.Println("t1==t2")
	}

	// dimensions need not be same in 2d slices
	var twoD = [][]int{{1, 2}, {1, 2, 3}}
	fmt.Println(twoD)
}
