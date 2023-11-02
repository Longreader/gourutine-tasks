package main

import "fmt"

func generator(data []int, out chan<- int) {
	for _, i := range data {
		out <- i
	}
	close(out)
}

func translator(in, out chan int) {
	for value := range in {
		out <- value * value
	}
	close(out)
}

func acceptor(in <-chan int) {
	for value := range in {
		fmt.Printf("Currant value is %d\n", value)
	}
}

func main() {
	Data := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go generator(Data, ch1)
	go translator(ch1, ch2)
	// Последняя функция конвеера не дает основному потоку завершиться
	// так как не исполняется как гоурутина
	acceptor(ch2)
}
