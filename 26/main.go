package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

func main() {
	m := make(map[rune]rune)
	str := "aabcd"
	str = strings.ToLower(str)

	arr := []rune(str)
	for _, v := range arr {
		m[v] = v
	}
	if len(m) < len(arr) {
		fmt.Println("false")
		return
	}
	fmt.Println("true")
}
