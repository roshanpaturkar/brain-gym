// Let's continue on finding the sum of subarrays. This time given a positive integer array nums,
// we want to find the length of the shortest subarray such that the subarray sum is at least target.
// Recall the same example with input nums = [1, 4, 1, 7, 3, 0, 2, 5] and target = 10,
// then the smallest window with the sum >= 10 is [7, 3] with length 2. So the output is 2.

// We'll assume for this problem that it's guaranteed target will not exceed the sum of all elements in nums.

package main

import (
	"fmt"
)

func subarraySumShortest(nums []int, target int) int {
    length := len(nums)
    window_sum := 0
    left, right := 0, 0
    
    for right < len(nums) {
        window_sum += nums[right]
        
        for window_sum >= target {
            length = min(length, right-left+1)
            window_sum -= nums[left]
            left ++
        }
        right ++
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
	fmt.Println(subarraySumShortest([]int{1, 4, 1, 7, 3, 0, 2, 5}, 10)) // Output: 2
}