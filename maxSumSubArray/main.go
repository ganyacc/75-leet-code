package main

import (
	"fmt"
)

func main() {

	arr := []int{2, 1, 5, 1, 3, 2}
	fmt.Printf("max sum :%d", maxSubarraySum(arr, 8))
}

func maxSubarraySum(array []int, k int) int {

	if k > len(array) {
		return 0
	}

	maxSum := 0
	windowSum := 0

	for i := 0; i < len(array); i++ {
		windowSum += array[i]

		if i >= k-1 {
			if windowSum > maxSum {
				maxSum = windowSum
			}
			windowSum -= array[i-(k-1)]
		}

	}
	return maxSum
}
