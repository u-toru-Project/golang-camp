package chapter04

import (
	"fmt"

	"ctci/library"
)

// CountPathsWithSum counts downward paths summing to target.
// Time O(n log n) balanced / O(n^2) skewed, space O(h).
func CountPathsWithSum(root *library.TreeNode, target int) int {
	if root == nil {
		return 0
	}
	return pathsFrom(root, target) + CountPathsWithSum(root.Left, target) + CountPathsWithSum(root.Right, target)
}

// CountPathsWithSumOptimized uses a running-sum hash map of prefix frequencies.
// Time O(n), space O(n).
func CountPathsWithSumOptimized(root *library.TreeNode, target int) int {
	return countPathsWithSumOptimized(root, target, 0, map[int]int{})
}

func CreateSampleTree() *library.TreeNode {
	root := library.NewTreeNode(10)
	five := library.NewTreeNode(5)
	minusThree := library.NewTreeNode(-3)
	three := library.NewTreeNode(3)
	two := library.NewTreeNode(2)
	eleven := library.NewTreeNode(11)
	threeLeaf := library.NewTreeNode(3)
	minusTwo := library.NewTreeNode(-2)
	one := library.NewTreeNode(1)
	eight := library.NewTreeNode(8)
	minusEight := library.NewTreeNode(-8)
	root.SetLeftChild(five)
	root.SetRightChild(minusThree)
	five.SetLeftChild(three)
	five.SetRightChild(two)
	three.SetLeftChild(threeLeaf)
	three.SetRightChild(minusTwo)
	two.SetLeftChild(one)
	minusThree.SetLeftChild(eleven)
	eleven.SetLeftChild(eight)
	eight.SetLeftChild(minusEight)
	return root
}

func pathsFrom(node *library.TreeNode, remaining int) int {
	if node == nil {
		return 0
	}
	remaining -= node.Data
	counter := 0
	if remaining == 0 {
		counter = 1
	}
	return counter + pathsFrom(node.Left, remaining) + pathsFrom(node.Right, remaining)
}

func countPathsWithSumOptimized(node *library.TreeNode, target, running int, prefixCounts map[int]int) int {
	if node == nil {
		return 0
	}
	running += node.Data
	total := prefixCounts[running-target]
	if running == target {
		total++
	}
	incrementPrefix(prefixCounts, running, 1)
	total += countPathsWithSumOptimized(node.Left, target, running, prefixCounts)
	total += countPathsWithSumOptimized(node.Right, target, running, prefixCounts)
	incrementPrefix(prefixCounts, running, -1)
	return total
}

func incrementPrefix(counts map[int]int, key, delta int) {
	next := counts[key] + delta
	if next == 0 {
		delete(counts, key)
	} else {
		counts[key] = next
	}
}

func RunQ412() {
	root := CreateSampleTree()
	fmt.Printf("target 8: %d / %d\n", CountPathsWithSum(root, 8), CountPathsWithSumOptimized(root, 8))
	fmt.Printf("target 6: %d / %d\n", CountPathsWithSum(root, 6), CountPathsWithSumOptimized(root, 6))
}
