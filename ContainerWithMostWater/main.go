package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("%v", maxArea([]int{4, 3, 2, 1, 3}))

}

func maxArea(height []int) int {

	maxArea := 0.0
	i := 0
	j := len(height) - 1

	for i < j {
		currArea := float64((j - i)) * (math.Min(float64(height[i]), float64(height[j])))
		fmt.Println("cur", currArea)
		maxArea = math.Max(currArea, maxArea)
		if height[i] < height[j] {
			i = i + 1
		} else {
			j = j - 1
		}
	}
	return int(maxArea)
}
