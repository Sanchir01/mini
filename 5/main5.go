package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	ok, val := intersection(a, b)
	fmt.Println("ok", ok, "val", val)
}
func intersection(a, b []int) (bool, []int) {
	lookup := make(map[int]bool)
	for _, val := range a {
		lookup[val] = true
	}

	var result []int
	for _, val := range b {
		if lookup[val] {
			result = append(result, val)
		}
	}

	hasIntersection := len(result) > 0
	return hasIntersection, result
}
