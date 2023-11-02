package main

import (
	"fmt"
)

func main() {
	str := "главрыба"

	rune_str := []rune(str)

	// Добавлена в пакете slices Go 1.21
	// reverse_rune_str := slices.Reverse(rune_str)

	for i, j := 0, len(rune_str)-1; i < j; i, j = i+1, j-1 {
		rune_str[i], rune_str[j] = rune_str[j], rune_str[i]
	}

	reverse_str := string(rune_str)

	fmt.Printf("Was %s\nNow %s\n", str, reverse_str)
}
