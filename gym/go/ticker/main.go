package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a ticker that ticks every 1 second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // Always stop the ticker to release resources

	done := make(chan bool)

	// Stop after 5 seconds
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	fmt.Println("Ticker started...")
	for {
		select {
		case <-done:
			fmt.Println("Ticker stopped.")
			return
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		}
	}
}
