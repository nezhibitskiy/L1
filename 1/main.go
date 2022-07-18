package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в
// структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	field1 int
}

type Action struct {
	Human
	field2 int
}

func main() {
	a := Action{}
	a.field1 = 1
	a.field2 = 2

	fmt.Println(a)
}
