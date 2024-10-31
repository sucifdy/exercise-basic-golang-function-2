package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {
	vowels := "aiueoAIUEO"
	vowelCount := 0
	consonantCount := 0

	for _, char := range str {
		if char == ' ' {
			continue // abaikan spasi
		}
		if strings.ContainsRune(vowels, char) {
			vowelCount++
		} else {
			consonantCount++
		}
	}

	// Jika tidak ada vokal dan/atau konsonan
	noVowelsOrConsonants := (vowelCount == 0 || consonantCount == 0)
	return vowelCount, consonantCount, noVowelsOrConsonants
}

// Fungsi utama untuk menguji
func main() {
	fmt.Println(CountVowelConsonant("kopi"))         // 2, 2, false
	fmt.Println(CountVowelConsonant("bbbbb ccccc"))   // 0, 10, true
	fmt.Println(CountVowelConsonant("Hidup Itu Indah")) // 6, 7, false
}
