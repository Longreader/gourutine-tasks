package main

import "fmt"

func main() {
	a := 15
	b := 16
	fmt.Printf("A is %d,B is %d\n", a, b)

	a, b = b, a
	fmt.Printf("A is %d,B is %d\n", a, b)

}
