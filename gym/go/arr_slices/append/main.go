package main

import "fmt"

func modifySlice(s []int) {
	s = append(s, 4, 5) // This append might create a new backing array for 's'
	fmt.Println("Inside function:", s)
}

func main() {
	mySlice := []int{1, 2, 3}
	modifySlice(mySlice)
	fmt.Println("Outside function:", mySlice) // Output: Outside function: [1 2 3]
}
