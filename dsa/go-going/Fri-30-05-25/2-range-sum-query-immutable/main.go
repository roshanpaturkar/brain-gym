package main

import "fmt"

func initArraySum(array []int) []int {
	sum := []int{0}
	
	for i := range(array) {
		sum = append(sum, sum[i] + array[i])
	}
	
	return sum
}

func rangeSumQueryImmutable(array []int, left, right int) int {
	prefixArraySum := initArraySum(array)
	
	return prefixArraySum[right + 1] - prefixArraySum[left]
}

func main() {
	fmt.Println(rangeSumQueryImmutable([]int{3, 0, 1, 4, 2}, 1, 3))			//	5
	fmt.Println(rangeSumQueryImmutable([]int{-2, 0, 3, -5, 2, -1}, 0, 2))	//	1
	fmt.Println(rangeSumQueryImmutable([]int{-2, 0, 3, -5, 2, -1}, 2, 5))	// -1
	fmt.Println(rangeSumQueryImmutable([]int{-2, 0, 3, -5, 2, -1}, 0, 5))	// -3
}