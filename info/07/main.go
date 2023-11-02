package main

import "fmt"

func main() {
	dict := make(map[int]int, 0)

	dict[0] = 1
	dict[1] = 343
	dict[2] = 234
	dict[32423432] = 342432

	for k, v := range dict {
		fmt.Println(k, v)
	}
}
