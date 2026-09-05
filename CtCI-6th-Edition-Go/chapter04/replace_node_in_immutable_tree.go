package chapter04

import (
	"fmt"

	"ctci/library"
)

func copyTreeNode(node *library.TreeNode) *library.TreeNode {
	return &library.TreeNode{
		Data:   node.Data,
		Size:   node.Size,
		Left:   node.Left,
		Right:  node.Right,
		Parent: node.Parent,
	}
}

// Replace path-copies ancestors and shares unmodified subtrees.
// Time O(h), space O(h).
func Replace(node *library.TreeNode, value int) *library.TreeNode {
	if node == nil {
		return nil
	}
	newNode := copyTreeNode(node)
	newNode.Data = value
	newRoot := newNode
	for node.Parent != nil {
		newRoot = copyTreeNode(node.Parent)
		newNode.Parent = newRoot
		if newRoot.Left == node {
			newRoot.Left = newNode
		}
		if newRoot.Right == node {
			newRoot.Right = newNode
		}
		node = node.Parent
		newNode = newRoot
	}
	return newRoot
}

func RunReplaceNode() {
	tree := Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	newTree := Replace(tree.Find(6), 11)
	fmt.Printf("original 6=%d replaced=%d\n", tree.Find(6).Data, newTree.Right.Left.Data)
}
