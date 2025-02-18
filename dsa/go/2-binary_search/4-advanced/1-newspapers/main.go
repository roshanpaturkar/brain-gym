// You've begun your new job to organize newspapers. Each morning, you are to separate the newspapers into smaller piles and assign each pile to a co-worker. This way, your co-workers can read through the newspapers and examine their contents simultaneously.

// Each newspaper is marked with a read time to finish all its contents. A worker can read one newspaper at a time, and, when they finish one, they can start reading the next. Your goal is to minimize the amount of time needed for your co-workers to finish all newspapers. Additionally, the newspapers came in a particular order, and you must not disarrange the original ordering when distributing the newspapers. In other words, you cannot pick and choose newspapers randomly from the whole pile to assign to a co-worker, but rather you must take a subsection of consecutive newspapers from the whole pile.

// What is the minimum amount of time it would take to have your coworkers go through all the newspapers? That is, if you optimize the distribution of newspapers, what is the longest reading time among all piles?

// Constraints
// 1 <= newspapers_read_times.length <= 10^5

// 1 <= newspapers_read_times[i] <= 10^5

// 1 <= num_coworkers <= 10^5

// Examples
// Example 1:
// Input: newspapers_read_times = [7,2,5,10,8], num_coworkers = 2
// Output: 18
// Explanation:
// Assign first 3 newspapers to one coworker then assign the rest to another. The time it takes for the first 3 newspapers is 7 + 2 + 5 = 14 and for the last 2 is 10 + 8 = 18.

// Example 2:
// Input: newspapers_read_times = [2,3,5,7], num_coworkers = 3
// Output: 7
// Explanation:
// Assign [2, 3], [5], and [7] separately to workers. The minimum time is 7.


package main

import (
	"fmt"
	"math"
)

func feasible(newspapersReadTimes []int, numCoworkers, limit int) bool {
    // time to keep track of the current worker's time spent, numWorkers to keep track of the number of coworkers used
    time, numWorkers := 0, 0
    for _, readTime := range newspapersReadTimes {
        // check if current time exceeds the given time limit
        if time+readTime > limit {
            time = 0
            numWorkers++
        }
        time += readTime
    }
    // edge case to check if we needed an extra worker at the end
    if time != 0 {
        numWorkers++
    }
    // check if the number of workers we need is more than what we have
    return numWorkers <= numCoworkers
}

func newspapersSplit(newspapersReadTimes []int, numCoworkers int) int {
    // Initializing 'low' to the maximum integer value and 'high' to 0
    low, high := math.MinInt32, 0
    // Finding the minimum value in the array and accumulating the total sum
    for _, readTime := range newspapersReadTimes {
        low = int(math.Max(float64(low), float64(readTime)))
        high += readTime
    }
    // Using binary search to find the optimal time limit
    ans := -1
    for low <= high {
        mid := (low + high) / 2
        // helper function to check if a time works
        if feasible(newspapersReadTimes, numCoworkers, mid) {
            ans = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return ans
}

func main() {
	newspapers_read_times := []int{7,2,5,10,8}
	num_coworkers := 2
	fmt.Println(newspapersSplit(newspapers_read_times, num_coworkers))

	newspapers_read_times = []int{2,3,5,7}
	num_coworkers = 3
	fmt.Println(newspapersSplit(newspapers_read_times, num_coworkers))
}