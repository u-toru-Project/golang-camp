package chapter01

import (
	"fmt"
	"unicode"
)

func toLetterIndex(character rune) int {
	letter := unicode.ToLower(character)
	if letter < 'a' || letter > 'z' {
		return -1
	}
	return int(letter - 'a')
}

// IsPermutationOfPalindrome ignores non-letters and casing.
// Time O(n), space O(1).
func IsPermutationOfPalindrome(phrase string) bool {
	table := make([]int, 26)
	for _, character := range phrase {
		if index := toLetterIndex(character); index != -1 {
			table[index]++
		}
	}
	foundOdd := false
	for _, count := range table {
		if count%2 != 1 {
			continue
		}
		if foundOdd {
			return false
		}
		foundOdd = true
	}
	return true
}

// IsPermutationOfPalindrome2 tracks the odd count as it goes.
func IsPermutationOfPalindrome2(phrase string) bool {
	countOdd := 0
	table := make([]int, 26)
	for _, character := range phrase {
		index := toLetterIndex(character)
		if index == -1 {
			continue
		}
		table[index]++
		if table[index]%2 == 1 {
			countOdd++
		} else {
			countOdd--
		}
	}
	return countOdd <= 1
}

// IsPermutationOfPalindrome3 uses a bit vector of odd/even counts.
func IsPermutationOfPalindrome3(phrase string) bool {
	bitVector := 0
	for _, character := range phrase {
		if index := toLetterIndex(character); index >= 0 {
			bitVector ^= 1 << index
		}
	}
	return bitVector&(bitVector-1) == 0
}

func RunQ104() {
	for _, phrase := range []string{"Tact Coa", "asda"} {
		fmt.Printf("%s: %t / %t / %t\n", phrase, IsPermutationOfPalindrome(phrase), IsPermutationOfPalindrome2(phrase), IsPermutationOfPalindrome3(phrase))
	}
}
