package main

import "fmt"

func initSumArray(nums []int) []int {
	sumArray := []int{0}
	for i := range(nums) {
		sumArray = append(sumArray, sumArray[i] + nums[i])
	}
	return sumArray
}

func rangeSumQueryImmutable(nums []int, left, right int) int {
	prefixSum := initSumArray(nums)
	return prefixSum[right + 1] - prefixSum[left]
}

func main() {
	fmt.Println(rangeSumQueryImmutable([]int{3, 0, 1, 4, 2}, 1, 3))			//	5
	fmt.Println(rangeSumQueryImmutable([]int{-2, 0, 3, -5, 2, -1}, 0, 2))	//	1
	fmt.Println(rangeSumQueryImmutable([]int{-2, 0, 3, -5, 2, -1}, 2, 5))	// -1
	fmt.Println(rangeSumQueryImmutable([]int{-2, 0, 3, -5, 2, -1}, 0, 5))	// -3
}