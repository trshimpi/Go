package main

import "fmt"

// enumerated type
type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delievered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to:", status)
}

func main() {
	changeOrderStatus(Delievered)
}
