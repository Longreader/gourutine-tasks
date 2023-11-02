package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type Worker struct {
	id int
	q  *Queue
}

type Queue struct {
	ch chan interface{}
}

func NewWorker(id int, queue *Queue) *Worker {
	return &Worker{
		id: id,
		q:  queue,
	}
}

func NewQueue() *Queue {
	return &Queue{
		ch: make(chan interface{}, 1),
	}
}

func (q *Queue) Push(value interface{}) {
	q.ch <- value
}

func (q *Queue) Close() {
	close(q.ch)
}

func (w *Worker) Loop(wg *sync.WaitGroup) {
	for {

		value, ok := <-w.q.ch
		if !ok {
			fmt.Println("Closed commection of worker", w.id)
			wg.Done()
			return
		}

		fmt.Printf("%d Worker is data: %d\n", w.id, value)

	}
}

func generateData() int {
	return rand.Intn(10000)
}

func main() {

	var (
		wg         sync.WaitGroup
		numWorkers int
		err        error
	)

	if len(os.Args) == 2 {
		numWorkers, err = strconv.Atoi(os.Args[1])
		if err != nil || numWorkers <= 0 {
			fmt.Println("Wrong arg number value")
			return
		}
	} else {
		fmt.Println("Wrongs args number")
		return
	}

	workers := make([]*Worker, 0, numWorkers)
	queue := NewQueue()

	for i := 0; i < numWorkers; i++ {
		workers = append(workers, NewWorker(i, queue))
	}

	for _, w := range workers {
		wg.Add(1)
		go w.Loop(&wg)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for {
		num := generateData()

		select {
		case <-sigs:
			// Stop
			queue.Close()
			// time.Sleep(time.Second * 1)
			wg.Wait()
			fmt.Println("Workers closed")
			return
		default:
			queue.Push(num)
		}
		time.Sleep(time.Second * 1)
	}

}
