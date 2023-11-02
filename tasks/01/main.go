package main

import (
	"fmt"
)

type Human struct {
}

func (h *Human) Hi() {
	fmt.Println("Human say hi")
}

type Action struct {
	Human
}

func (a *Action) Hello() {
	fmt.Println("Action say hello")
}

// Переопределение метода "родительской структуры"
// func (a *Action) Hi() {
// 	fmt.Println("Action say hi")
// }

func main() {

	human := &Human{}
	action := &Action{}

	human.Hi()

	// Если добавить переопределение метода
	// то обращение к "родительской" структуре доступно
	// через следующую форму
	// action.Human.Hi()
	action.Hi()
	action.Hello()
}
