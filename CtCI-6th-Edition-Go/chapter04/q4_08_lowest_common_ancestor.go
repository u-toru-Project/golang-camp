package chapter04

import (
	"fmt"

	"ctci/library"
)

// LowestCommonAncestor finds the LCA in a binary tree that is not necessarily a BST.
// Time O(n) on a balanced tree, space O(h).
func LowestCommonAncestor(root, p, q *library.TreeNode) *library.TreeNode {
	if !covers(root, p) || !covers(root, q) {
		return nil
	}
	return lowestCommonAncestorHelper(root, p, q)
}

func lowestCommonAncestorHelper(root, p, q *library.TreeNode) *library.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	pIsOnLeft := covers(root.Left, p)
	qIsOnLeft := covers(root.Left, q)
	if pIsOnLeft != qIsOnLeft {
		return root
	}
	if pIsOnLeft {
		return lowestCommonAncestorHelper(root.Left, p, q)
	}
	return lowestCommonAncestorHelper(root.Right, p, q)
}

func covers(root, node *library.TreeNode) bool {
	if root == nil || node == nil {
		return false
	}
	if root == node {
		return true
	}
	return covers(root.Left, node) || covers(root.Right, node)
}

func RunQ408() {
	root := Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	for _, pair := range [][2]int{{1, 3}, {4, 7}, {7, 9}, {2, 4}} {
		lca := LowestCommonAncestor(root, root.Find(pair[0]), root.Find(pair[1]))
		fmt.Printf("LCA of %d,%d is %d\n", pair[0], pair[1], lca.Data)
	}
}
