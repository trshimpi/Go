package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// person struct represents an individual with JSON tags for marshaling/unmarshaling
type person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	// Sample JSON data representing an array of persons
	myJson := ` [
			{
				"first_name": "John",
				"last_name": "Doe",
				"hair_color": "brown",
				"has_dog": true
			},
			{
				"first_name": "Jane",
				"last_name": "Smith",
				"hair_color": "blonde",
				"has_dog": false
			},
			{
				"first_name": "Michael",
				"last_name": "Johnson",
				"hair_color": "black",
				"has_dog": true
			}

	]`

	// Slice to store unmarshalled JSON data
	var unmarshalled []person

	// Unmarshal JSON data into the slice of person structs
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		fmt.Println("Error unmarshalling json", err)
	}

	// Print the unmarshalled data
	fmt.Printf("unmarshalled: %v", unmarshalled)

	// Create JSON from struct

	// Slice to store person structs
	var mySlice []person

	// Create two person instances
	var m1 = person{
		FirstName: "Walley",
		LastName:  "White",
		HairColor: "White",
		HasDog:    true,
	}

	var m2 = person{
		FirstName: "Jessie",
		LastName:  "Pinkman",
		HairColor: "Black",
		HasDog:    false,
	}

	// Append the person instances to the slice
	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	// Marshal the slice of person structs into JSON with indentation
	newJson, err := json.MarshalIndent(mySlice, "", "   ")
	if err != nil {
		fmt.Println("Error marshalling json", err)
	}

	// Print the marshalled JSON as a string
	fmt.Println(string(newJson))

	// Write the marshalled JSON to a file
	os.WriteFile("26_json/json.json", newJson, 0644)

}
