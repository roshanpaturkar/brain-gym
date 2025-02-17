package main

import (
	"fmt"
)

func findMinRotated(arr []int) int {
    left, right := 0, len(arr)-1
    boundaryIndex := -1

    for left <= right {
        mid := (left + right) / 2
        // if <= last element, then belongs to lower half
        if arr[mid] <= arr[len(arr)-1] {
            boundaryIndex = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return boundaryIndex
}

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMinRotated(arr))

	arr = []int{11, 13, 15, 17}
	fmt.Println(findMinRotated(arr))

	arr = []int{5, 6, 7, 8, 9, 10, 1, 2, 3}
	fmt.Println(findMinRotated(arr))

	arr = []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 1, 2, 3}
	fmt.Println(findMinRotated(arr))
}