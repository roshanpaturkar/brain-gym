package main

import "fmt"

func main() {
    a := [...]int{1, 2, 3, 4, 5}

	b := a[:1]
	c := a[2:]

	x := append(b, c...)
	y := append(b, x...)

	fmt.Println("Original array:", a)
	fmt.Println("Slice x:", x)
	fmt.Println("Slice y:", y)
	fmt.Println("Length of x:", len(x))
	fmt.Println("Length of y:", len(y))
}