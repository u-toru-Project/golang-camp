package chapter11

import "fmt"

// CountUniqueCharsBuggy ignores non-letters then adds all chars again.
func CountUniqueCharsBuggy(value string) int {
	seen := make(map[rune]struct{})
	letters := 0
	for _, ch := range value {
		if isLetter(ch) {
			lower := toLower(ch)
			if _, ok := seen[lower]; !ok {
				seen[lower] = struct{}{}
				letters++
			}
		}
	}
	return letters + len([]rune(value))
}

// CountUniqueChars counts distinct runes.
func CountUniqueChars(value string) int {
	seen := make(map[rune]struct{})
	for _, ch := range value {
		seen[ch] = struct{}{}
	}
	return len(seen)
}

func isLetter(ch rune) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')
}

func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + ('a' - 'A')
	}
	return ch
}

func RunQ1101() {
	fmt.Printf("fixed aabBC=%d\n", CountUniqueChars("aabBC"))
	fmt.Printf("buggy !!=%d fixed=%d\n", CountUniqueCharsBuggy("!!"), CountUniqueChars("!!"))
}
