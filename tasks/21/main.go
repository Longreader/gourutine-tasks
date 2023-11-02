package main

import "fmt"

// Паттерн небходим на случай если функции
// находяться в другом модуле и недоступны для изменения

type Agent struct {
}

func (a *Agent) MyNameIs() {
	fmt.Println("My name is classified")
}

type Human struct {
}

func (h *Human) CallMyself() {
	fmt.Println("I call myself Alex")
}

// ------------------------------------------------------

type SpeakAdapter interface {
	Speak()
}

type AgentAdapter struct {
	*Agent
}

func (a *AgentAdapter) Speak() {
	a.MyNameIs()
}

func NewAgentAdapter(a *Agent) SpeakAdapter {
	return &AgentAdapter{a}
}

type HumanAdapter struct {
	*Human
}

func (h *HumanAdapter) Speak() {
	h.CallMyself()
}

func NewHumanAdapter(h *Human) SpeakAdapter {
	return &HumanAdapter{h}
}

func main() {
	people := []SpeakAdapter{&AgentAdapter{&Agent{}}, &HumanAdapter{&Human{}}}
	for _, hum := range people {
		hum.Speak()
	}
}
