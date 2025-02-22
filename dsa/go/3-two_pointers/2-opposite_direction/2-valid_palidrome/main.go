// Determine whether a string is a palindrome, ignoring non-alphanumeric characters and case. 
// Examples:
// Input: Do geese see God? Output: True
// Input: Was it a car or a cat I saw? Output: True
// Input: A brown fox jumping over Output: False

package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    left := 0
    right := len(s) - 1
    
    for left < right {
        for left < right && !isAlphanumeric(s[left]) {
            left ++
        }
        
        for left < right && !isAlphanumeric(s[right]) {
            right --
        }
        
        if s[left] != s[right] {
            return false
        }
        
        left ++
		right --
    }
    
    return true
}

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func main() {
	s := "Do geese see God?"
	fmt.Println(isPalindrome(s))
	
	s = "Was it a car or a cat I saw?"
	fmt.Println(isPalindrome(s))
	
	s = "A brown fox jumping over"
	fmt.Println(isPalindrome(s))
}