package main

import "fmt"

func subarraySum(arr []int, target int) []int {
    sum := 0
    left, right := 0, 0
    
    for right < len(arr) {
        sum += arr[right]
        
        for sum > target {
            sum -= arr[left]
            left ++
        }
        right ++
        
        if sum == target {
            return []int{left, right}
        }
    } 
    
    return []int{}
}

func main() {
	fmt.Println(subarraySum([]int{1, 3, -3, 8, 5, 7}, 5))		//	2, 4
	fmt.Println(subarraySum([]int{1, 1, 1}, 3))					//	0, 3
	fmt.Println(subarraySum([]int{1, -20, -3, 30, 5, 7}, 7))	//	1, 4
}