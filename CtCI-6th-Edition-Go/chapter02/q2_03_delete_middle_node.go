package chapter02

import (
	"fmt"

	"ctci/library"
)

// DeleteNode deletes a middle node by copying the next node over.
// Time O(1), space O(1). Cannot delete the last node.
func DeleteNode(node *library.LinkedListNode) bool {
	if node == nil || node.Next == nil {
		return false
	}
	node.Data = node.Next.Data
	node.Next = node.Next.Next
	return true
}

func RunQ203() {
	head := library.CreateLinkedListFromArray([]int{1, 2, 3, 4, 5})
	fmt.Println(head.PrintForward())
	fmt.Printf("deleted? %t\n", DeleteNode(head.Next.Next))
	fmt.Println(head.PrintForward())
}
