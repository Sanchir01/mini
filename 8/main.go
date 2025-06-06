package main

import (
	"sync"
	"sync/atomic"
)

type CustomWaitGroup struct {
	counter int64
	sema    chan struct{}
	done    chan struct{}
	mu      sync.Mutex
	waiting bool
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		sema: make(chan struct{}, 1),
		done: make(chan struct{}),
	}
}

func (cwg *CustomWaitGroup) Add(delta int) {
	newCount := atomic.AddInt64(&cwg.counter, int64(delta))

	if newCount < 0 {
		panic("CustomWaitGroup: negative WaitGroup counter")
	}

	if newCount == 0 {
		cwg.mu.Lock()
		if cwg.waiting {
			select {
			case cwg.done <- struct{}{}:
			default:

			}
		}
		cwg.mu.Unlock()
	}
}

func (cwg *CustomWaitGroup) Done() {
	cwg.Add(-1)
}

func (cwg *CustomWaitGroup) Wait() {
	if atomic.LoadInt64(&cwg.counter) == 0 {
		return
	}

	cwg.mu.Lock()

	if atomic.LoadInt64(&cwg.counter) == 0 {
		cwg.mu.Unlock()
		return
	}

	cwg.waiting = true
	cwg.mu.Unlock()

	<-cwg.done

	cwg.mu.Lock()
	cwg.waiting = false
	cwg.mu.Unlock()
}

func (cwg *CustomWaitGroup) Count() int64 {
	return atomic.LoadInt64(&cwg.counter)
}

type SemaphoreWaitGroup struct {
	sema chan struct{}
	done chan struct{}
	mu   sync.Mutex
}

func NewSemaphoreWaitGroup() *SemaphoreWaitGroup {
	return &SemaphoreWaitGroup{
		sema: make(chan struct{}, 0),
		done: make(chan struct{}, 1),
	}
}

func (swg *SemaphoreWaitGroup) Add(delta int) {
	swg.mu.Lock()
	defer swg.mu.Unlock()

	if delta > 0 {
		newSema := make(chan struct{}, cap(swg.sema)+delta)
		for len(swg.sema) > 0 {
			<-swg.sema
			newSema <- struct{}{}
		}
		swg.sema = newSema
	}
	if delta < 0 {
		for i := 0; i < -delta; i++ {
			select {
			case <-swg.sema:
			default:
				panic("SemaphoreWaitGroup: negative WaitGroup counter")
			}
		}

		if len(swg.sema) == 0 {
			select {
			case swg.done <- struct{}{}:
			default:
			}
		}
	}
}

func (swg *SemaphoreWaitGroup) Done() {
	swg.Add(-1)
}

func (swg *SemaphoreWaitGroup) Wait() {
	swg.mu.Lock()
	if len(swg.sema) == 0 {
		swg.mu.Unlock()
		return
	}
	swg.mu.Unlock()

	<-swg.done
}

func main() {

}
