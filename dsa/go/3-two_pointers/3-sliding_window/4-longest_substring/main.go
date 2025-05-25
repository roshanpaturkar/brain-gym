package main

import (
	"fmt"
)

func longestSubstringWithoutRepeatingCharacters(s string) int {
	longest := 0
	left, right := 0, 0
    sub := make(map[rune]int)
    
    for right < len(s) {
        sub[rune(s[right])] ++
        
        for sub[rune(s[right])] > 1 {
            sub[rune(s[left])] --
            left ++
        }
        longest = max(longest, right-left+1)
        right ++
    }
    return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestSubstringWithoutRepeatingCharacters("abccabcabcc")) // Output: 3
	fmt.Println(longestSubstringWithoutRepeatingCharacters("aaaabaaa"))    // Output: 2
}