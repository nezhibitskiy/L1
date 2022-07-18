package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	arr1 := []int{5, 4, 3, 6, 8, 3, 5, 78, 9, 14}
	arr2 := []int{6, 5, 9, 4, 3, 51, 5, 98, 31, 3}

	join := make(map[int]int)
	res := make([]int, 0)
	for _, v := range arr1 {
		join[v] = 1
	}
	for _, v := range arr2 {
		if j, ok := join[v]; ok {
			if j == 1 {
				res = append(res, v)
				join[v]++
			}
		}
	}

	fmt.Println(res)
}
