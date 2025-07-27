package main

import (
	"fmt"
	"time"
)

func worker(id int, semaphore chan struct{}) {
	semaphore <- struct{}{} // Acquire the semaphore
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(2 * time.Second) // Simulate work
	fmt.Printf("Worker %d done\n", id)
	<-semaphore // Release the semaphore
}

func main() {
	semaphore := make(chan struct{}, 2) // Allow 2 workers at a time
	for i := 1; i <= 5; i++ {
		go worker(i, semaphore)
	}
	time.Sleep(10 * time.Second) // Wait for all workers to finish
	fmt.Println("All workers done")
}
