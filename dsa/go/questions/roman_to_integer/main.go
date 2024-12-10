package main

import "fmt"

func roman_to_integer(s string) int {
	roman_map := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && roman_map[s[i]] < roman_map[s[i+1]] {
			result -= roman_map[s[i]]
		} else {
			result += roman_map[s[i]]
		}
	}
	return result
}

func main() {
	s := "III"
	fmt.Println(roman_to_integer(s))

	s = "IV"
	fmt.Println(roman_to_integer(s))

	s = "IX"
	fmt.Println(roman_to_integer(s))

	s = "LVIII"
	fmt.Println(roman_to_integer(s))

	s = "MCMXCIV"
	fmt.Println(roman_to_integer(s))

	s = "MMXXI"
	fmt.Println(roman_to_integer(s))

	s = "MMXX"
	fmt.Println(roman_to_integer(s))

	s = "DCXXI"
	fmt.Println(roman_to_integer(s))
}