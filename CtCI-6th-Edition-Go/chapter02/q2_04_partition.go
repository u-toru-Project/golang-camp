package chapter02

import (
	"fmt"

	"ctci/library"
)

// Partition moves values < pivot to the left, preserving relative order.
// Time O(n), space O(1).
func Partition(node *library.LinkedListNode, pivot int) *library.LinkedListNode {
	var beforeStart, beforeEnd, afterStart, afterEnd *library.LinkedListNode
	for node != nil {
		next := node.Next
		node.Next = nil
		if node.Data < pivot {
			if beforeStart == nil {
				beforeStart = node
				beforeEnd = beforeStart
			} else {
				beforeEnd.Next = node
				beforeEnd = node
			}
		} else if afterStart == nil {
			afterStart = node
			afterEnd = afterStart
		} else {
			afterEnd.Next = node
			afterEnd = node
		}
		node = next
	}
	if beforeStart == nil {
		return afterStart
	}
	beforeEnd.Next = afterStart
	return beforeStart
}

// PartitionByPrepending partitions by prepending, reversing order within each side.
// Time O(n), space O(1).
func PartitionByPrepending(node *library.LinkedListNode, pivot int) *library.LinkedListNode {
	var beforeStart, afterStart *library.LinkedListNode
	for node != nil {
		next := node.Next
		if node.Data < pivot {
			node.Next = beforeStart
			beforeStart = node
		} else {
			node.Next = afterStart
			afterStart = node
		}
		node = next
	}
	if beforeStart == nil {
		return afterStart
	}
	head := beforeStart
	for beforeStart.Next != nil {
		beforeStart = beforeStart.Next
	}
	beforeStart.Next = afterStart
	return head
}

func RunQ204() {
	head := library.CreateLinkedListFromArray([]int{1, 3, 7, 5, 2, 9, 4})
	fmt.Println(head.PrintForward())
	fmt.Println(Partition(head.Clone(), 5).PrintForward())
	fmt.Println(PartitionByPrepending(head.Clone(), 5).PrintForward())
}
