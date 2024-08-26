package main

import (
	"fmt"
)

func main() {

	var word1, word2 string

	fmt.Println("enter word1")
	fmt.Scanf("%v\n", &word1)

	fmt.Println("enter word2")
	fmt.Scanf("%v", &word2)

	fmt.Printf("%v", mergeAlternately1(word1, word2))

}

func mergeAlternately1(word1, word2 string) string {
	var mergedString []byte
	len1, len2 := len(word1), len(word2)
	maxLen := len1
	if len2 > len1 {
		maxLen = len2
	}

	for i := 0; i < maxLen; i++ {
		if i < len1 {
			mergedString = append(mergedString, word1[i])
		}
		if i < len2 {
			mergedString = append(mergedString, word2[i])
		}
	}

	return string(mergedString)
}
