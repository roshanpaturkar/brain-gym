package main

import "fmt"
import "time"

func main() {
	ch := make(chan int, 2) // Buffered channel for demonstration

	ch <- 1
	ch <- 2
	close(ch) // Close the channel after sending values

	// Receive existing values
	val1, ok1 := <-ch
	fmt.Printf("Received: %d, Open: %t\n", val1, ok1)

	val2, ok2 := <-ch
	fmt.Printf("Received: %d, Open: %t\n", val2, ok2)

	// Receive from closed channel (after all sent values are consumed)
	val3, ok3 := <-ch
	fmt.Printf("Received: %d, Open: %t\n", val3, ok3)

	// Another receive from closed channel
	val4, ok4 := <-ch
	fmt.Printf("Received: %d, Open: %t\n", val4, ok4)

	// Demonstrating blocking behavior with unbuffered channel
	unbufferedCh := make(chan int)
	go func() {
		time.Sleep(100 * time.Millisecond) // Simulate some work
		close(unbufferedCh)
	}()

	// Receiving from a closed unbuffered channel
	val5, ok5 := <-unbufferedCh
	fmt.Printf("Received from unbuffered: %d, Open: %t\n", val5, ok5)
}