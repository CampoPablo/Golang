package main

import (
	"fmt"
	"strings"
)

func main() {
	simplePredicate()
	anyPredicate()
	allPredicate()
	filterPRedicate()
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

// we use the any function to check if at least one word of the words slice has letter 'w'.
func anyPredicate() {
	words := []string{"falcon", "war", "sun", "cup", "rock"}

	w := "a"
	r := any(words, func(s string) bool {
		return strings.Contains(s, w)
	})

	if r {
		fmt.Printf("The slice contains an element with %s\n", w)
	} else {
		fmt.Printf("The slice does not contain an element with %s\n", w)
	}
}

//takes a slice and a predicate as parameters. The predicate is an anonymous function which utilizes the strings.Contains function.
func any(data []string, f func(string) bool) bool {
	for _, v := range data {
		if f(v) {
			return true
		}
	}

	return false
}

// allPredicate iterates over elements of a collection and returns true if the predicate is valid for all elements.
func allPredicate() {
	words := []string{"war", "water", "cup", "tree", "storm"}
	n := 4

	res := all(words, func(s string) bool {
		return len(s) >= n
	})

	if res {
		fmt.Printf("All words have %d+ characters", n)
	} else {
		fmt.Printf("It is not true that all words have %d+ characters", n)
	}
}

// allPredicate check if all words have n or more letters
func all(data []string, f func(string) bool) bool {
	for _, e := range data {
		if !f(e) {
			return false
		}
	}

	return true
}

// find all words that start with 'w'.
func filterPRedicate() {
	words := []string{"war", "water", "cup", "tree", "storm"}

	p := "w"

	res := filter(words, func(s string) bool {
		return strings.HasPrefix(s, p)
	})

	fmt.Println(res)
}

// takes the collection and the predicate as parameters. It creates a new fltd slice into which we append all the elements that satisfy the passed predicate function.
func filter(data []string, f func(string) bool) []string {
	fltd := make([]string, 0)

	for _, e := range data {
		if f(e) {
			fltd = append(fltd, e)
		}
	}

	return fltd
}
