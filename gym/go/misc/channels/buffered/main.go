package main

import (
	"fmt"
	"time"
)

// producer sends numbers to a channel
func producerBuffered(ch chan int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Println("Producer (Buffered): sending", i)
		ch <- i // Send the number to the channel
		fmt.Printf("Producer (Buffered): buffer size %d, length %d\n", cap(ch), len(ch))
		time.Sleep(100 * time.Millisecond) // Simulate some work
	}
	close(ch) // Close the channel when done sending
	fmt.Println("Producer (Buffered): finished and closed channel")
}

func main() {
	// Create a buffered channel for integers with a capacity of 3
	bufferedDataChannel := make(chan int, 3)
	numberOfItems := 5

	fmt.Println("Main: Starting producer (Buffered) Goroutine...")
	// Start the producer Goroutine
	go producerBuffered(bufferedDataChannel, numberOfItems)

	fmt.Println("Main: Waiting to receive data from buffered channel...")
	time.Sleep(1 * time.Second) // Give producer some time to fill the buffer

	// Receive data from the channel until it's closed
	for item := range bufferedDataChannel {
		fmt.Println("Main: received (Buffered)", item)
		fmt.Printf("Main: received (Buffered), buffer size %d, length %d\n", cap(bufferedDataChannel), len(bufferedDataChannel))
		time.Sleep(700 * time.Millisecond) // Simulate slower processing
	}

	fmt.Println("Main: All data received from buffered channel. Program finished.")
}