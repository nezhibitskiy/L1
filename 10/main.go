package main

import (
	"fmt"
	"math"
)

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	//sort.Slice(arr, func(i, j int) bool {
	//	return arr[i] < arr[j]
	//})

	groups := make(map[float64][]float64)

	var g float64
	for _, v := range arr {
		g = math.Trunc(v/10) * 10
		groups[g] = append(groups[g], v)
	}

	for g, t := range groups {
		fmt.Printf("%v: %v\n", g, t)
	}
}
