package main

import (
	"fmt"
)

func countSquare(c chan<- int, nums []int) {
	for _, i := range nums {
		c <- i * i
	}
	close(c)
}

func main() {

	c := make(chan int)
	nums := []int{2, 4, 6, 8, 10}
	total := 0
	go countSquare(c, nums)
	for num := range c {
		total += num
	}
	fmt.Println(total)
}
