package main

import (
	"fmt"
	"sync"
)

// post represents a blog post with a view count and a mutex for safe concurrent access
type post struct {
	views int
	mu    sync.Mutex // good practice to add mutex in struct itself
}

// inc increments the view count of the post
// It uses a mutex to ensure thread-safe access to the views field
func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock() // Unlock the mutex when the function returns
		wg.Done()     // Notify the WaitGroup that this goroutine is done
	}()

	p.mu.Lock() // Lock the mutex before modifying the views field
	p.views += 1
	// The unlock is deferred, which is a best practice to ensure it always happens
}

func main() {
	var wg sync.WaitGroup
	myPost := post{
		views: 0,
	}

	// Commented out code for non-concurrent increment
	// for i := 0; i < 100; i++ {
	// 	myPost.inc()
	// } // 100

	// Create 1000 goroutines to increment the view count concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)          // Increment the WaitGroup counter before starting each goroutine
		go myPost.inc(&wg) // Start a new goroutine to increment the view count
	}
	// The main goroutine will wait for all increments to complete after this loop

	wg.Wait()                 // Wait for all goroutines to finish
	fmt.Println(myPost.views) // Print the final view count
}
