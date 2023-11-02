package main

import (
	"fmt"
	"sync"
)

type Set struct {
	sync.RWMutex
	st map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		st: make(map[int]struct{}),
	}
}

func (s *Set) Add(value int) {
	s.st[value] = struct{}{}
}

func (s *Set) Get(value int) (int, bool) {
	_, ok := s.st[value]
	if ok {
		return value, true
	} else {
		return 0, false
	}
}

func (s *Set) Cross(s1 *Set) *Set {
	res_set := NewSet()
	for k := range s.st {
		if _, ok := s1.Get(k); ok {
			res_set.Add(k)
		}
	}
	return res_set
}

func (s *Set) Elems() []int {
	result := make([]int, 0)
	for k := range s.st {
		result = append(result, k)
	}
	return result
}

func (s *Set) FullSet(wg *sync.WaitGroup, left, right int) {
	defer wg.Done()
	for i := left; i < right; i++ {
		s.Lock()
		s.Add(i)
		s.Unlock()
	}
}

func main() {

	set1 := NewSet()
	set2 := NewSet()

	var wg sync.WaitGroup

	wg.Add(1)
	go set1.FullSet(&wg, 0, 10)

	wg.Add(1)
	go set2.FullSet(&wg, 5, 15)

	wg.Wait()

	set3 := set1.Cross(set2)
	fmt.Println(set3.Elems())
}
