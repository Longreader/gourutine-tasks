package main

import "fmt"

func someAction(v []int8, b int8) []int8 {
	v[0] = 100
	return append(v, b)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	a = someAction(a, 6)
	fmt.Println(a)
}
