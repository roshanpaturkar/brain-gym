package main

import "fmt"

func bubbleSort(unsoertedList []int) []int {
	n := len(unsoertedList)

	for i:=n-1; i>=0; i-- {
		swap := false;

		for j:=0; j<i; j++ {
			if unsoertedList[j] > unsoertedList[j+1] {
				unsoertedList[j], unsoertedList[j+1] = unsoertedList[j+1], unsoertedList[j]
				swap = true
			}
		}

		if !swap {
			return unsoertedList
		}
	}

	return unsoertedList
}

func main() {
	unsortedList := []int{5,7,2,5,1,4,8,3,10,-1,45,100,34,22,65,87,65,0}
	
	sortedList := bubbleSort(unsortedList)
	fmt.Println("Sorted list by bubble sort: ", sortedList)
}