package chapter16

import (
	"fmt"
	"sort"
	"unicode"
)

var t9Letters = map[int]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

type t9Node struct {
	children map[rune]*t9Node
	word     string
}

func T9Words(digits string, dictionary []string) []string {
	if digits == "" {
		return nil
	}
	for _, ch := range digits {
		if !unicode.IsDigit(ch) || ch == '0' || ch == '1' {
			panic("Digits must be 2-9 only.")
		}
	}
	root := &t9Node{children: make(map[rune]*t9Node)}
	for _, word := range dictionary {
		if word != "" && allLetters(word) {
			insertT9(root, toLowerASCII(word))
		}
	}
	found := searchT9(root, digits, 0)
	uniq := make(map[string]struct{})
	for _, w := range found {
		uniq[w] = struct{}{}
	}
	out := make([]string, 0, len(uniq))
	for w := range uniq {
		out = append(out, w)
	}
	sort.Strings(out)
	return out
}

func allLetters(word string) bool {
	for _, ch := range word {
		if !unicode.IsLetter(ch) {
			return false
		}
	}
	return true
}

func toLowerASCII(word string) string {
	b := make([]byte, len(word))
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if ch >= 'A' && ch <= 'Z' {
			ch += 'a' - 'A'
		}
		b[i] = ch
	}
	return string(b)
}

func insertT9(root *t9Node, word string) {
	node := root
	for _, ch := range word {
		child := node.children[ch]
		if child == nil {
			child = &t9Node{children: make(map[rune]*t9Node)}
			node.children[ch] = child
		}
		node = child
	}
	node.word = word
}

func searchT9(node *t9Node, digits string, index int) []string {
	if index == len(digits) {
		if node.word == "" {
			return nil
		}
		return []string{node.word}
	}
	results := make([]string, 0)
	digit := int(digits[index] - '0')
	for _, ch := range t9Letters[digit] {
		if child := node.children[ch]; child != nil {
			results = append(results, searchT9(child, digits, index+1)...)
		}
	}
	return results
}

func RunQ1620() {
	fmt.Println(T9Words("8733", []string{"tree", "used", "trend", "apple"}))
}
