// Given an array of integers sorted in ascending order, find two numbers that add up to a given target. Return the indices of the two numbers in ascending order. You can assume elements in the array are unique and there is only one solution. Do this in O(n) time and with constant auxiliary space.
// Input:
// arr: a sorted integer array
// target: the target sum we want to reach
// Sample Input: [2, 3, 4, 5, 8, 11, 18], 8
// Sample Output: 1 3

package main

import (
	"fmt"
)

func twoSumSorted(arr []int, target int) []int {
    left := 0
    right := len(arr) - 1
    
    for left < right {
        sum := arr[left] + arr[right]
        if sum == target {
            return []int{left, right}
        } else if sum > target {
            right --
        } else {
            left ++
        }
    }
    
    return []int{}
}

func main() {
	arr := []int{2, 3, 4, 5, 8, 11, 18}
	target := 8
	fmt.Println(twoSumSorted(arr, target))
	
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target = 10
	fmt.Println(twoSumSorted(arr, target))
	
	arr = []int{1, 2, 3, 4, 5, 7, 8, 9}
	target = 6
	fmt.Println(twoSumSorted(arr, target))
	
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target = 18
	fmt.Println(twoSumSorted(arr, target))
}