package main

import (
	"fmt"
	"sync"
)

func worker(id int, data <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range data {
		fmt.Printf("Worker %d processing %d\n", id, num)
		results <- num * 2 // Double the number and send it to the results channel
	}
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	dataChan := make(chan int, 3) // Buffered channel to hold data
	resultsChan := make(chan int, 3) // Buffered channel to receive results

	// Launch 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, dataChan, resultsChan, &wg)
	}

	// Send data to the workers
	for i := 1; i <= 5; i++ {
		dataChan <- i
	}
	close(dataChan) // Close the data channel to signal no more data

	// Wait for all workers to finish
	wg.Wait()
	close(resultsChan) // Close the results channel

	// Process results
	for result := range resultsChan {
		fmt.Println("Result:", result)
	}
	fmt.Println("All done!")
}
