package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func main() {
	data := "snow dog sun"

	arr := strings.Split(data, " ")

	for i := len(arr); i > 0; i-- {
		fmt.Print(arr[i-1], " ")
	}
}
