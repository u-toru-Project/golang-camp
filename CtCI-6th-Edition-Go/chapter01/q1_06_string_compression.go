package chapter01

import (
	"fmt"
	"strconv"
	"strings"
)

// Compress returns the original when compression is not strictly shorter.
// Time O(n), space O(n).
func Compress(value string) string {
	if value == "" {
		return value
	}
	var b strings.Builder
	count := 1
	for i := 1; i < len(value); i++ {
		if value[i] == value[i-1] {
			count++
			continue
		}
		b.WriteByte(value[i-1])
		b.WriteString(strconv.Itoa(count))
		count = 1
	}
	b.WriteByte(value[len(value)-1])
	b.WriteString(strconv.Itoa(count))
	if b.Len() < len(value) {
		return b.String()
	}
	return value
}

func RunQ106() {
	original := "abbccccccde"
	fmt.Printf("Original  : %s\nCompressed: %s\n", original, Compress(original))
}
