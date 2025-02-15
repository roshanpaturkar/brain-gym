// take an integer array and returns a hash table with the array elements as keys and their frequencies as values.

package main

import "fmt"

func counter(arr []int) map[int]int {
	counter := make(map[int]int)
    for _, num := range arr {
        if _, exists := counter[num]; exists {
            counter[num]++
        } else {
            counter[num] = 1
        }
    }
    return counter
}

func main() {
	arr := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	result := counter(arr)
	fmt.Println(result)
}

