package main

import "fmt"

func main() {
	s := make([]int, 3, 5) // len=3, cap=5
	// s[3] = 10 // This will cause a runtime panic: index out of range [3] with length 3
	fmt.Println(len(s), cap(s))
}
