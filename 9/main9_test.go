package main

import (
	"math"
	"testing"
)

func collectResults(out <-chan float64) []float64 {
	var results []float64
	for v := range out {
		results = append(results, v)
	}
	return results
}
func TestStartPipeline(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	cubeChanel(in, out)

	input := []uint8{2, 3, 4, 5}
	expected := []float64{8, 27, 64, 125}

	go func() {
		for _, n := range input {
			in <- n
		}
		close(in)
	}()

	results := collectResults(out)

	if len(results) != len(expected) {
		t.Fatalf("unexpected result length: got %d, want %d", len(results), len(expected))
	}

	for i := range results {
		if math.Abs(results[i]-expected[i]) > 1e-9 {
			t.Errorf("wrong result at index %d: got %.2f, want %.2f", i, results[i], expected[i])
		}
	}
}
