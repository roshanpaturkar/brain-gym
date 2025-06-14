package main

import "fmt"

func getMinimumWindow(original string, check string) string {
	// If either string is empty, return an empty string
	// This is a base case to handle edge conditions
	// where no valid substring can be formed
	// that contains all characters of 'check'
	// in the original string 'original'.
	// This is important to avoid unnecessary processing
	// and to ensure that we return a valid result.
	if len(original) == 0 || len(check) == 0 {
		return ""
	}

	// Create a map to count the occurrences of each character in 'check'
	checkCount := make(map[rune]int)
	for _, char := range check {
		checkCount[char]++
	}

	// Two pointers to maintain the sliding window
	left, right := 0, 0
	minLength := len(original) + 1
	minSubstring := ""

	// Map to keep track of the count of characters in the current window
	// and the number of matched characters
	currentCount := make(map[rune]int)
	matchedChars := 0

	for right < len(original) {
		// Expand the window by including the character at 'right'
		currentChar := rune(original[right])
		currentCount[currentChar]++

		// If the current character's count matches the required count, increment matchedChars
		if currentCount[currentChar] == checkCount[currentChar] {
			matchedChars++
		}

		// Try to contract the window from the left
		for matchedChars == len(checkCount) {
			// Update the minimum substring if the current window is smaller
			// and lexicographically smaller if lengths are equal
			if right-left+1 < minLength || (right-left+1 == minLength && original[left:right+1] < minSubstring) {
				minLength = right - left + 1
				minSubstring = original[left : right+1]
			}

			// Contract the window by moving 'left' pointer
			leftChar := rune(original[left])
			currentCount[leftChar]--
			if currentCount[leftChar] < checkCount[leftChar] {
				matchedChars--
			}
			left++
		}
		right++
	}

	return minSubstring
}

func main() {
	original := "cdbaebaecd"
	check := "abc"
	result := getMinimumWindow(original, check)
	fmt.Println("Minimum window substring:", result) // Output: "baec"

	original = "aabbcc"
	check = "abc"
	result = getMinimumWindow(original, check)
	fmt.Println("Minimum window substring:", result) // Output: "abbc"

	original = "aabbcc"
	check = "xyz"
	result = getMinimumWindow(original, check)
	fmt.Println("Minimum window substring:", result) // Output: ""
}