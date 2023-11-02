package main

import (
	"fmt"
	"reflect"
)

func GetTypeOf(value interface{}) {
	fmt.Println(reflect.ValueOf(value).Type())
}

func main() {

	v1 := 1
	v2 := true
	v3 := "1"
	v4 := make(chan int, 1)
	var v5 interface{} = 5

	GetTypeOf(v1)
	GetTypeOf(v2)
	GetTypeOf(v3)
	GetTypeOf(v4)
	GetTypeOf(v5)

}
