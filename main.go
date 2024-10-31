package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {
	vowelCount := 0
	consonantCount := 0

	for _, c := range str {
		cc := strings.ToLower(string(c))
		if cc == "a" || cc == "i" || cc == "u" || cc == "e" || cc == "o" {
			vowelCount++
		} else if cc >= "a" && cc <= "z" { // Check for consonants
			consonantCount++
		}
	}

	isValid := vowelCount == 0 || consonantCount == 0
	return vowelCount, consonantCount, isValid
}

func main() {
	// Test cases
	fmt.Println(CountVowelConsonant("kopi"))           // Expected output: 2, 2, false
	fmt.Println(CountVowelConsonant("bbbbb ccccc"))    // Expected output: 0, 10, true
	fmt.Println(CountVowelConsonant("Hidup Itu Indah")) // Expected output: 6, 7, false
}
