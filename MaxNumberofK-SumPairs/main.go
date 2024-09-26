package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Printf("%v", maxOperations([]int{3, 1, 3, 4, 3}, 6))
}

func maxOperations(nums []int, k int) int {

	// Sort the array first

	slices.Sort(nums)

	left, right := 0, len(nums)-1
	operations := 0

	// Use two-pointer approach
	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			operations++
			left++
			right--
		} else if sum < k {
			left++
		} else {
			right--
		}
	}

	return operations
}
