package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	st map[int]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		st: make(map[int]int),
	}
}

func (sm *SafeMap) Set(key, value int) {
	sm.Lock()
	defer sm.Unlock()
	sm.st[key] = value
}

func (sm *SafeMap) Get(key int) (int, bool) {
	sm.RLock()
	defer sm.RUnlock()
	value, ok := sm.st[key]
	return value, ok
}

func main() {

	var wg sync.WaitGroup

	storage := NewSafeMap()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 20; i++ {
			storage.Set(i, i+1)
			fmt.Printf("Set value %d and key %d", i+1, i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			val, ok := storage.Get(i)
			if ok {
				fmt.Println("Get value", val)
			} else {
				fmt.Println("No such key")
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 10; i <= 20; i++ {
			val, ok := storage.Get(i)
			if ok {
				fmt.Println("Get value", val)
			} else {
				fmt.Println("No such key")
			}
		}
	}()

	wg.Wait()

}
