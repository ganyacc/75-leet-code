package main

import "fmt"

func main() {
	moveZeroes([]int{1, 0, 2, 0})

}

func moveZeroes(nums []int) {

	j := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] != 0 {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			j++
		}
	}

	fmt.Println("", nums)

}
