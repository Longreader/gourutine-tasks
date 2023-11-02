package main

import "fmt"

func main() {
	a := new([]int)     // *[]int
	b := make([]int, 0) // []int

	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", &b)

	*a = append(*a, 2)
	b = append(b, 2)

	fmt.Printf("%v\n", *a)
	fmt.Printf("%v\n", b)

	// make алоцирует память для каналов/слайсов/хеш-таблиц
	// new алоцирует память для элемента и возвращает указатель на ячейку памяти
}
