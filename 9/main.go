package main

import "fmt"

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		array := []int{1, 2, 3, 4, 5}
		defer close(ch1)
		for _, v := range array {
			ch1 <- v
		}
	}()

	go func() {
		defer close(ch2)
		for v := range ch1 {
			ch2 <- v * 2
		}
	}()

	for v := range ch2 {
		fmt.Println(v)
	}
}
