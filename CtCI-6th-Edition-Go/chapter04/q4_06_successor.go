package chapter04

import (
	"fmt"

	"ctci/library"
)

// Successor returns the in-order successor, or nil.
// Time O(h), space O(1).
func Successor(node *library.TreeNode) *library.TreeNode {
	if node == nil {
		return nil
	}
	if node.Right != nil {
		return leftMost(node.Right)
	}
	current := node
	parent := current.Parent
	for parent != nil && parent.Right == current {
		current = parent
		parent = current.Parent
	}
	return parent
}

func leftMost(node *library.TreeNode) *library.TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func RunQ406() {
	root := Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	for _, value := range []int{6, 1, 7, 4, 10} {
		node := root.Find(value)
		successor := Successor(node)
		if successor == nil {
			fmt.Printf("Successor of %d is none\n", value)
			continue
		}
		fmt.Printf("Successor of %d is %d\n", value, successor.Data)
	}
}
