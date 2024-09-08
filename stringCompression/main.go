package main

import (
	"fmt"
	"strconv"
)

func compress(input []byte) int {
	if len(input) == 0 {
		return 0
	}

	writeIndex := 0
	count := 1

	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			input[writeIndex] = input[i-1] // Add the character
			writeIndex++

			if count > 1 {
				// Convert count to string, then append each byte of the count string
				for _, digit := range []byte(strconv.Itoa(count)) {
					input[writeIndex] = digit
					writeIndex++
				}
			}
			count = 1
		}
	}

	// Handle the last group
	input[writeIndex] = input[len(input)-1]
	writeIndex++
	if count > 1 {
		// Convert count to string, then append each byte of the count string
		for _, digit := range []byte(strconv.Itoa(count)) {
			input[writeIndex] = digit
			writeIndex++
		}
	}

	// Resize the input slice to match the new size
	fmt.Println("--", input[:writeIndex])
	return len(input[:writeIndex])
}

func main() {
	input := []byte{'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'a'}
	compressed := compress(input)
	fmt.Println("Compressed array", compressed)
}
