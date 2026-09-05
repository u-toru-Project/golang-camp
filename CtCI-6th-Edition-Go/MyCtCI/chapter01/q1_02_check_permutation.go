package chapter01

import (
	"fmt"
	"slices"
)

func IsPermutationBySorting(original, valueToTest string) bool {
	if len(original) != len(valueToTest) {
		return false
	}
	left := []rune(original)
	right := []rune(valueToTest)
	slices.Sort(left)
	slices.Sort(right)
	return string(left) == string(right)
}

func IsPermutationByCounts(original, valueToTest string) bool {
	if len(original) != len(valueToTest) {
		return false
	}
	counts := make(map[rune]int)
	for _, character := range original {
		counts[character]++
	}
	for _, character := range valueToTest {
		if counts[character] == 0 {
			return false
		}
		counts[character]--
	}
	return true
}

func RunQ102() {
	pairs := [][2]string{{"apple", "papel"}, {"carrot", "tarroc"}, {"hello", "llloh"}}
	for _, pair := range pairs {
		fmt.Printf("%s, %s: %t / %t\n", pair[0], pair[1], IsPermutationBySorting(pair[0], pair[1]), IsPermutationByCounts(pair[0], pair[1]))
	}
}
