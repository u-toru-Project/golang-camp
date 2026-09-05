package chapter17

import (
	"fmt"
	"unicode"
)

func LongestBalancedSubarray(text string) (int, int, int) {
	balance := 0
	firstSeen := map[int]int{0: -1}
	bestLength, bestStart, bestEnd := 0, 0, -1
	for i, ch := range text {
		if unicode.IsDigit(ch) {
			balance--
		} else if unicode.IsLetter(ch) {
			balance++
		}
		if firstIndex, ok := firstSeen[balance]; ok {
			length := i - firstIndex
			if length > bestLength {
				bestLength = length
				bestStart = firstIndex + 1
				bestEnd = i
			}
		} else {
			firstSeen[balance] = i
		}
	}
	return bestLength, bestStart, bestEnd
}

func RunQ1705() {
	text := "aaa222aa"
	length, start, end := LongestBalancedSubarray(text)
	fmt.Printf("%s -> len=%d [%d,%d]\n", text, length, start, end)
}
