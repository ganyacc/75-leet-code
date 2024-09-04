package main

import "fmt"

func reverseVowels(s string) string {
	if len(s) <= 1 {
		return s
	}

	vowels := "aAeEiIoOuU"
	strRune := []rune(s)
	j := len(strRune) - 1

	for i := 0; i < len(strRune); i++ {
		if isVowel(strRune[i], vowels) {
			for j > i && !isVowel(strRune[j], vowels) {
				j--
			}
			if j > i {
				strRune[i], strRune[j] = strRune[j], strRune[i]
				j--
			}
		}
	}

	return string(strRune)
}

func isVowel(r rune, vowels string) bool {
	for _, v := range vowels {
		if r == v {
			return true
		}
	}
	return false
}

func main() {
	str := "leetcode"
	swapped := reverseVowels(str)
	fmt.Println(swapped)
}
