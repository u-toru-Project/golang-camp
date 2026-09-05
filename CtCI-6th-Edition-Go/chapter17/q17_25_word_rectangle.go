package chapter17

import "fmt"

type wrNode struct {
	children   map[byte]*wrNode
	terminates bool
}

func FindLargestWordRectangle(words []string) []string {
	if len(words) == 0 {
		return nil
	}
	byLength := make(map[int][]string)
	maxLen := 0
	for _, word := range words {
		byLength[len(word)] = append(byLength[len(word)], word)
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}
	trie := buildWordTrie(words)
	var best []string
	bestArea := 0
	for side := maxLen; side >= 1; side-- {
		candidates := byLength[side]
		if len(candidates) == 0 {
			continue
		}
		pointers := make([]*wrNode, side)
		for i := range pointers {
			pointers[i] = trie
		}
		found := searchRectangle(candidates, nil, pointers)
		if found != nil && len(found)*side > bestArea {
			best = found
			bestArea = len(found) * side
		}
		if best != nil && bestArea == side*side {
			break
		}
	}
	return best
}

func buildWordTrie(words []string) *wrNode {
	root := &wrNode{children: make(map[byte]*wrNode)}
	for _, word := range words {
		node := root
		for i := 0; i < len(word); i++ {
			ch := word[i]
			child := node.children[ch]
			if child == nil {
				child = &wrNode{children: make(map[byte]*wrNode)}
				node.children[ch] = child
			}
			node = child
		}
		node.terminates = true
	}
	return root
}

func searchRectangle(wordsOfLen []string, used []string, triePointers []*wrNode) []string {
	var best []string
	usedSet := make(map[string]struct{}, len(used))
	for _, w := range used {
		usedSet[w] = struct{}{}
	}
	for _, word := range wordsOfLen {
		if _, ok := usedSet[word]; ok {
			continue
		}
		newPointers := advanceTrieRow(triePointers, word)
		if newPointers == nil {
			continue
		}
		candidate := append(append([]string(nil), used...), word)
		if columnWordsComplete(newPointers) && (best == nil || len(candidate) > len(best)) {
			best = candidate
		}
		deeper := searchRectangle(wordsOfLen, candidate, newPointers)
		if deeper != nil && (best == nil || len(deeper) > len(best)) {
			best = deeper
		}
	}
	return best
}

func advanceTrieRow(triePointers []*wrNode, word string) []*wrNode {
	newPointers := make([]*wrNode, len(triePointers))
	for i := 0; i < len(word); i++ {
		next := triePointers[i].children[word[i]]
		if next == nil {
			return nil
		}
		newPointers[i] = next
	}
	return newPointers
}

func columnWordsComplete(triePointers []*wrNode) bool {
	for _, pointer := range triePointers {
		if !pointer.terminates {
			return false
		}
	}
	return true
}

func RunQ1725() {
	words := []string{"area", "lead", "wall", "lady", "ball"}
	rectangle := FindLargestWordRectangle(words)
	if rectangle == nil {
		fmt.Println("(none)")
		return
	}
	for _, row := range rectangle {
		fmt.Println(row)
	}
}
