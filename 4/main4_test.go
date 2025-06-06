package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestAll(t *testing.T) {
	testcases := []struct {
		name   string
		slice1 []string
		slice2 []string
		want   []string
	}{
		{
			name:   "basic",
			slice1: []string{"a", "b", "c"},
			slice2: []string{"b"},
			want:   []string{"a", "c"},
		},
		{
			name:   "all",
			slice1: []string{"a", "b", "c"},
			slice2: []string{"a", "b", "c"},
			want:   []string{},
		},
		{
			name:   "all excluded",
			slice1: []string{},
			slice2: []string{},
			want:   []string{},
		},
		{
			name:   "no excluded",
			slice1: []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2: []string{"banana", "date", "fig"},
			want:   []string{"apple", "cherry", "43", "lead", "gno1"},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			result := difference(tt.slice1, tt.slice2)

			sortedResult := make([]string, len(result))
			sort.Strings(sortedResult)

			sortedExpected := make([]string, len(tt.want))
			sort.Strings(sortedExpected)

			if !reflect.DeepEqual(sortedResult, sortedExpected) {
				t.Errorf("difference() = %v, expected %v", result, tt.want)
			}
		})
	}
}
