package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала —
// читать. По истечению N секунд программа должна завершаться.

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	ch := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		internalCtx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
		i := 0
		for {
			select {
			case <-internalCtx.Done():
				ch <- i
				i++
				internalCtx, cancel = context.WithTimeout(context.Background(), time.Millisecond*500)
			case <-ctx.Done():
				cancel()
				close(ch)
				wg.Done()
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}()
	wg.Wait()
	cancel()
}
