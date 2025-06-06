package main

import "testing"

func TestGenerateRandomIntInRange(t *testing.T) {
	const iterations = 1000
	for i := 0; i < iterations; i++ {
		n := <-generateRandomInt()
		if n < 1 || n > 1000 {
			t.Errorf("generated number %d is out of range [1, 1000]", n)
		}
	}
}
