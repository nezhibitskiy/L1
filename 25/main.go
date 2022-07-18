package main

// Реализовать собственную функцию sleep

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	seconds := 5

	start := make(chan struct{})
	wg := new(sync.WaitGroup)
	wg.Add(seconds)
	go func() {
		for {
			select {
			case <-start:
				sleep(seconds, wg)
			}
		}
	}()
	fmt.Println("Wait")
	start <- struct{}{}
	wg.Wait()
	fmt.Println("Done")
}

func sleep(n int, wg *sync.WaitGroup) {
	t := time.Now()
	for i := n; i >= 0; i-- {
		fmt.Printf("t: %d\n", int(time.Since(t).Seconds()))
		<-time.After(1 * time.Second)
		wg.Done()
	}
}
