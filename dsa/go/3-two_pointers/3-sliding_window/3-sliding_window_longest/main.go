// Recall finding the largest size k subarray sum of an integer array in Largest Subarray Sum.
// What if we don't need the largest sum among all subarrays of fixed size k, but instead,
// we want to find the length of the longest subarray with sum smaller than or equal to a target?
// Given input nums = [1, 6, 3, 1, 2, 4, 5] and target = 10, 
// then the longest subarray that does not exceed 10 is [3, 1, 2, 4], so the output is 4 (length of [3, 1, 2, 4]).

package main

import (
	"fmt"
)

func subarraySumLongest(nums []int, target int) int {
	// Initialize the sum and length of the longest subarray
	sum, longest := 0, 0

	// Initialize the left and right pointers
	left, right := 0, 0

	// Iterate over the array
	for right < len(nums) {
		// Add the current element to the sum
		sum += nums[right]

		// If the sum exceeds the target
		for sum > target {
			// Subtract the leftmost element from the sum
			sum -= nums[left]

			// Move the left pointer to the right
			left++
		}

		// Update the length of the longest subarray
		longest = max(longest, right-left+1)

		// Move the right pointer to the right
		right++
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 6, 3, 1, 2, 4, 5}
	target := 10
	fmt.Println(subarraySumLongest(nums, target)) // Output: 4
	
	nums = []int{1, 6, 3, 1, 2, 4, 5, 10, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target = 20
	fmt.Println(subarraySumLongest(nums, target)) // Output: 6
}