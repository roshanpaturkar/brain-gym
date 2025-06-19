package main

import "fmt"

// Function to check if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func sumEven(numbers []int, ch chan int) {
	sum := 0
	for _, n := range numbers {
		if n%2 == 0 {
			sum += n
		}
	}
	ch <- sum
}

func sumOdd(numbers []int, ch chan int) {
	sum := 0
	for _, n := range numbers {
		if n%2 != 0 {
			sum += n
		}
	}
	ch <- sum
}

func sumPrime(numbers []int, ch chan int) {
	sum := 0
	for _, n := range numbers {
		if isPrime(n) {
			sum += n
		}
	}
	ch <- sum
}

func main() {
	// Generate numbers from 1 to 100
	numbers := make([]int, 100)
	for i := range numbers {
		numbers[i] = i + 1
	}

	// Create channels
	evenCh := make(chan int)
	oddCh := make(chan int)
	primeCh := make(chan int)

	// Launch goroutines
	go sumEven(numbers, evenCh)
	go sumOdd(numbers, oddCh)
	go sumPrime(numbers, primeCh)

	// Receive from channels
	evenSum := <-evenCh
	oddSum := <-oddCh
	primeSum := <-primeCh

	// Print results
	fmt.Println("Sum of Even Numbers:", evenSum)
	fmt.Println("Sum of Odd Numbers:", oddSum)
	fmt.Println("Sum of Prime Numbers:", primeSum)
}
