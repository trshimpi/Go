package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	// run at the end of function completion => defer
	defer w.Done()
	fmt.Println("Doing task:", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		// to run tasks concurrently
		go task(i, &wg)

	}

	wg.Wait()
}
