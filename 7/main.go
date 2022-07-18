package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	//mu := sync.Mutex{}

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			//mu.Lock()
			m[1] = i
			//mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(m[1])
}
