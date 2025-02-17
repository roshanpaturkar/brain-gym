package main

import (
	"fmt"
) 

func peakOfMountainArray(arr []int) int {
    left := 0
    right := len(arr) - 1
    peak := -1
    
    for left <= right {
        mid := (left + right) / 2
        
        if mid < len(arr)-1 && arr[mid] > arr[mid+1] {
            peak = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    
    return peak
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1}
	fmt.Println(peakOfMountainArray(arr))

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(peakOfMountainArray(arr))

	arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(peakOfMountainArray(arr))

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(peakOfMountainArray(arr))
}