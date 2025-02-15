package main

import "fmt"


func selectionSort(unsortedList []int) []int {
	n := len(unsortedList)

	for i:=0; i<n; i++ {
		minIndex := i;

		for j:=i; j<n; j++ {
			if unsortedList[j] < unsortedList[minIndex] {
				minIndex = j
			}
			unsortedList[i] = unsortedList[minIndex]
		} 
	}

	return unsortedList
}

func main() {
	unsortedList := []int{5,7,2,5,1,4,8,3}
	
	sortedList := selectionSort(unsortedList)
	fmt.Println("Sorted list by selection sort: ", sortedList)
}