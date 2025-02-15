package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	left := 0
    right := len(arr) - 1

    for left <= right { // <= here because left and right could point to the same element, < would miss it
        mid := left + (right-left)/2 // use `(right - left) / 2` to prevent `left + right` potential overflow
        // found target, return its index
        if arr[mid] == target {
            return mid
        }
        if arr[mid] < target {
            // middle less than target, discard left half by making left search boundary `mid + 1`
            left = mid + 1
        } else {
            // middle greater than target, discard right half by making right search boundary `mid - 1`
            right = mid - 1
        }
    }
    return -1 // if we get here we didn't hit above return so we didn't find target
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 5
	
	if idx := binarySearch(arr, target); idx != -1 {
		fmt.Printf("Target %d found at index %d\n", target, idx)
	} else {
		fmt.Printf("Target %d not found\n", target)
	}
}