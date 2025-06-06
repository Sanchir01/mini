package main

import "testing"

func TestGenerateSlice(t *testing.T) {
	randomslice := randomSlice(10, 1, 10)
	t.Log(randomslice)
	evenslice := sliceExample(randomslice)
	t.Log(evenslice)
	addsliceelem := addElements(evenslice, 0)
	newSlice := copySlice(addsliceelem)
	addsliceelem[0] = 100
	newSlice[0] = 0
	t.Log("origin", addsliceelem, "copy", newSlice)
	t.Log(removeElement(newSlice, 0))
}
