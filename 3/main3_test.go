package main

import (
	"testing"
)

func TestAddAndGet(t *testing.T) {
	m := NewStringIntMap()
	key := "key1"
	m.Add(key, 100)

	ok := m.IsExist(key)
	if !ok {
		t.Errorf("not exist key")
	}
	val, ok := m.Get(key)
	if !ok {
		t.Errorf("expected 100, got %d (exists: %v)", val, ok)
	}

	newmap := m.Copy()
	t.Log(newmap)
	m.Remove(key)
}
