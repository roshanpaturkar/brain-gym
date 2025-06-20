package main

import "fmt"

func even(n int, ch chan int) {
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			ch <- i
		}
	}
	close(ch)

}

func odd(n int, ch chan int) {
	for i := 1; i < n; i++ {
		if i%2 != 0 {
			ch <- i
		}
	}
	close(ch)
}

func main() {
	ech := make(chan int, 10)

	och := make(chan int, 10)

	n := 10

	go even(n, ech)
	go odd(n, och)

	for ch := range ech {
		fmt.Println(ch)
	}

	for ch := range och {
		fmt.Println(ch)
	}
}
