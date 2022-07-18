package main

// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором.

import (
	"fmt"

	"24/Point"
)

func main() {
	p1 := Point.NewPoint(1, 1)
	p2 := Point.NewPoint(2, 1)

	fmt.Println(Point.CalcDistance(&p1, &p2))
}
