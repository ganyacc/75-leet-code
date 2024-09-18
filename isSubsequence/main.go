package main

import (
	"fmt"
)

func main() {

	fmt.Printf("%v", isSubsequence("acb", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	substring := []rune(s)
	mainString := []rune(t)

	if len(s) == 0 {
		return true
	}

	idx := 0
	flag := false
	for i := 0; i < len(substring); i++ {

		for j := idx; j < len(mainString); j++ {

			if substring[i] == mainString[j] && i <= idx {
				idx = j
				idx++
				flag = true
				break
			} else {
				flag = false
			}
		}

		if flag == false {
			return flag
		}

		if idx == len(mainString) && i < len(substring)-1 {
			return false
		}
	}
	return flag
}
