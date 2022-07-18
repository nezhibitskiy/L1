package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func calcValues(arr []int) chan int {
	ch := make(chan int)

	go func(ch chan int) {
		wg := sync.WaitGroup{}
		for _, v := range arr {
			wg.Add(1)
			go func(ch chan int, v int) {
				ch <- v * v
				wg.Done()
			}(ch, v)
		}
		wg.Wait()
		close(ch)
	}(ch)

	return ch
}

func main() {
	arr := []int{2, 4, 6, 8, 10}

	ch := calcValues(arr)

	var sum int
	for v := range ch {
		sum += v
	}
	fmt.Println(sum)
}
