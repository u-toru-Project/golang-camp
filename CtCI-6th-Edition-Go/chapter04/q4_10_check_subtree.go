package chapter04

import (
	"fmt"

	"ctci/library"
)

// ContainsTree reports whether subtree appears inside root.
// Time O(n m), space O(h).
func ContainsTree(root, subtree *library.TreeNode) bool {
	if subtree == nil {
		return true
	}
	if root == nil {
		return false
	}
	if root.Data == subtree.Data && matchTree(root, subtree) {
		return true
	}
	return ContainsTree(root.Left, subtree) || ContainsTree(root.Right, subtree)
}

func matchTree(first, second *library.TreeNode) bool {
	if first == nil && second == nil {
		return true
	}
	if first == nil || second == nil {
		return false
	}
	return first.Data == second.Data && matchTree(first.Left, second.Left) && matchTree(first.Right, second.Right)
}

func RunQ410() {
	tree := Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	subtree := Create(6, 7, 8, 9, 10)
	fmt.Printf("6..10 under 1..10: %t\n", ContainsTree(tree, subtree))
	subtree = Create(7, 8, 9, 10)
	fmt.Printf("7..10 under 1..10: %t\n", ContainsTree(tree, subtree))
}
