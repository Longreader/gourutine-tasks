package main

import "fmt"

// Интерфейсный тип определяет поведение других типов
// путем предствления списка методов для реализации
// Таким путем реализуется полиморфизм

type Caller interface {
	Call()
}

func SendSount(c Caller) {
	c.Call()
}

type One struct {
}

func (o *One) Call() {
	fmt.Println("One")
}

type Two struct {
}

func (o *Two) Call() {
	fmt.Println("Two")
}

func main() {
	one := &One{}
	two := &Two{}

	SendSount(one)
	SendSount(two)
}
