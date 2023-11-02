package main

import (
	"bytes"
	"fmt"
)

func main() {

	data := []string{"First", "Second"}

	// Самый эфективный способ

	total1 := make([]byte, len(data)*10)
	length := 0

	for _, value := range data {
		length += copy(total1[length:], []byte(value))
	}

	fmt.Println(string(total1[:]))

	// Второй по эффективности

	var total2 bytes.Buffer

	for _, value := range data {
		total2.WriteString(value)
	}

	fmt.Println(total2.String())

	// Третий

	var total3 string

	for _, value := range data {
		total3 += value
	}

	fmt.Println(total3)
}
