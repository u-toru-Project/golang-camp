package chapter05

import (
	"fmt"
	"strings"
)

// PrintBinary prints the fractional binary, or ERROR if it needs more than 32 bits.
// Time O(1), space O(1).
func PrintBinary(number float64) string {
	if number >= 1 || number <= 0 {
		return "ERROR"
	}
	var binary strings.Builder
	binary.WriteByte('.')
	for number > 0 {
		if binary.Len() > 32 {
			return "ERROR"
		}
		doubled := number * 2
		if doubled >= 1 {
			binary.WriteByte('1')
			number = doubled - 1
		} else {
			binary.WriteByte('0')
			number = doubled
		}
	}
	return binary.String()
}

// PrintBinary2 uses successive fractions 0.5, 0.25, ...
// Time O(1), space O(1).
func PrintBinary2(number float64) string {
	if number >= 1 || number <= 0 {
		return "ERROR"
	}
	var binary strings.Builder
	binary.WriteByte('.')
	fraction := 0.5
	for number > 0 {
		if binary.Len() >= 32 {
			return "ERROR"
		}
		if number >= fraction {
			binary.WriteByte('1')
			number -= fraction
		} else {
			binary.WriteByte('0')
		}
		fraction /= 2
	}
	return binary.String()
}

func RunQ502() {
	fmt.Printf("0.125 -> %s / %s\n", PrintBinary(0.125), PrintBinary2(0.125))
}
