package main

import "fmt"

func remove(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Array contain %v\n", array)
	array = remove(array, 3)
	fmt.Printf("Array contain %v\n", array)
}
