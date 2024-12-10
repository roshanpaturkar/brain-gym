package main

import (
	"fmt"
)

// Given an integer, convert it to a roman numeral.

func int_to_roman(num int) string {
	roman_map := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	result := ""
	for num > 0 {
		if num >= 1000 {
			result += roman_map[1000]
			num -= 1000
		} else if num >= 900 {
			result += roman_map[900]
			num -= 900
		} else if num >= 500 {
			result += roman_map[500]
			num -= 500
		} else if num >= 400 {
			result += roman_map[400]
			num -= 400
		} else if num >= 100 {
			result += roman_map[100]
			num -= 100
		} else if num >= 90 {
			result += roman_map[90]
			num -= 90
		} else if num >= 50 {
			result += roman_map[50]
			num -= 50
		} else if num >= 40 {
			result += roman_map[40]
			num -= 40
		} else if num >= 10 {
			result += roman_map[10]
			num -= 10
		} else if num >= 9 {
			result += roman_map[9]
			num -= 9
		} else if num >= 5 {
			result += roman_map[5]
			num -= 5
		} else if num >= 4 {
			result += roman_map[4]
			num -= 4
		} else {
			result += roman_map[1]
			num -= 1
		}
	}
	return result
}

func main() {
	num := 3
	fmt.Println(int_to_roman(num))

	num = 4
	fmt.Println(int_to_roman(num))

	num = 9
	fmt.Println(int_to_roman(num))

	num = 58
	fmt.Println(int_to_roman(num))

	num = 1994
	fmt.Println(int_to_roman(num))

	num = 2021
	fmt.Println(int_to_roman(num))

	num = 2020
	fmt.Println(int_to_roman(num))

	num = 621
	fmt.Println(int_to_roman(num))
}