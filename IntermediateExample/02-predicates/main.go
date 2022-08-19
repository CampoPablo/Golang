package main

import (
	"fmt"
)

func main() {
	simplePredicate()
	// FALTA Completar
}

// simplePredicate takes all positive numbers from a slice into another slice.
func simplePredicate() {
	vals := []int{-2, 0, 4, 3, 1, 9, 7, -3, -5, 6}
	vals2 := []int{}

	for _, val := range vals {

		if isPositive(val) {

			vals2 = append(vals2, val)
		}
	}

	fmt.Println(vals)
	fmt.Println(vals2)
}

func isPositive(val int) bool {
	if val > 0 {
		return true
	} else {
		return false
	}
}
