package main

import (
	"fmt"
)

func main() {

	fmt.Printf("%v", canPlaceFlowers([]int{0, 0, 1, 0, 0}, 1))
	//	math.Pow(2, 10)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	flag := false

	length := len(flowerbed)
	if n == 0 {
		return true
	} else {
		if length == 1 && flowerbed[0] == 0 && n == 1 {
			return true
		} else if length == 1 && flowerbed[0] == 1 && n == 1 {
			return false
		}
	}

	if len(flowerbed) > 1 {

		for i := 0; i < len(flowerbed); {
			if n == 0 {
				break
			}

			if i == len(flowerbed)-1 && flowerbed[i] == 0 && flowerbed[i-1] == 0 {
				n--
				flowerbed[i] = 1
				flag = true
				break
			}

			if flowerbed[i] == 1 {
				i += 2
				continue
			} else if flowerbed[i+1] == 0 {
				n--
				flowerbed[i] = 1
				flag = true
				i += 2
			} else {
				i++
			}

		}
	}

	if n == 0 {
		return flag
	} else {
		return false
	}

}
