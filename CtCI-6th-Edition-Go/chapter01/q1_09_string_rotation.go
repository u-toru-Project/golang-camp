package chapter01

import (
	"fmt"
	"strings"
)

func IsSubstring(value, possibleSubstring string) bool {
	return strings.Contains(value, possibleSubstring)
}

// IsRotation uses a single substring check on s1+s1.
// Time O(n), space O(n).
func IsRotation(first, second string) bool {
	if len(first) == 0 || len(first) != len(second) {
		return false
	}
	return IsSubstring(first+first, second)
}

func RunQ109() {
	pairs := [][2]string{{"waterbottle", "erbottlewat"}, {"camera", "macera"}}
	for _, pair := range pairs {
		fmt.Printf("%s, %s: %t\n", pair[0], pair[1], IsRotation(pair[0], pair[1]))
	}
}
