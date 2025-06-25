package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	max := 10
	num := 0
	isEvenTurn := true

	// Even number printer
	go func() {
		defer wg.Done()
		for num < max {
			mu.Lock()
			for !isEvenTurn {
				cond.Wait()
			}
			if num < max {
				fmt.Println("Even:", num)
				num++
				isEvenTurn = false
				cond.Signal()
			}
			mu.Unlock()
		}
	}()

	// Odd number printer
	go func() {
		defer wg.Done()
		for num < max {
			mu.Lock()
			for isEvenTurn {
				cond.Wait()
			}
			if num < max {
				fmt.Println("Odd:", num)
				num++
				isEvenTurn = true
				cond.Signal()
			}
			mu.Unlock()
		}
	}()

	wg.Wait()
}
