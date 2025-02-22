// Given an array representing heights of vertical lines, find the maximum area of water trapped between two lines.
// Input: [1,8,6,2,5,4,8,3,7].
// Output: 49.

package main

import (
	"fmt"
)

func containerWithMostWater(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		area := (right - left) * min(height[left], height[right])
		maxArea = max(maxArea, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(containerWithMostWater(height))
}