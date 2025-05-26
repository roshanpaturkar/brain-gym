package main

import "fmt"

func leastConsecutiveCardToMatch(cards []int) int {
	left, right := 0, 0
	minimum := len(cards) + 1
	count := make(map[int]int)
	
	for right < len(cards) {
		count[cards[right]] ++
		
		if count[cards[right]] == 2 {
			minimum = min(minimum, right - left + 1)
			count[cards[left]]--
			left ++
		}
		right ++
	}
	
	if 	minimum == len(cards) + 1 {
		return -1
	}
	return minimum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cards := []int{3, 4, 2, 3, 4, 7}
	fmt.Println(leastConsecutiveCardToMatch(cards))		//	4
	
	fmt.Println(leastConsecutiveCardToMatch([]int{1, 0, 5, 3}))	//	-1
	fmt.Println(leastConsecutiveCardToMatch([]int{5}))				//	-1
	fmt.Println(leastConsecutiveCardToMatch([]int{7, 7})) 			//	2
}