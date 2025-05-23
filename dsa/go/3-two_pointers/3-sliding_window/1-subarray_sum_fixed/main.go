// Given an array (list) nums consisted of only non-negative integers, find the largest sum among all subarrays of length k in nums.
// For example, if the input is nums = [1, 2, 3, 7, 4, 1], k = 3,
//then the output would be 14 as the largest length 3 subarray sum is given by [3, 7, 4] which sums to 14.

package main

import (
	"fmt"
)

func maxSumSubarray(nums []int, k int) int {
	sum := 0
	maxSum := 0

	for i := range nums {
		sum += nums[i]

		if i >= k-1 {
			maxSum = max(maxSum, sum)
			sum -= nums[i-k+1]
		}
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 3, 7, 4, 1}
	k := 3
	fmt.Println(maxSumSubarray(nums, k)) // Output: 14
	
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k = 4
	fmt.Println(maxSumSubarray(nums, k)) // Output: 34
}