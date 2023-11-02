package main

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func() {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println(slice)
	}()
	fmt.Println(slice)
}
