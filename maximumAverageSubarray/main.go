package main

import "fmt"

func main() {
	fmt.Printf("%.5f", findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)) // Example case
	fmt.Printf("\n%.5f", findMaxAverage([]int{-1}, 1))                 // Test case with [-1] and k=1
}

func findMaxAverage(nums []int, k int) float64 {

	// Initialize the sum of the first `k` elements
	var sum float64
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}

	// Set the initial maximum average as the average of the first window
	maxAvg := sum / float64(k)

	// Slide the window across the array
	for i := k; i < len(nums); i++ {
		// Update the sum by removing the first element of the previous window
		// and adding the new element of the current window
		sum += float64(nums[i]) - float64(nums[i-k])
		// Calculate the new average
		newAvg := sum / float64(k)
		// Update maxAvg if the new average is higher
		if newAvg > maxAvg {
			maxAvg = newAvg
		}
	}

	return maxAvg
}
