package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

type counter struct {
	mu    sync.Mutex
	value int32
}

//func (c *counter) Add(value int32, wg *sync.WaitGroup) {
//	atomic.AddInt32(&c.value, value)
//	wg.Done()
//}

func (c *counter) Add(value int32, wg *sync.WaitGroup) {
	c.mu.Lock()
	c.value += value
	c.mu.Unlock()
	wg.Done()
}

func main() {
	var c counter
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go c.Add(int32(i), &wg)
	}
	wg.Wait()
	fmt.Println(c.value)
}
