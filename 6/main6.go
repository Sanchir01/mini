package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := generateRandomInt()
	fmt.Println(<-ch)
}

func generateRandomInt() <-chan int {
	ch := make(chan int)

	go func() {
		ch <- rand.Intn(1000) + 1
	}()
	return ch
}
