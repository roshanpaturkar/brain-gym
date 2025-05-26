package main

import "fmt"

func shortestSubarraySum(nums []int, target int) int {
	left, right := 0, 0
	window_sum := 0
	length := len(nums)

	for right < len(nums) {
		window_sum += nums[right]

		for window_sum >= target {
			length = min(length, right-left+1)
			window_sum -= nums[left]
			left++
		}
		right++
	}

	return length
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 4, 1, 7, 3, 0, 2, 5}
	fmt.Println(shortestSubarraySum(nums, 10)) // 2

	nums = []int{1, 6, 8}
	fmt.Println(shortestSubarraySum(nums, 1)) // 1

	nums = []int{6, 6, 6, 6, 6, 6, 6}
	fmt.Println(shortestSubarraySum(nums, 19)) // 4

	nums = []int{1, 1, 1}
	fmt.Println(shortestSubarraySum(nums, 3)) // 3
}
