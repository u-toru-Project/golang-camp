package chapter04

import (
	"fmt"

	"ctci/library"
)

// Create builds a minimal-height BST from a sorted unique array.
// Time O(n), space O(n).
func Create(sortedArray ...int) *library.TreeNode {
	if len(sortedArray) == 0 {
		return nil
	}
	return createRange(sortedArray, 0, len(sortedArray)-1)
}

func createRange(sortedArray []int, left, right int) *library.TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	node := library.NewTreeNode(sortedArray[mid])
	node.SetLeftChild(createRange(sortedArray, left, mid-1))
	node.SetRightChild(createRange(sortedArray, mid+1, right))
	return node
}

func RunQ402() {
	root := Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("Empty array: %t\n", Create() == nil)
	fmt.Printf("Single node: %d\n", Create(1).Data)
	_ = root
}
