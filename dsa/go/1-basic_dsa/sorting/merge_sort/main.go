package main

import "fmt"

func sortList(unsortedList []int) []int {
	n := len(unsortedList)

    // Base case: A list of size 1 or 0 is already sorted
    if n <= 1 {
        return unsortedList
    }

    // Split the list into left and right halves
    midpoint := n / 2
    leftList := sortList(unsortedList[:midpoint])
    rightList := sortList(unsortedList[midpoint:])

    resultList := make([]int, 0, n)
    leftPointer, rightPointer := 0, 0

    // Merge the sorted left and right lists with two pointers
    for leftPointer < midpoint || rightPointer < n-midpoint {
        if leftPointer == midpoint { // If left list is empty, append element from right
            resultList = append(resultList, rightList[rightPointer])
            rightPointer++
        } else if rightPointer == n-midpoint { // If right list is empty, append element from left
            resultList = append(resultList, leftList[leftPointer])
            leftPointer++
        } else if leftList[leftPointer] <= rightList[rightPointer] { // Append smaller element from left
            resultList = append(resultList, leftList[leftPointer])
            leftPointer++
        } else { // Append smaller element from right
            resultList = append(resultList, rightList[rightPointer])
            rightPointer++
        }
    }

    return resultList
}

func main() {
	unsortedList := []int{5,7,2,5,1,4,8,3,10,-1,45,100,34,22,65,87,65,0}
	
	sortedList := sortList(unsortedList)
	fmt.Println("Sorted list by merge sort: ", sortedList)
}