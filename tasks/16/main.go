package main

import (
	"fmt"
	"math/rand"
)

func partition(array []int, pivot int) ([]int, []int, []int) {
	var left []int
	var center []int
	var right []int

	for _, x := range array {
		if x < pivot {
			left = append(left, x)
		} else if x == pivot {
			center = append(center, x)
		} else {
			right = append(right, x)
		}
	}

	return left, center, right
}

func quicksort(array []int) []int {
	if len(array) < 2 {
		return array
	} else {
		pivot := array[rand.Intn(len(array))]
		left, center, right := partition(array, pivot)
		return append(append(quicksort(left), center...), quicksort(right)...)
	}
}

func main() {
	array := []int{543, 57, 37, 5, 74, 56, 4, 63, 56, 46, 543, 32, 6, 4, 7, 6, 2, 64, 72, 7544, 1, 534, 3, 155, 5, 7, 6, 8, 4, 2}
	sorted := quicksort(array)
	fmt.Println(sorted)
}
