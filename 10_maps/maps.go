package main

import (
	"fmt"
	"maps"
)

func main() {

	//creating a map using make => map[key]value
	var m = make(map[string]int, 0)

	// setting key-value pair in map
	m["key1"] = 1
	m["key2"] = 2

	fmt.Println("map m:", m)

	//getting value based on key in map
	fmt.Println("m key1 value:", m["key1"])

	//map returns zero value based on data type if key doesn't exits
	fmt.Println("m key3 value:", m["Key3"])

	// get length of map using len , it is equal to number of key-value pairs
	fmt.Println("length of map m:", len(m))

	// to remove any key value pair from map use delete
	delete(m, "key1")

	fmt.Println("map m after delete of key1:", m)

	m["key1"] = 1
	//you are remove all the key-val pairs using clear
	clear(m)
	fmt.Println("m after clear:", m)

	// second param gives true or false , after checking wheather that key exists in map or not
	_, ok := m["key3"]
	if ok {
		fmt.Println("All ok")
	} else {
		fmt.Println("not ok")
	}

	// direct declartion with values
	m1 := map[string]int{"key1": 1, "key2": 2}

	m2 := map[string]int{"key1": 1, "key2": 2}

	// maps give equal function to check equality
	if maps.Equal(m1, m2) {
		fmt.Println("m1==m2")
	}

}
