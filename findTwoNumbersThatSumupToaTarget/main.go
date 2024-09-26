package main

import "fmt"

func main() {

	array := []int{1, 2, 4, 5, 6}
	fmt.Printf("target elements %v", twoSum(array, 9))
}

func twoSum(array []int, target int) []int {
	left := 0
	right := len(array) - 1

	result := []int{}
	for left < right {

		sum := array[left] + array[right]
		if sum == target {
			result = append(result, array[left], array[right])
			return result
		} else if sum < target {
			left++
		} else {
			right--
		}

	}
	return result
}
