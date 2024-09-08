package main

import (
	"fmt"
)

func main() {

	fmt.Printf("%v", canPlaceFlowers([]int{0, 0, 1, 0, 0}, 3))
	//	math.Pow(2, 10)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			// Checking if left plot is empty or out of bounds
			left := (i == 0) || (flowerbed[i-1] == 0)
			// Checking if right plot is empty or out of bounds
			right := (i == len(flowerbed)-1) || (flowerbed[i+1] == 0)

			// If both left and right are empty, plant a flower
			if left && right {
				flowerbed[i] = 1
				n--
				// If no more flowers need to be planted, return true
				if n == 0 {
					return true
				}
			}
		}
	}
	// If after the loop there are still flowers to plant, return false
	return n <= 0

}
