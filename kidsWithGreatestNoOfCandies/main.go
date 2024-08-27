package main

import (
	"fmt"
)

func main() {
	result := kidsWithCandies([]int{12, 1, 12}, 10)
	if result != nil {
		fmt.Println(result)
	}
}

func findMax(list []int) int {
	if len(list) == 0 {
		return 0
	}

	max := 0

	for i := 0; i < len(list); i++ {

		if list[i] > max {
			max = list[i]
		}

	}
	return max

}

func kidsWithCandies(candies []int, extraCandies int) []bool {

	switch {

	case len(candies) < 2:
		fmt.Println("length of candies should be greater than or equal to 2")
		return nil

	case len(candies) > 100:
		fmt.Println("length of candies should be less than or equal to 100")
		return nil
	}

	max := findMax(candies)

	var result = make([]bool, len(candies))

	length := len(candies)

	for i := 0; i < length; i++ {
		if candies[i] > 100 {
			fmt.Printf("%dperson has more candies than 100 ", i)
			return nil
		} else if candies[i] < 1 {
			fmt.Printf("%d person has less candies than 1 ", i)
			return nil
		}

		//flag to insert final value to result array
		isMax := false

		num := candies[i] + extraCandies

		if num >= max {
			isMax = true
		} else {
			isMax = false
		}

		result[i] = isMax

	}

	return result

}
