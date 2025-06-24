package main

import (
    "sync"
)

func main() {
    var m sync.Map
    var wg sync.WaitGroup
    // Concurrent write operations
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            m.Store(n, n*10)
        }(i)
    }

    wg.Wait()
    // Concurrent read operations
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            if val, ok := m.Load(n); ok {
                println(val.(int))
            }
        }(i)
    }

    wg.Wait()
}