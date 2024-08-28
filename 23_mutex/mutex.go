package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex // good practice to add mutex in struct itself
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	// use mutex to lock the resource to avoid race conditions
	p.mu.Lock()
	p.views += 1
	// best practice is to keep unlock in defer
}

func main() {
	var wg sync.WaitGroup
	myPost := post{
		views: 0,
	}

	// for i := 0; i < 100; i++ {
	// 	myPost.inc()
	// } // 100

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
	fmt.Println(myPost.views)
}
