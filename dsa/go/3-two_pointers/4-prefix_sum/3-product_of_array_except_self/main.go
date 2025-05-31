package main

import "fmt"

func initResult(nums []int) []int {
	result := make([]int, len(nums))
	for i := range result {
		result[i] = 1
	}

	return result
}

func productOfArrayExceptSelf(nums []int) []int {
	result := initResult(nums)

	left := 1
	for i := 0; i < len(nums); i++ {
		result[i] = left
		left *= nums[i]
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}

func main() {
	fmt.Println(productOfArrayExceptSelf([]int{1, 2, 3, 4}))      //  [24, 12, 8, 6]
	fmt.Println(productOfArrayExceptSelf([]int{1, 2, 3, 4, 5}))   //  [120, 60, 40, 30, 24]
	fmt.Println(productOfArrayExceptSelf([]int{-1, 1, 0, -3, 3})) //  [0, 0, 9, 0, 0]
}
