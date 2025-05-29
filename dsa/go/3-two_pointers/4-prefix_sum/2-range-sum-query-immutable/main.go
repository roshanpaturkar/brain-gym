package main

import "fmt"

func rangeSumQueryImmutable(nums []int, left int, right int) int {
    currentSum := 0
    
    for left <= right {
        currentSum += nums[left]
        left ++
    }
    
    return currentSum
}

func main() {
	fmt.Println(rangeSumQueryImmutable([]int{3, 0, 1, 4, 2}, 1, 3))			//	5
	fmt.Println(rangeSumQueryImmutable([]int{-2, 0, 3, -5, 2, -1}, 0, 2))	//	1
	fmt.Println(rangeSumQueryImmutable([]int{-2, 0, 3, -5, 2, -1}, 2, 5))	// -1
	fmt.Println(rangeSumQueryImmutable([]int{-2, 0, 3, -5, 2, -1}, 0, 5))	// -3
}