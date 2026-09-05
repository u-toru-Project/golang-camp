package chapter04

import (
	"fmt"
	"strings"

	"ctci/library"
)

// ListOfDepths groups nodes by depth.
// Time O(n), space O(n).
func ListOfDepths(root *library.TreeNode) [][]*library.TreeNode {
	lists := make([][]*library.TreeNode, 0)
	if root == nil {
		return lists
	}
	current := []*library.TreeNode{root}
	for len(current) > 0 {
		lists = append(lists, current)
		parents := current
		current = nil
		for _, node := range parents {
			if node.Left != nil {
				current = append(current, node.Left)
			}
			if node.Right != nil {
				current = append(current, node.Right)
			}
		}
	}
	return lists
}

func RunQ403() {
	tree := Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	for _, level := range ListOfDepths(tree) {
		parts := make([]string, len(level))
		for i, node := range level {
			parts[i] = fmt.Sprintf("%d", node.Data)
		}
		fmt.Println(strings.Join(parts, ","))
	}
}
