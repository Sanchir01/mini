package main

import (
	"testing"
)

func TestAddAndGet(t *testing.T) {
	testcases := []struct {
		name  string
		key   string
		value int
	}{
		{
			name:  "test1",
			key:   "key1",
			value: 30,
		},
		{
			name:  "test2",
			key:   "key2",
			value: 2,
		},
		{
			name:  "test3",
			key:   "key2",
			value: 7899,
		},
	}
	m := NewStringIntMap()

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			m.Add(tt.key, tt.value)

			ok := m.IsExist(tt.key)
			if !ok {
				t.Errorf("not exist key")
			}
			val, ok := m.Get(tt.key)
			if !ok {
				t.Errorf("expected 100, got %d (exists: %v)", val, ok)
			}

			newmap := m.Copy()
			t.Log(newmap)
			m.Remove(tt.key)
		})
	}

}
