// Date: 2021/10/10
// Palindrome Number
// Given an integer x, return true if x is palindrome integer.
// An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.

package main

import "fmt"

func palindrome_number(x int) bool {
	if x < 0 {
		return false
	}
	// reverse the number
	rev := 0
	original := x
	for x != 0 {
		rev = rev*10 + x%10
		x /= 10
	}
	return original == rev
}

func main() {
	x := 121
	fmt.Println(palindrome_number(x))

	x = -121
	fmt.Println(palindrome_number(x))

	x = 10
	fmt.Println(palindrome_number(x))

	x = -101
	fmt.Println(palindrome_number(x))
}