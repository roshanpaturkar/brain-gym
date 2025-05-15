package main

import "fmt"

func countChar(chars string) map[rune]int {
	charCount := make(map[rune]int)

	for _, char := range chars {
		_, exists := charCount[char]

		if exists {
			charCount[char] += 1
		} else {
			charCount[char] = 1
		}
	}

	return charCount
}

func main() {
	data := "hello world"

	counts := countChar(data)
	fmt.Println("Character counts:")
	for char, count := range counts {
		fmt.Printf("\"%c\": %d\n", char, count)
	}
}