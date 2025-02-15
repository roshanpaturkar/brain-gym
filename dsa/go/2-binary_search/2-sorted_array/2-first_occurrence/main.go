package main

import (
	"fmt"
)

func findFirstOccurrence(arr []int, target int) int {
    l := 0
    r := len(arr) - 1
    ans := -1
    for l <= r {
        mid := (l + r) / 2
        if arr[mid] == target {
            ans = mid
            r = mid - 1
        } else if arr[mid] < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return ans
}

func main() {
	arr := []int{1, 2, 3, 3, 3, 3, 3, 3, 3, 4, 5, 6, 7, 8, 9}
	target := 3
	fmt.Println(findFirstOccurrence(arr, target))
	
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target = 4
	fmt.Println(findFirstOccurrence(arr, target))
}