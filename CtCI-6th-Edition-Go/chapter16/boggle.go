package chapter16

import (
	"fmt"
	"sort"

	"ctci/library"
)

func Solve(matrix [][]rune, trie *library.TrieNode) []string {
	if trie == nil || len(matrix) == 0 {
		return nil
	}
	h, w := len(matrix), len(matrix[0])
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	words := make(map[string]struct{})
	for row := range h {
		for col := range w {
			searchBoggle(matrix, visited, row, col, trie, "", words)
		}
	}
	out := make([]string, 0, len(words))
	for w := range words {
		out = append(out, w)
	}
	sort.Strings(out)
	return out
}

func searchBoggle(matrix [][]rune, visited [][]bool, row, col int, node *library.TrieNode, prefix string, words map[string]struct{}) {
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || visited[row][col] {
		return
	}
	child := node.GetChild(matrix[row][col])
	if child == nil {
		return
	}
	word := prefix + string(matrix[row][col])
	if child.Terminates {
		words[word] = struct{}{}
	}
	visited[row][col] = true
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			searchBoggle(matrix, visited, row+dr, col+dc, child, word, words)
		}
	}
	visited[row][col] = false
}

func RunBoggle() {
	matrix := [][]rune{
		{'b', 'c', 'e', 'p'},
		{'e', 'o', 'r', 'o'},
		{'e', 'm', 't', 'n'},
		{'s', 'e', 'a', 'i'},
	}
	trie := library.NewTrie([]string{"bee", "become", "seem", "certain", "top"})
	fmt.Printf("Words found: %v\n", Solve(matrix, trie.Root))
}
