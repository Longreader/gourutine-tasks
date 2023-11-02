package main

import (
	"fmt"
	"reflect"
)

func main() {
	one := []int{}

	two := make([]int, 0)

	fmt.Println(reflect.TypeOf(one))
	fmt.Println(reflect.TypeOf(two))

}
