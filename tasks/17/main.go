package main

import "fmt"

func binarySearch(array []int, x int, left int, right int) int {
	if right <= left {
		return -1
	}
	mid := (left + right) / 2
	if array[mid] == x {
		return mid
	} else if x < array[mid] {
		return binarySearch(array, x, left, mid)
	} else {
		return binarySearch(array, x, mid+1, right)
	}
}

func main() {

	array := []int{543, 57, 37, 5, 74, 56, 4, 63, 56, 46, 543, 32, 6, 4, 7, 6, 2, 64, 72, 7544, 1, 534, 3, 155, 5, 7, 6, 8, 4, 2}
	x := 72
	index := binarySearch(array, x, 0, len(array))
	fmt.Println(index)
}
