package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 50, 30, 40, 45}
	var nums2 []int
	max := 0
	min := 1000
	for i := len(nums) - 1; i >= 0; i-- {
	nums2 =	append(nums2, nums[i])

		if nums[i] > max {
			max = nums[i]
		}

		if nums[i] < min {
			min = nums[i]
		}
	}

	avg := (max + min) / 2

	fmt.Println("Reverse: ", nums2)
	fmt.Println("Max: ", max)
	fmt.Println("Min: ", min)
	fmt.Println("Avg: ", avg)
}