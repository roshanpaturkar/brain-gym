package main

import (
	"fmt"
)

func findBoundary(arr []bool) int {
    left := 0
    right := len(arr) - 1
    boundaryIndex := -1
    for left <= right {
        mid := (left + right) / 2
        if arr[mid] {
            boundaryIndex = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return boundaryIndex
}

func main() {
	arr := []bool{false, false, false, false, false, true, true, true, true}
	fmt.Println(findBoundary(arr))
}