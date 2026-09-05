package chapter04

import (
	"fmt"
	"math/rand/v2"

	"ctci/library"
)

type RandomBinarySearchTree struct {
	Root *library.TreeNode
}

func (t *RandomBinarySearchTree) Insert(key int) *library.TreeNode {
	created := library.NewTreeNode(key)
	if t.Root == nil {
		t.Root = created
		return created
	}
	current := t.Root
	for {
		current.Size++
		if key <= current.Data {
			if current.Left == nil {
				current.SetLeftChild(created)
				return created
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.SetRightChild(created)
				return created
			}
			current = current.Right
		}
	}
}

func (t *RandomBinarySearchTree) Delete(key int) error {
	root, removed := deleteNode(t.Root, key)
	t.Root = root
	if !removed {
		return fmt.Errorf("No such value in the tree: %d", key)
	}
	return nil
}

func (t *RandomBinarySearchTree) GetNode(key int) (*library.TreeNode, error) {
	if t.Root == nil {
		return nil, fmt.Errorf("No such value in the tree: %d", key)
	}
	found := t.Root.Find(key)
	if found == nil {
		return nil, fmt.Errorf("No such value in the tree: %d", key)
	}
	return found, nil
}

func (t *RandomBinarySearchTree) GetRandomNode() (*library.TreeNode, error) {
	if t.Root == nil {
		return nil, fmt.Errorf("tree is empty")
	}
	return GetRandomNode(t.Root), nil
}

// GetRandomNode selects a node uniformly via subtree sizes.
// Time O(h), space O(1).
func GetRandomNode(root *library.TreeNode) *library.TreeNode {
	return GetIthNode(root, rand.IntN(root.Size))
}

// GetIthNode returns the in-order node at 0-based index using subtree sizes.
// Time O(h), space O(1).
func GetIthNode(root *library.TreeNode, index int) *library.TreeNode {
	if root == nil || index < 0 || index >= root.Size {
		panic("index out of range")
	}
	current := root
	remaining := index
	for {
		leftSize := 0
		if current.Left != nil {
			leftSize = current.Left.Size
		}
		if remaining < leftSize {
			current = current.Left
		} else if remaining == leftSize {
			return current
		} else {
			remaining -= leftSize + 1
			current = current.Right
		}
	}
}

func deleteNode(node *library.TreeNode, key int) (*library.TreeNode, bool) {
	if node == nil {
		return nil, false
	}
	if key < node.Data {
		left, removed := deleteNode(node.Left, key)
		node.SetLeftChild(left)
		if removed {
			node.Size--
		}
		return node, removed
	}
	if key > node.Data {
		right, removed := deleteNode(node.Right, key)
		node.SetRightChild(right)
		if removed {
			node.Size--
		}
		return node, removed
	}
	if node.Left == nil {
		if node.Right != nil {
			node.Right.Parent = node.Parent
		}
		return node.Right, true
	}
	if node.Right == nil {
		node.Left.Parent = node.Parent
		return node.Left, true
	}
	successor := minNode(node.Right)
	node.Data = successor.Data
	right, _ := deleteNode(node.Right, successor.Data)
	node.SetRightChild(right)
	node.Size--
	return node, true
}

func minNode(node *library.TreeNode) *library.TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func RunQ411() {
	tree := &RandomBinarySearchTree{}
	for _, key := range []int{20, 9, 25, 5, 12, 11, 14} {
		tree.Insert(key)
	}
	_ = tree.Delete(12)
	random, _ := tree.GetRandomNode()
	fmt.Printf("size=%d random=%d\n", tree.Root.Size, random.Data)
}
