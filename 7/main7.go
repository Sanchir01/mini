package main

import (
	"context"
	"sync"
)

func main() {}

func fanin(ctx context.Context, chans []chan int) chan int {

	out := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for _, ch := range chans {
			wg.Add(1)
			go func(ch chan int) {
				defer wg.Done()
				for {
					select {
					case v, ok := <-ch:
						if !ok {
							return
						}
						select {
						case out <- v:
						case <-ctx.Done():
							return
						}
					case <-ctx.Done():
						return
					}
				}
			}(ch)
		}
		wg.Wait()
		close(out)
	}()
	return out
}
