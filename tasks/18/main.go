package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int64
)

func thread(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		atomic.AddInt64(&counter, 1)
	}
}

func main() {
	var (
		wg sync.WaitGroup
	)

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go thread(&wg)
	}

	wg.Wait()

	fmt.Printf("Counter equal %d\n", counter)
}
