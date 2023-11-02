package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "snow dog sun"

	list := strings.Split(str, " ")

	for i := len(list) - 1; i >= 0; i-- {
		fmt.Printf("%s ", list[i])
	}
}
