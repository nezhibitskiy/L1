package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func Solution(arr []string) []string {
	if len(arr) < 2 {
		return arr
	}
	m := make(map[string]int)
	res := make([]string, 0)
	for _, v := range arr {
		if _, ok := m[v]; ok {
			continue
		} else {
			res = append(res, v)
			m[v] = 1
		}
	}
	if len(res) == 0 {
		return arr
	}
	return res
}

func main() {
	fmt.Println(Solution([]string{"cat", "cat", "dog", "cat", "tree"}))
}
