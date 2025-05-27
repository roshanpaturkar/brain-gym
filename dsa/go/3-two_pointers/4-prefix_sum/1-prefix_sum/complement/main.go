package main

import "fmt"

func subarraySum(arr []int, target int) []int {
	prefix_sums := make(map[int]int)
	
	curent_sum := 0
	prefix_sums[0] = 0
	
	for i := range(arr) {
		curent_sum += arr[i]
		complement := curent_sum - target
		
		if _, exists := prefix_sums[complement]; exists {
			return []int{prefix_sums[complement], i + 1}
		}
		prefix_sums[curent_sum] = i + 1
	}
	return []int{}
}

func main() {
	fmt.Println(subarraySum([]int{1, 3, -3, 8, 5, 7}, 5))		//	2, 4
	fmt.Println(subarraySum([]int{1, 1, 1}, 3))					//	0, 3
	fmt.Println(subarraySum([]int{1, -20, -3, 30, 5, 7}, 7))	//	1, 4
}