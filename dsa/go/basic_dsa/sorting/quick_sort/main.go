package main

import "fmt"


func sortListInterval(unsortedList []int, start, end int) {
    // If segment is 1 or 0, it's sorted
    if end-start <= 1 {
        return
    }

    // Using last element as the pivot
    pivot := unsortedList[end-1]
    startPtr, endPtr := start, end-1

    // Partitioning process
    for startPtr < endPtr {
        // Find the next element from the left that is larger than the pivot
        for unsortedList[startPtr] < pivot && startPtr < endPtr {
            startPtr++
        }

        // Find the next element from the right that is smaller than or equal to the pivot
        for unsortedList[endPtr] >= pivot && startPtr < endPtr {
            endPtr--
        }

        // Swap if pointers haven't met
        unsortedList[startPtr], unsortedList[endPtr] = unsortedList[endPtr], unsortedList[startPtr]
    }

    // Place pivot in its final position
    unsortedList[startPtr], unsortedList[end-1] = unsortedList[end-1], unsortedList[startPtr]

    // Sort left and right of the pivot
    sortListInterval(unsortedList, start, startPtr)
    sortListInterval(unsortedList, startPtr+1, end)
}

func sortList(unsortedList []int) []int {
    sortListInterval(unsortedList, 0, len(unsortedList))

    return unsortedList
}

func main() {
	unsortedList := []int{5,7,2,5,1,4,8,3,10,-1,45,100,34,22,65,87,65,0}
	
	sortedList := sortList(unsortedList)
	fmt.Println("Sorted list by quick sort: ", sortedList)
}