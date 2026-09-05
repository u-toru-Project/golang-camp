package chapter17

import (
	"fmt"
	"sort"
)

func LongestCompositeWord(wordList []string) *string {
	mapping := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		mapping[word] = true
	}
	ordered := append([]string(nil), wordList...)
	sort.SliceStable(ordered, func(i, j int) bool {
		return len(ordered[i]) > len(ordered[j])
	})
	for _, originalWord := range ordered {
		if containsSubwords(originalWord, true, mapping) {
			w := originalWord
			return &w
		}
	}
	return nil
}

func containsSubwords(word string, isOriginalWord bool, mapping map[string]bool) bool {
	if known, ok := mapping[word]; ok && !isOriginalWord {
		return known
	}
	for i := 1; i < len(word); i++ {
		left := word[:i]
		right := word[i:]
		if leftKnown, ok := mapping[left]; ok && leftKnown && containsSubwords(right, false, mapping) {
			return true
		}
	}
	mapping[word] = false
	return false
}

func RunQ1715() {
	words := []string{"cat", "banana", "dog", "nana", "walk", "walker", "dogwalker"}
	if result := LongestCompositeWord(words); result != nil {
		fmt.Println(*result)
	} else {
		fmt.Println("(none)")
	}
}
