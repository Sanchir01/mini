package main

import (
	"testing"
	"time"
)

func TestCustomWaitGroup(t *testing.T) {
	t.Run("базовая функциональность", func(t *testing.T) {
		cwg := NewCustomWaitGroup()
		completed := make(chan int, 5)

		cwg.Add(5)
		for i := 0; i < 5; i++ {
			go func(id int) {
				defer cwg.Done()
				time.Sleep(10 * time.Millisecond)
				completed <- id
			}(i)
		}

		cwg.Wait()

		if len(completed) != 5 {
			t.Errorf("Ожидалось 5 завершенных горутин, получено %d", len(completed))
		}

		if cwg.Count() != 0 {
			t.Errorf("Ожидался счетчик 0, получен %d", cwg.Count())
		}
	})

	t.Run("Wait без Add", func(t *testing.T) {
		cwg := NewCustomWaitGroup()

		done := make(chan bool)
		go func() {
			cwg.Wait()
			done <- true
		}()

		select {
		case <-done:
			
		case <-time.After(100 * time.Millisecond):
			t.Error("Wait заблокировался без причины")
		}
	})

}
