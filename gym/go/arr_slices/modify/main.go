package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[0:3] // slice1 is [1 2 3]
	slice2 := arr[0:4] // slice2 is [1 2 3 4]

	slice1[0] = 99 // Modifies the underlying array

	fmt.Println(arr)    // Output: [99 2 3 4 5]
	fmt.Println(slice2) // Output: [99 2 3 4] - slice2 is also affected
}
