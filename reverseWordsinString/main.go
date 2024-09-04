package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	stack := []string{}
	var wordBuilder strings.Builder

	for _, char := range s {
		if char == ' ' {
			if wordBuilder.Len() > 0 {
				stack = append(stack, wordBuilder.String())
				wordBuilder.Reset()
			}
		} else {
			wordBuilder.WriteRune(char)
		}

	}

	if wordBuilder.Len() > 0 {
		stack = append(stack, wordBuilder.String())
	}

	var result strings.Builder
	for i := len(stack) - 1; i >= 0; i-- {

		result.WriteString(stack[i])
		if i > 0 {
			result.WriteString(" ")
		}
	}

	return result.String()
}

func main() {
	str := "the sky is blue"
	reversed := reverseWords(str)
	fmt.Println(reversed) // Output: "Go is this world hello"
}
