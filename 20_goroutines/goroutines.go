package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing task:", id)
}

func main() {
	for i := 0; i <= 10; i++ {
		// to run tasks concurrently
		go task(i)
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	// just to stop program from exit immediately
	time.Sleep(time.Second * 2)
}
