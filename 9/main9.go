package main

import "fmt"

func cubeChanel(in <-chan uint8, out chan<- float64) {
	go func() {
		for n := range in {
			x := float64(n)
			res := x * x * x
			out <- res
		}
		close(out)
	}()
}

func main() {
	in := make(chan uint8)
	out := make(chan float64)
	cubeChanel(in, out)
	go func() {
		for i := uint8(1); i <= 10; i++ {
			in <- i
		}
		close(in)
	}()
	for res := range out {
		fmt.Println(res)
	}
}
