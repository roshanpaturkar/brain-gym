package main

import "fmt"

func containerWithMostWater(arr []int) int {
    left := 0
    right := len(arr) - 1
    maxArea := 0
    
    for left < right {
        area := (right - left) * min(arr[left], arr[right])
        maxArea = max(maxArea, area)
        
        if arr[left] < arr[right] {
            left ++
        } else {
            right --
        }
    }
    
    return maxArea
}

func main() {
	fmt.Println(containerWithMostWater([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))	// 49
	fmt.Println(containerWithMostWater([]int{1, 1}))	// 1
	fmt.Println(containerWithMostWater([]int{4, 3, 2, 1, 4}))	// 16
	fmt.Println(containerWithMostWater([]int{1, 2, 1}))	// 2
}