package main

import (
	"fmt"
	"time"
)

// send
// func processNow(myNum chan int) {
// 	for num := range myNum {
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second)
// 	}

// }

// receive
// func sum(result chan int, num1 int, num2 int) {
// 	sum := num1 + num2
// 	result <- sum
// }

// goroutine synchronizer
// func task(done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()
// 	fmt.Println("Task processing...")
// }

// email sender
// <-chan - receive only channel , chan<- - send only channel
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() {
		done <- true
	}()

	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {
	// myChannel := make(chan int)

	// myChannel <- 1  // channels are blocking - till second side is ready for receiving

	// fmt.Println("channel value:", <-myChannel) // will give fatal error - deadlock

	//======
	// myNum := make(chan int)

	// go processNow(myNum)

	// for {
	// 	myNum <- rand.Intn(100)
	// }

	// time.Sleep(time.Second)

	//======

	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <-result // blocking

	// fmt.Println("result:", res)

	//============

	// done := make(chan bool)

	// go task(done)

	// <-done

	//======

	// emailChannel := make(chan string, 100)

	// emailChannel <- "1@email.com"
	// emailChannel <- "2@gmail.com"

	// fmt.Println(<-emailChannel)
	// fmt.Println(<-emailChannel)

	//=========

	// email sender
	// buffered channel with space 5
	// emailChan := make(chan string, 5)

	// done := make(chan bool)

	// go emailSender(emailChan, done)
	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("done sending...")

	// // closing buffered channels is imp
	// close(emailChan)

	// <-done

	//===========================

	// chan1 := make(chan int)
	// chan2 := make(chan string)

	// go func() {
	// 	chan1 <- 10
	// }()

	// go func() {
	// 	chan2 <- "pong"
	// }()

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case chan1Val := <-chan1:
	// 		fmt.Println("received data from chan1:", chan1Val)
	// 	case chan2Val := <-chan2:
	// 		fmt.Println("received data from chan1:", chan2Val)
	// 	}
	// }

}
