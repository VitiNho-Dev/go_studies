package main

import "fmt"

type exemple interface {
	som() int
}

type exempleImpl struct {
	num1 int
	num2 int
}

func (e exempleImpl) som() int {
	result := e.num1 + e.num2

	return result
}

func calculate(e exemple) {
	fmt.Println(e.som())
}

func main() {
	teste := exempleImpl{num1: 1, num2: 1}

	calculate(teste)
}
