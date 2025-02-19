// Given an array of integers, move all the 0s to the back of the array while maintaining the relative order of the non-zero elements. Do this in-place using constant auxiliary space.
// Input:
// [1, 0, 2, 0, 0, 7]
// Output:
// [1, 2, 7, 0, 0, 0]

package main

import (
	"fmt"
)

func moveZeros(nums []int) []int {
    slow := 0
    
    for fast := 0; fast < len(nums); fast++ {
        if nums[fast] != 0 {
            nums[slow], nums[fast] = nums[fast], nums[slow]
            slow++
        }
    }
    
    return nums
}

func main() {
	nums := []int{1, 0, 2, 0, 0, 7}
	fmt.Println(moveZeros(nums))
	
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(moveZeros(nums))
	
	nums = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(moveZeros(nums))
}