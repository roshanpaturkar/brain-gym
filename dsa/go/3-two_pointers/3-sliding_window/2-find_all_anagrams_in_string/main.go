// Given a string original and a string check, find the starting index of all substrings of original that is an anagram of check.
// The output must be sorted in ascending order.

// Parameters
// original: A string
// check: A string
// Result
// A list of integers representing the starting indices of all anagrams of check.
// Examples
// Example 1
// Input: original = "cbaebabacd", check = "abc"
// Output: [0, 6]
// Explanation: The substring from 0 to 2, "cba", is an anagram of "abc", and so is the substring from 6 to 8, "bac".

// Example 2
// Input: original = "abab", check = "ab"
// Output: [0, 1, 2]
// Explanation: All substrings with length 2 from "abab" are anagrams of "ab".
	
// Constraints
// 1 <= len(original), len(check) <= 10^5
// Each string consists of only lowercase characters in the standard English alphabet.

package main

import (
	"fmt"
)

func findAnagrams(original string, check string) []int {
	// Create a map to store the frequency of each character in check
	checkMap := make(map[rune]int)
	for _, char := range check {
		checkMap[char]++
	}

	// Create a map to store the frequency of each character in the current window
	windowMap := make(map[rune]int)

	// Create a list to store the starting indices of all anagrams of check
	var result []int

	// Initialize the left and right pointers
	left, right := 0, 0

	// Iterate over the original string
	for right < len(original) {
		// Expand the window
		windowMap[rune(original[right])]++

		// If the window size is equal to the size of check
		if right-left+1 == len(check) {
			// Check if the window is an anagram of check
			if isAnagram(windowMap, checkMap) {
				// If it is an anagram, add the starting index of the window to the result
				result = append(result, left)
			}

			// Shrink the window
			windowMap[rune(original[left])]--
			if windowMap[rune(original[left])] == 0 {
				delete(windowMap, rune(original[left]))
			}

			// Move the left pointer
			left++
		}

		// Move the right pointer
		right++
	}

	return result
}

func isAnagram(windowMap map[rune]int, checkMap map[rune]int) bool {
	for char, count := range checkMap {
		if windowMap[char] != count {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(findAnagrams("cbaebabacd", "abc")) // [0 6]
	fmt.Println(findAnagrams("abab", "ab"))         // [0 1 2]
}