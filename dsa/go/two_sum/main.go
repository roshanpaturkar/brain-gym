// # Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// # You may assume that each input would have exactly one solution, and you may not use the same element twice.
// # You can return the answer in any order.

// # Example 1:
// # Input: nums = [2,7,11,15], target = 9
// # Output: [0,1]
// # Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

package main

import "fmt"

func two_sum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		j, n := m[target - num]
		if n {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(two_sum(nums, target))

	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(two_sum(nums, target))

	nums = []int{3, 3}
	target = 6
	fmt.Println(two_sum(nums, target))
}