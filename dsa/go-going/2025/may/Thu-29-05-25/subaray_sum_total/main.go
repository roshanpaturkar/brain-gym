package main

import "fmt"

func subarraySumTotal(array []int, target int) int {
	prefixSum := make(map[int]int)
	prefixSum[0] = 1
	
	currentSum := 0
	count := 0
	
	for i := range(array) {
		currentSum += array[i]
		complement := currentSum - target
		
		if _, exists := prefixSum[complement]; exists {
			count += prefixSum[complement]
		}
		prefixSum[currentSum]++
	}
	
	return count
}

func main() {
	fmt.Println(subarraySumTotal([]int{1, 2, 3}, 3))				//	2
	fmt.Println(subarraySumTotal([]int{1, 1, 1}, 2))				//	2
	fmt.Println(subarraySumTotal([]int{10, 5, -5, -20, 10}, -10))	//	3
}