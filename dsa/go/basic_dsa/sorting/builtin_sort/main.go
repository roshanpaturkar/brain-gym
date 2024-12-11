package main

import (
	"fmt"
	"sort"
)

func main() {
	// sort words alphabetically
	words := []string{"zebra", "fat", "apple", "lion", "ink"}
	sort.Strings(words) // words = ["apple", "fat", "ink", "lion", "zebra"]
	fmt.Println("Sorted Strings :", words)

	// sort nums in ascending order
	nums := []int{40, 100, 1, 5, 25, 10}
	sort.Ints(nums) // nums = [1, 5, 10, 25, 40, 100]
	fmt.Println("Sorted Nums :", nums)

	// sort words in descending order
	sort.Sort(sort.Reverse(sort.StringSlice(words)))
	fmt.Println("Sorted string slice in decending order: ", words)

	// sort nums in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // nums = [100, 40, 25, 10, 5, 1]
	fmt.Println("Sorted number in decending order :", nums)

	// sort float64 in ascending order
	floats := []float64{40.1, 100.5, 1.2, 5.6, 25.3, 10.4}
	sort.Float64s(floats) // floats = [1.2, 5.6, 10.4, 25.3, 40.1, 100.5]
	fmt.Println("Sorted Floats :", floats)

	// sort float64 in descending order
	sort.Sort(sort.Reverse(sort.Float64Slice(floats))) // floats = [100.5, 40.1, 25.3, 10.4, 5.6, 1.2]
	fmt.Println("Sorted Floats in decending order :", floats)

	// sort tasks by priority in ascending order
	tasks := []struct {
		description string
		priority    int
	}{
		{"Cook dinner", 5},
		{"Buy grocery", 3},
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].priority < tasks[j].priority
	})
	// tasks = [{"Buy grocery", 3}, {"Cook dinner", 5}]
	fmt.Println("Sorted Structures accroding to priority: ", tasks)
}
