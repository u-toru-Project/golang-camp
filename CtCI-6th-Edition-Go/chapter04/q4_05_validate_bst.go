package chapter04

import (
	"fmt"

	"ctci/library"
)

// IsValidBST reports whether the tree is a BST (left <= n < right).
// Time O(n), space O(h).
func IsValidBST(root *library.TreeNode) bool {
	return isValidBST(root, nil, nil)
}

func isValidBST(root *library.TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Data <= *min {
		return false
	}
	if max != nil && root.Data > *max {
		return false
	}
	data := root.Data
	return isValidBST(root.Left, min, &data) && isValidBST(root.Right, &data, max)
}

func RunQ405() {
	valid := Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("Minimal BST: %t\n", IsValidBST(valid))
	invalid := library.NewTreeNode(2)
	invalid.SetLeftChild(library.NewTreeNode(1))
	invalid.Left.SetLeftChild(library.NewTreeNode(3))
	fmt.Printf("3 under 1: %t\n", IsValidBST(invalid))
}
