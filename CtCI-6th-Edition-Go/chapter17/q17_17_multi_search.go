package chapter17

import "fmt"

type msNode struct {
	children map[byte]*msNode
	isWord   bool
}

type msTrie struct {
	root *msNode
}

func newMsTrie() *msTrie {
	return &msTrie{root: &msNode{children: make(map[byte]*msNode)}}
}

func (t *msTrie) insert(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		child := node.children[ch]
		if child == nil {
			child = &msNode{children: make(map[byte]*msNode)}
			node.children[ch] = child
		}
		node = child
	}
	node.isWord = true
}

func (t *msTrie) checkExistence(word string) []string {
	matches := make([]string, 0)
	node := t.root
	for i := 0; i < len(word); i++ {
		child := node.children[word[i]]
		if child == nil {
			break
		}
		node = child
		if node.isWord {
			matches = append(matches, word[:i+1])
		}
	}
	return matches
}

func MultiSearch(text string, searchTerms []string) map[string][]int {
	trie := newMsTrie()
	for _, word := range searchTerms {
		trie.insert(word)
	}
	result := make(map[string][]int)
	for i := 0; i < len(text); i++ {
		for _, word := range trie.checkExistence(text[i:]) {
			result[word] = append(result[word], i)
		}
	}
	return result
}

func RunQ1717() {
	found := MultiSearch("mississippi", []string{"i", "is", "pp", "ms"})
	for word, locations := range found {
		fmt.Printf("%s: %v\n", word, locations)
	}
}
