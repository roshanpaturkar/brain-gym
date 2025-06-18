package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	arraySize     = 100000 // Total size of the array
	chunkSize     = 10000  // Size of each chunk processed by a goroutine
	maxRandomVal  = 9      // Maximum random value (0-9)
	numWorkers    = arraySize / chunkSize
)

func main() {
	// Initialize a new random source based on current time
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Create and initialize the array with random values
	data := make([]int, arraySize)
	for i := 0; i < arraySize; i++ {
		data[i] = random.Intn(maxRandomVal + 1) // Generate random number between 0 and 9
	}

	fmt.Printf("Array of size %d initialized with random values (0-%d).\n", arraySize, maxRandomVal)

	// Use a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup
	// Use a mutex to protect the shared final counts map
	var mu sync.Mutex
	// Store counts from all goroutines (shared resource)
	finalCounts := make(map[int]int)

	// Launch goroutines to process chunks
	for i := 0; i < numWorkers; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		startIndex := i * chunkSize
		endIndex := startIndex + chunkSize
		// Ensure endIndex doesn't go beyond arraySize for the last chunk
		if endIndex > arraySize {
			endIndex = arraySize
		}

		// Run the counting logic in a goroutine
		go func(chunk []int, workerID int) {
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine finishes

			localCounts := make(map[int]int)
			for _, num := range chunk {
				localCounts[num]++
			}

			fmt.Printf("Worker %d processed chunk from index %d to %d. Local counts: %v\n",
				workerID, startIndex, endIndex-1, localCounts)

			// Safely merge local counts into the final shared counts map
			mu.Lock()         // Acquire mutex lock
			for num, count := range localCounts {
				finalCounts[num] += count
			}
			mu.Unlock()       // Release mutex lock
		}(data[startIndex:endIndex], i) // Pass a slice (chunk) to the goroutine
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("\n--- Final Occurrences Count ---")
	// Print final counts in a sorted order for better readability
	for i := 0; i <= maxRandomVal; i++ {
		fmt.Printf("Number %d: %d occurrences\n", i, finalCounts[i])
	}
}