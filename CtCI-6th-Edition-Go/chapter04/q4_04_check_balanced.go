package chapter04

import (
	"fmt"

	"ctci/library"
)

func absInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// IsBalanced reports whether every node's subtree heights differ by at most 1.
// Time O(n log n), space O(h).
func IsBalanced(root *library.TreeNode) bool {
	if root == nil {
		return true
	}
	if absInt(height(root.Left)-height(root.Right)) > 1 {
		return false
	}
	return IsBalanced(root.Left) && IsBalanced(root.Right)
}

// IsBalancedBook reports balance in a single O(n) walk.
func IsBalancedBook(root *library.TreeNode) bool {
	return checkHeight(root) != -1
}

func height(root *library.TreeNode) int {
	if root == nil {
		return 0
	}
	return maxInt(height(root.Left), height(root.Right)) + 1
}

func checkHeight(root *library.TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := checkHeight(root.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := checkHeight(root.Right)
	if rightHeight == -1 {
		return -1
	}
	if absInt(leftHeight-rightHeight) > 1 {
		return -1
	}
	return maxInt(leftHeight, rightHeight) + 1
}

func RunQ404() {
	balanced := Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("Minimal BST: %t / %t\n", IsBalanced(balanced), IsBalancedBook(balanced))
	unbalanced := library.NewTreeNode(2)
	unbalanced.SetLeftChild(library.NewTreeNode(1))
	unbalanced.Left.SetLeftChild(library.NewTreeNode(3))
	fmt.Printf("Skewed left: %t / %t\n", IsBalanced(unbalanced), IsBalancedBook(unbalanced))
}
