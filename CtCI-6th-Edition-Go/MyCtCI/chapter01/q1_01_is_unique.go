package chapter01

import "fmt"

func IsUniqueChars(value string) bool {
	seen := make(map[rune]struct{})
	for _, character := range value {
		if _, ok := seen[character]; ok {
			return false
		}
		seen[character] = struct{}{}
	}
	return true
}

func IsUniqueAsciiLetters(value string) (bool, error) {
	checker := 0
	for _, character := range value {
		if character < 'a' || character > 'z' {
			return false, fmt.Errorf("only lowercase ASCII letters are supported")
		}
		mask := 1 << (character - 'a')
		if checker&mask != 0 {
			return false, nil
		}
		checker |= mask
	}
	return true, nil
}

func RunQ101() {
	for _, word := range []string{"abcde", "hello", "apple", "kite", "padle"} {
		ok, err := IsUniqueAsciiLetters(word)
		if err != nil {
			fmt.Printf("%s: %t %v\n", word, IsUniqueChars(word), err)
			continue
		}
		fmt.Printf("%s: %t %v\n", word, IsUniqueChars(word), ok)
	}
}
