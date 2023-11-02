package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {

	// via context

	ch1 := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Done writing")
				close(ch1)
				return
			default:
				ch1 <- rand.Intn(10)
				fmt.Println("Goroutine 1 works fine, send data")
			}
			time.Sleep(time.Second * 1)
		}
	}(ctx)

	go func(cancel context.CancelFunc) {
		time.Sleep(time.Second * 3)
		cancel()
	}(cancel)

	// via channel

	ch2 := make(chan int)

	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Done writing")
				close(ch2)
				return
			default:
				ch2 <- rand.Intn(10)
				fmt.Println("Goroutine 2 works fine, send data")
			}
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		time.Sleep(time.Second * 3)
		quit <- struct{}{}
	}()

	for ch1 != nil || ch2 != nil {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Printf("Get data from one, %d \n", v)

		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Printf("Get data from two, %d \n", v)

		}
	}

	fmt.Println("Goroutins closed.. exit")

}
