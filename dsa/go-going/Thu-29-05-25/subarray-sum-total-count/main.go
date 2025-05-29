package main

import "fmt"

func subarraySum(array []int, target int) []int {
	prefixSum := make(map[int]int)
	prefixSum[0] = 0
	
	currentSum := 0
	
	for i := range(array) {
		currentSum += array[i]
		complement := currentSum - target
		
		if _, exists := prefixSum[complement]; exists {
			return []int{prefixSum[complement], i + 1}
		}
		prefixSum[currentSum] = i + 1
	}
	return []int{}
}

func main() {
	fmt.Println(subarraySum([]int{1, 3, -3, 8, 5, 7}, 5))		//	2, 4
	fmt.Println(subarraySum([]int{1, 1, 1}, 3))					//	0, 3
	fmt.Println(subarraySum([]int{1, -20, -3, 30, 5, 7}, 7))	//	1, 4
}