package main

import "fmt"

func main() {

	fmt.Printf("%v", increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
}

func increasingTriplet(nums []int) bool {
	first, second := int(^uint(0)>>1), int(^uint(0)>>1) // Max int value
	fmt.Println("first", first)
	fmt.Println("second", second)

	for _, num := range nums {
		if num <= first {
			first = num // Update first to the smallest element
		} else if num <= second {
			second = num // Update second to the next smallest element
		} else {
			// If we find a number greater than both first and second
			return true
		}
	}

	return false // No such triplet found
}
