package chapter16

import (
	"fmt"
	"strings"
)

func GetWordFrequency(book string) map[string]int {
	frequency := make(map[string]int)
	for token := range strings.FieldsSeq(book) {
		word := strings.ToLower(token)
		frequency[word]++
	}
	return frequency
}

func GetFrequency(wordFrequency map[string]int, word string) int {
	return wordFrequency[strings.ToLower(word)]
}

func RunQ1602() {
	book := "The dog is the DOG"
	freq := GetWordFrequency(book)
	fmt.Printf("the: %d\n", GetFrequency(freq, "the"))
}
