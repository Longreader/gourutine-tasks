package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Reader(out chan int, ctx context.Context, wg *sync.WaitGroup) {
	for {
		val, ok := <-out
		if !ok {
			fmt.Println("Reader closed")
			wg.Done()
			return
		}
		fmt.Println("Get value", val)
	}
}

func Writer(in chan int, ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			close(in)
			fmt.Println("Writer closed")
			return
		default:
			val := rand.Intn(100)
			in <- val
		}
	}
}

func main() {

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(time.Second*2, cancel)

	ch := make(chan int)

	wg.Add(1)
	go Writer(ch, ctx, &wg)
	wg.Add(1)
	go Reader(ch, ctx, &wg)

	wg.Wait()

}
