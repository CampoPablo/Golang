package main

import "log"

func main() {
	var myString string
	myString = "Green"
	log.Println("Mystring is set to", myString)
	changeUsingPointer(&myString)
	log.Println("After func call my string is", myString)
}

func changeUsingPointer(s *string) {
	var newValue string
	log.Println("s is set to", s)
	newValue = "Blue"
	*s = newValue
}
