package main

import "fmt"

// counter returns a function that increments and returns a counter value.
// Each call to the returned function will return the next integer.
// This is a closure that captures the `count` variable.
// It demonstrates how closures can maintain state across function calls.
// This is a simple example of a closure in Go.
func counter() func() int {
    count := 0

	// The returned function increments the count and returns it.
    return func() int {
        count++
        return count
    }
}

func main() {
	// Create a new counter function
    next := counter()

	// Call the counter function multiple times to get sequential integers
    fmt.Println(next()) // 1
    fmt.Println(next()) // 2
    fmt.Println(next()) // 3
}
