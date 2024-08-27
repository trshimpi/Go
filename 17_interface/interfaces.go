package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	// gateway should not be of concrete implementation / it should be an interface
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)

	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)
	p.gateway.pay(amount)
}

// no need to use implements keyword as it's implicit ,
//struct should implement all the functions written inside interface
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payment using razorpay:", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Making payment using stripe:", amount)
}

type fakePayment struct{}

func (f fakePayment) pay(amount float32) {
	fmt.Println("making payment using fake testing purpose:", amount)
}

func main() {
	// instead of doing like this we need to create a common interface to avoid constant changes in all the places
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}

	paymenter := fakePayment{}

	newPayment := payment{
		gateway: paymenter,
	}
	newPayment.makePayment(100)
}
