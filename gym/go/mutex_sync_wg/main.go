package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup         // WaitGroup to wait for all goroutines
    var mu sync.Mutex             // Mutex to protect shared resource
    counter := 0                  // Shared resource

    numWorkers := 5
    wg.Add(numWorkers)

    for i := 0; i < numWorkers; i++ {
        go func(workerID int) {
            defer wg.Done() // Notify WaitGroup when done

            for j := 0; j < 1000; j++ {
                mu.Lock()         // Lock before updating counter
                counter++
                mu.Unlock()       // Unlock after updating
            }
            fmt.Printf("Worker %d done\n", workerID)
        }(i)
    }

    wg.Wait() // Wait for all goroutines to complete
    fmt.Println("Final Counter:", counter)
}
