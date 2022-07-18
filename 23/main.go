package main

// Удалить i-ый элемент из слайса.

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	i := 1
	arr = append(arr[:i], arr[i+1:]...)
	fmt.Println(arr)
}
