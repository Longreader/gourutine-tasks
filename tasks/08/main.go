package main

import (
	"fmt"
	"os"
	"strconv"
)

func XOR(X int64, pos uint) int64 {
	pos -= 1
	Y := int64(1)
	Y = Y << pos
	return (X | Y) & ^(X & Y)

}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Enter number of bit")
		return
	}

	pos, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error ocured")
		return
	}

	if pos > 8 || pos < 1 {
		fmt.Println("Bad position for int64")
		return
	}

	u_pos := uint(pos)

	num := int64(64)
	fmt.Printf("Number old: %d\nBy bytes: %b\n", num, num)

	newNum := XOR(num, u_pos)
	fmt.Printf("Number new: %d\nBy bytes: %b\n", newNum, newNum)

}
