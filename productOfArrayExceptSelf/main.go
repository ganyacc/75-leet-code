package main

import "fmt"

func main() {

	fmt.Printf("%v", productExceptSelf([]int{1, 2, 3, 4}))

}

func productExceptSelf(nums []int) []int {

	n := len(nums)
	result := make([]int, n)

	// Step 1: Compute the prefix product for each element
	result[0] = 1 // There is no element before the first element
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Step 2: Compute the suffix product and multiply it with the prefix product
	suffixProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] = result[i] * suffixProduct
		suffixProduct *= nums[i]
	}

	return result

}
