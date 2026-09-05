package chapter16

import "fmt"

func IsMatch(text, pattern string) bool {
	return match(text, pattern, 0, 0)
}

func match(text, pattern string, textIndex, patternIndex int) bool {
	if patternIndex == len(pattern) {
		return textIndex == len(text)
	}
	isStar := patternIndex+1 < len(pattern) && pattern[patternIndex+1] == '*'
	if isStar {
		ch := pattern[patternIndex]
		for {
			if match(text, pattern, textIndex, patternIndex+2) {
				return true
			}
			if textIndex == len(text) {
				break
			}
			if ch != '.' && text[textIndex] != ch {
				break
			}
			textIndex++
		}
		return false
	}
	if textIndex == len(text) {
		return false
	}
	if pattern[patternIndex] != '.' && pattern[patternIndex] != text[textIndex] {
		return false
	}
	return match(text, pattern, textIndex+1, patternIndex+1)
}

func RunQ1618() {
	fmt.Println(IsMatch("aab", "c*a*b"))
	fmt.Println(IsMatch("aa", "a"))
}
