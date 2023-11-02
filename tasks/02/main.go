package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	nums := []int{2, 4, 6, 8, 10}

	for _, num := range nums {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i * i)
		}(num)
	}
	wg.Wait()
}
