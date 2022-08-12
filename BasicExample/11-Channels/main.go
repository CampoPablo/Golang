package main

import (
	"log"

	"com.pablocampo/myprogram/helpers"
)

const numPool = 1000

// We make a channel with the type int and start function calculateValue as a goroutine.
// The function calculateValue is blocked when it encounters <- c and waits to receive a value.
// The main goroutine sends a value to the channel.
func calculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	// Making a channel of value type string
	intChan := make(chan int)
	//// Closing channel
	defer close(intChan)
	// Starting a concurrent goroutine
	go calculateValue(intChan)
	// Reciving values to the channel
	num := <-intChan
	log.Println(num)
}
