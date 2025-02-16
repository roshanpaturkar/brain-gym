package main

import (
	"fmt"
)

func squareRoot(n int) int {
    if n == 0 {
        return 0
    }
    left, right := 1, n
    res := -1
    for left <= right {
        mid := (left + right) / 2
        if mid*mid == n {
            return mid
        } else if mid*mid > n {
            res = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return res - 1
}

func main() {
	n := 16
	fmt.Println(squareRoot(n))

	n = 17
	fmt.Println(squareRoot(n))
	
	n = 0
	fmt.Println(squareRoot(n))
	
	n = 6
	fmt.Println(squareRoot(n))
}