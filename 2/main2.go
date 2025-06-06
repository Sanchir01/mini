package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randomint := randomSlice(20, 1, 10)
	fmt.Println(randomint)
	evenSlice := sliceExample(randomint)
	fmt.Println(evenSlice)
	tesint := addElements(evenSlice, 1)
	fmt.Println(tesint)
	newSlice := copySlice(tesint)
	newSlice[1] = 2
	fmt.Println(removeElement(newSlice, 1))
}

func randomSlice(length, min, max int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = rand.Intn(max-min+1) + min
	}
	return slice
}

func sliceExample(intingslice []int) []int {
	var myNewSlice []int
	for _, v := range intingslice {
		if v%2 == 0 {
			myNewSlice = append(myNewSlice, v)
		}
	}
	return myNewSlice
}

func addElements(slice []int, value int) []int {
	return append(slice, value)
}

func copySlice(original []int) []int {
	copied := make([]int, len(original))
	copy(copied, original)
	return copied
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
