package main

import "fmt"

func maximumScore(arr1 []int, arr2 []int) int {
	const modulo_amt = 1_000_000_007
	
	// Initialize pointers for both arrays
	// and variables to keep track of the current sums
	// and the maximum score
	ptr1, ptr2 := 0, 0
	sum1, sum2 := 0, 0
	maxScore := 0

	for ptr1 < len(arr1) && ptr2 < len(arr2) {
		// Compare the current elements of both arrays
		// and decide which one to take
		// If they are equal, we can teleport
		// and add the maximum of the two sums
		// to the maxScore
		// and reset the sums for the next segment
		if arr1[ptr1] < arr2[ptr2] {
			sum1 += arr1[ptr1]
			ptr1++
		} else if arr1[ptr1] > arr2[ptr2] {
			sum2 += arr2[ptr2]
			ptr2++
		} else {
			maxScore += max(sum1, sum2) + arr1[ptr1]
			maxScore %= modulo_amt
			sum1, sum2 = 0, 0
			ptr1++
			ptr2++
		}
	}

	// If we reach the end of one array, we need to add the remaining 
	// elements of the other array
	for ptr1 < len(arr1) {
		sum1 += arr1[ptr1]
		ptr1++
	}

	// If we reach the end of the second array, we need to add the remaining
	// elements of the first array
	for ptr2 < len(arr2) {
		sum2 += arr2[ptr2]
		ptr2++
	}

	// Add the maximum of the two sums to the maxScore
	maxScore += max(sum1, sum2)
	// Ensure we take modulo to prevent overflow
	return maxScore % modulo_amt
}

func main() {
	arr1 := []int{2, 4, 5, 8, 10}
	arr2 := []int{4, 6, 8, 9}
	result := maximumScore(arr1, arr2)
	fmt.Println(result) // Output: 30

	arr1 = []int{1, 4, 5, 8, 9}
	arr2 = []int{2, 3, 6, 7, 10}
	result = maximumScore(arr1, arr2)
	fmt.Println(result) // Output: 28

	arr1 = []int{2, 3, 5, 6, 7, 9, 11, 13, 14, 16, 17, 19, 20}
	arr2 = []int{3, 4, 5, 7, 8, 10, 11, 12, 15, 16, 18, 20}
	result = maximumScore(arr1, arr2)
	fmt.Println(result) // Output: 155

	arr1 = []int{2, 4, 6, 7, 9, 15, 16, 17, 18, 20, 21, 22, 24, 29, 30, 33, 43}
	arr2 = []int{2, 3, 5, 8, 9, 10, 11, 12, 13, 14, 20, 23, 25, 26, 27, 30, 36, 39}
	result = maximumScore(arr1, arr2)
	fmt.Println(result) // Output: 321
}