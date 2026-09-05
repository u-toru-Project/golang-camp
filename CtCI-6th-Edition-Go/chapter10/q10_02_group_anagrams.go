package chapter10

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func GroupAnagrams(words []string) {
	groups := map[string][]string{}
	order := make([]string, 0)
	for _, word := range words {
		key := sortedKey(word)
		if _, ok := groups[key]; !ok {
			order = append(order, key)
		}
		groups[key] = append(groups[key], word)
	}
	index := 0
	for _, key := range order {
		for _, word := range groups[key] {
			words[index] = word
			index++
		}
	}
}

func SortByAnagram(words []string) {
	sort.SliceStable(words, func(i, j int) bool {
		return sortedKey(words[i]) < sortedKey(words[j])
	})
}

func sortedKey(word string) string {
	runes := []rune(word)
	slices.Sort(runes)
	return string(runes)
}

func RunQ1002() {
	words := []string{"apple", "banana", "carrot", "ele", "duck", "papel", "tarroc", "cudk", "eel", "lee"}
	fmt.Println(strings.Join(words, " "))
	SortByAnagram(words)
	fmt.Println(strings.Join(words, " "))
	GroupAnagrams(words)
	fmt.Println(strings.Join(words, " "))
}
