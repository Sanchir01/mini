package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	res := difference(slice1, slice2)
	fmt.Println(res)
}

func difference(slice1, slice2 []string) []string {
	lookup := make(map[string]bool)
	for _, item := range slice2 {
		lookup[item] = true
	}

	var result []string
	for _, item := range slice1 {
		if !lookup[item] {
			result = append(result, item)
		}
	}

	return result
}
