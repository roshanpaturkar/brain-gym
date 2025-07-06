package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch) // Close the channel

	// Attempting to send to a closed channel will cause a panic
	// ch <- 10 

	fmt.Println("Program continues after potential panic (if line above was uncommented)")
}