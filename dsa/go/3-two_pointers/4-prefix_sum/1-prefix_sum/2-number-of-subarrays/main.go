package main

import "fmt"

func subarraySumTotal(arr []int, target int) int {
	prefix_sum := make(map[int]int)
	prefix_sum[0] = 1
	
	current_sum := 0
	count := 0
	
	for _, value := range(arr) {
		current_sum += value
		complement := current_sum - target
		
		if _, exists := prefix_sum[complement]; exists {
			count += prefix_sum[complement]
		}
		prefix_sum[current_sum] ++
	}
	return count
}

func main() {
	fmt.Println(subarraySumTotal([]int{1, 2, 3}, 3))	//	2
}

