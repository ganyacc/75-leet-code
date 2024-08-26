package main

import (
	"fmt"
)

func findGcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
		fmt.Println("a", a)
		fmt.Println("b", b)
	}

	return a
}

// Function to check if s can be constructed by repeating t
func canConstruct(s, t string) bool {
	if len(s)%len(t) != 0 {
		return false
	}
	repeated := ""
	for i := 0; i < len(s)/len(t); i++ {
		repeated += t
	}
	return repeated == s
}

// Function to find the largest string x that divides both str1 and str2
func gcdOfStrings(str1, str2 string) string {
	// Calculate the GCD of the lengths of the two strings
	len1, len2 := len(str1), len(str2)
	gcdLength := findGcd(len1, len2)

	// Potential gcd string is the substring of length gcdLength from str1
	potentialGcd := str1[:gcdLength]

	// Check if this potentialGcd can construct both str1 and str2
	if canConstruct(str1, potentialGcd) && canConstruct(str2, potentialGcd) {
		return potentialGcd
	}
	return ""
}

func main() {
	fmt.Println(gcdOfStrings("ABCABCAAAAAAAAAAAAAA", "ABAAAAAA")) // Output: "ABC"
	// fmt.Println(gcdOfStrings("ABABAB", "ABAB")) // Output: "AB"
	//fmt.Println(gcdOfStrings("LEET", "CODE"))   // Output: ""
}
