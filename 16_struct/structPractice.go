package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	email string
}

// struct embedding is also possible
type order struct {
	id        string
	price     float32
	status    string
	createdAt time.Time
	customer
}

// if you don't pass pointer over here , it won't change actual value
func (o *order) changePrice(price float32) {
	o.price = price // no need to derefrence pointers over here struct does it automatically
}

func (o order) getPrice() float32 {
	return o.price
}

func main() {

	myOrder := order{
		id:        "abc",
		price:     50.50,
		status:    "received",
		createdAt: time.Now(),
	}

	myOrder2 := order{
		id:    "xyz",
		price: 110.50,
		customer: customer{
			"sahil",
			"srshimpi@gmail.com",
		},
	}

	// fmt.Println(myOrder)
	// fmt.Println(myOrder.getPrice())
	// myOrder.changePrice(100.00)
	// fmt.Println(myOrder)

	// fmt.Println(myOrder2)

	// // define and assign values to struct directly
	// customer := struct {
	// 	id   string
	// 	name string
	// }{
	// 	"123", "tanmay",
	// }

	// fmt.Println(customer)
	myOrder.customer.name = "Tanmay"
	myOrder.customer.email = "trshimpi25@gmail.com"

	fmt.Println(myOrder.customer)
	fmt.Println(myOrder2)
}
