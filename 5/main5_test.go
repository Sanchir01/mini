package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name              string
		a                 []int
		b                 []int
		expectedHas       bool
		expectedIntersect []int
	}{
		{
			name:              "основной пример",
			a:                 []int{65, 3, 58, 678, 64},
			b:                 []int{64, 2, 3, 43},
			expectedHas:       true,
			expectedIntersect: []int{64, 3},
		},
		{
			name:              "нет пересечений",
			a:                 []int{1, 2, 3},
			b:                 []int{4, 5, 6},
			expectedHas:       false,
			expectedIntersect: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasIntersection, intersections := intersection(tt.a, tt.b)
			if hasIntersection != tt.expectedHas {
				t.Errorf("intersection() hasIntersection = %v, expected %v", hasIntersection, tt.expectedHas)
			}

			sortedResult := make([]int, len(intersections))
			sort.Ints(sortedResult)

			sortedExpected := make([]int, len(tt.expectedIntersect))
			sort.Ints(sortedExpected)

			if !reflect.DeepEqual(sortedResult, sortedExpected) {
				t.Errorf("intersection() intersections = %v, expected %v", intersections, tt.expectedIntersect)
			}
		})
	}
}
