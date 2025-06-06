package main

import (
	"context"
	"sort"
	"testing"
	"time"
)

func TestFanin(t *testing.T) {
	t.Run("базовая функциональность с двумя каналами", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		ch1 := make(chan int)
		ch2 := make(chan int)

		result := fanin(ctx, []chan int{ch1, ch2})

		go func() {
			ch1 <- 1
			ch1 <- 2
			close(ch1)
		}()

		go func() {
			ch2 <- 3
			ch2 <- 4
			close(ch2)
		}()

		var received []int
		for v := range result {
			received = append(received, v)
		}

		sort.Ints(received)
		expected := []int{1, 2, 3, 4}

		if len(received) != len(expected) {
			t.Fatalf("Ожидалось %d значений, получено %d", len(expected), len(received))
		}

		for i, v := range expected {
			if received[i] != v {
				t.Errorf("Ожидалось %d, получено %d", v, received[i])
			}
		}
	})

	t.Run("пустой список каналов", func(t *testing.T) {
		ctx := context.Background()
		result := fanin(ctx, []chan int{})

		select {
		case v, ok := <-result:
			if ok {
				t.Errorf("Ожидался закрытый канал, но получено значение %d", v)
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Канал не был закрыт в течение 100мс")
		}
	})

}
