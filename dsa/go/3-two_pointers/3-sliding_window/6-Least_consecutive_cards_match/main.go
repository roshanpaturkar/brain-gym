// A bunch of cards is laid out in front of you in a line, where the value of each card ranges from 0 to 10^6. 
// A pair of cards is matching if they have the same number value.
// Given a list of integers cards, your goal is to match a pair of cards, but you can only pick up cards in a consecutive manner.
// What's the minimum number of cards that you need to pick up to make a pair? If there are no matching pairs, return -1.
// For example, given cards = [3, 4, 2, 3, 4, 7], then picking up [3, 4, 2, 3] makes a pair of 3s and picking up [4, 2, 3, 4] matches two 4s.
// We need 4 consecutive cards to match a pair of 3s and 4 consecutive cards to match 4s, so you need to pick up at least 4 cards to make a match.

package main

import (
	"fmt"
)

func leastConsecutiveCardsToMatch(cards []int) int {
    shortest := len(cards) + 1
    counts := make(map[int]int)
    left, right := 0, 0
    
    for right < len(cards) {
        counts[cards[right]] ++
        
        if counts[cards[right]] == 2 {
            shortest = min(shortest, right - left + 1)
            counts[cards[left]] --
            left ++
        }
        right ++
    }
    
    if shortest == len(cards) + 1 {
        return -1
    }
    return shortest
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
	fmt.Println(leastConsecutiveCardsToMatch([]int{3, 4, 2, 3, 4, 7})) // Output: 4
	fmt.Println(leastConsecutiveCardsToMatch([]int{1, 0, 5, 3})) // Output: -1
	fmt.Println(leastConsecutiveCardsToMatch([]int{5})) // Output: -1
	fmt.Println(leastConsecutiveCardsToMatch([]int{7, 7})) // Output: 2
}