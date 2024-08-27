package main

import "fmt"

// enumerated type
type OrderStatus string

const (
	Received   OrderStatus = "received"
	Confirmed              = "confirmed"
	Prepared               = "prepared"
	Delievered             = "delievered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to:", status)
}

func main() {
	changeOrderStatus(Confirmed)
}
