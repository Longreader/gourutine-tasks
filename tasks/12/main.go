package main

import "fmt"

type Set struct {
	st map[string]struct{}
}

func NewSet(data []string) *Set {
	set := &Set{
		st: make(map[string]struct{}),
	}
	for _, key := range data {
		set.Add(key)
	}
	return set
}

func (s *Set) Add(value string) {
	s.st[value] = struct{}{}
}

func (s *Set) Get(value string) (string, bool) {
	_, ok := s.st[value]
	if ok {
		return value, true
	} else {
		return "", false
	}
}

func (s *Set) Cross(s1 *Set) *Set {
	res_set := NewSet(make([]string, 0))
	for k := range s.st {
		if _, ok := s1.Get(k); ok {
			res_set.Add(k)
		}
	}
	return res_set
}

func main() {

	args := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewSet(args)

	for k := range set.st {
		fmt.Println(k)
	}

}
