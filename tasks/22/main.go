package main

import (
	"fmt"
	"math/big"
)

func sum(x, y *big.Int) *big.Int {
	return new(big.Int).Add(x, y)
}

func sub(x, y *big.Int) *big.Int {
	return new(big.Int).Sub(x, y)
}

func div(x, y *big.Int) *big.Int {
	return new(big.Int).Div(x, y)
}

func mul(x, y *big.Int) *big.Int {
	return new(big.Int).Mul(x, y)
}

func main() {
	firstValue := big.NewInt(int64(2 << 32))
	secondValue := big.NewInt(int64(2 << 31))

	fmt.Printf("First value is %v\n", firstValue)
	fmt.Printf("Second value is %v\n", secondValue)
	fmt.Printf("Sum is %v\n", sum(firstValue, secondValue))
	fmt.Printf("Sub is %v\n", sub(firstValue, secondValue))
	fmt.Printf("Mul is %v\n", mul(firstValue, secondValue))
	fmt.Printf("Div is %v\n", div(firstValue, secondValue))
}
