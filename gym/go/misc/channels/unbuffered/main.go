package main

import (
	"fmt"
	"time"
)

// producer sends numbers to a channel
func producer(ch chan int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Println("Producer: sending", i)
		ch <- i // Send the number to the channel
		time.Sleep(500 * time.Millisecond) // Simulate some work
	}
	close(ch) // Close the channel when done sending
	fmt.Println("Producer: finished and closed channel")
}

func main() {
	// Create an unbuffered channel for integers
	dataChannel := make(chan int)
	numberOfItems := 5

	fmt.Println("Main: Starting producer Goroutine...")
	// Start the producer Goroutine
	go producer(dataChannel, numberOfItems)

	fmt.Println("Main: Waiting to receive data...")

	// Receive data from the channel until it's closed
	// The loop `for item := range dataChannel` automatically
	// handles checking if the channel is closed and breaks when it is.
	for item := range dataChannel {
		fmt.Println("Main: received", item)
		time.Sleep(200 * time.Millisecond) // Simulate some processing
	}

	fmt.Println("Main: All data received. Program finished.")
}