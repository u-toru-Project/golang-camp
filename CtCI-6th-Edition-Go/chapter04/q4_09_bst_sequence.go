package chapter04

import (
	"fmt"
	"strings"

	"ctci/library"
)

// AllSequences returns every array that could have produced the BST.
// Time O(n * 4^n / sqrt(n)), space O(n * 4^n / sqrt(n)).
func AllSequences(node *library.TreeNode) [][]int {
	result := make([][]int, 0)
	if node == nil {
		result = append(result, []int{})
		return result
	}
	prefix := []int{node.Data}
	weave(resultPtr(&result), AllSequences(node.Left), AllSequences(node.Right), prefix)
	return result
}

func resultPtr(result *[][]int) *[][]int { return result }

func weave(result *[][]int, leftSeq, rightSeq [][]int, prefix []int) {
	for _, left := range leftSeq {
		for _, right := range rightSeq {
			weaved := make([][]int, 0)
			weaveLists(append([]int(nil), left...), append([]int(nil), right...), &weaved, append([]int(nil), prefix...))
			*result = append(*result, weaved...)
		}
	}
}

func weaveLists(first, second []int, weaved *[][]int, prefix []int) {
	if len(first) == 0 || len(second) == 0 {
		sequence := append(append([]int(nil), prefix...), first...)
		sequence = append(sequence, second...)
		*weaved = append(*weaved, sequence)
		return
	}
	moveHeadToPrefixAndWeave(&first, second, weaved, &prefix)
	moveHeadToPrefixAndWeave(&second, first, weaved, &prefix)
}

func moveHeadToPrefixAndWeave(first *[]int, second []int, weaved *[][]int, prefix *[]int) {
	head := (*first)[0]
	*first = (*first)[1:]
	*prefix = append(*prefix, head)
	weaveLists(*first, second, weaved, *prefix)
	*prefix = (*prefix)[:len(*prefix)-1]
	*first = append([]int{head}, *first...)
}

func RunQ409() {
	root := Create(1, 2, 3)
	for _, sequence := range AllSequences(root) {
		parts := make([]string, len(sequence))
		for i, value := range sequence {
			parts[i] = fmt.Sprintf("%d", value)
		}
		fmt.Println(strings.Join(parts, ","))
	}
}
