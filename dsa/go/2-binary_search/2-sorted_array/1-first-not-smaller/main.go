package main

import (
	"fmt"
)

func firstNotSmaller(arr []int, target int) int {
    left := 0
    right := len(arr) - 1
    boundaryIndex := -1
    for left <= right {
        mid := (left + right) / 2
        if arr[mid] >= target {
            boundaryIndex = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return boundaryIndex
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 5
	fmt.Println(firstNotSmaller(arr, target))
	
	
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target = 10
	fmt.Println(firstNotSmaller(arr, target))
	
	
	arr = []int{1, 2, 3, 4, 5, 7, 8, 9}
	target = 6
	fmt.Println(firstNotSmaller(arr, target))
}