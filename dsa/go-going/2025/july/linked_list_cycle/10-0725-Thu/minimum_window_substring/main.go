package main

import "fmt"

func getMinimumWindow(original string, check string) string {
	if len(original) == 0 || len(check) == 0 {
		return ""
	}

	checkCount := make(map[rune]int)
	for _, char := range check {
		checkCount[char]++
	}

	left, right := 0, 0
	minLength := len(original) + 1
	minSubstring := ""

	currentCount := make(map[rune]int)
	matchedChars := 0

	for right < len(original) {
		currentChar := rune(original[right])
		currentCount[currentChar]++

		if currentCount[currentChar] == checkCount[currentChar] {
			matchedChars++
		}

		for matchedChars == len(checkCount) {
			if right-left+1 < minLength || (right-left+1 == minLength && original[left:right+1] < minSubstring) {
				minLength = right - left + 1
				minSubstring = original[left : right+1]
			}

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