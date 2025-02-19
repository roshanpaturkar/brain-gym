package main

import (
	"fmt"
)

func removeDuplicates(arr []int) int {
    slow := 0
//     for _, num := range(arr) {
//         if arr[slow] != num {
//             slow += 1
//             arr[slow] = num
//         }
//     }
 
    for fast :=1; fast < len(arr); fast++ {
        if arr[slow] != arr[fast] {
            slow += 1
            arr[slow] = arr[fast]
        }
    }
    return slow + 1
}

func main() {
	arr := []int{1, 1, 2, 2, 3, 4, 5, 5, 5, 6, 6, 7, 8, 9}
	fmt.Println(removeDuplicates(arr))
	fmt.Println(arr)
	
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(removeDuplicates(arr))
	fmt.Println(arr)
	
	arr = []int{1, 2, 3, 4, 5, 7, 8, 9}
	fmt.Println(removeDuplicates(arr))
	fmt.Println(arr)
}