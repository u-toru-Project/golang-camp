package chapter02

import (
	"fmt"

	"ctci/library"
)

// KthToLast returns the kth node from the end. k=1 is the last node.
// Time O(n), space O(1).
func KthToLast(head *library.LinkedListNode, k int) *library.LinkedListNode {
	if k <= 0 {
		return nil
	}
	leading, trailing := head, head
	for range k {
		if leading == nil {
			return nil
		}
		leading = leading.Next
	}
	for leading != nil {
		trailing = trailing.Next
		leading = leading.Next
	}
	return trailing
}

type indexAndNode struct {
	node  *library.LinkedListNode
	index int
}

// KthToLastRecursive returns the kth node from the end using recursion.
// Time O(n), space O(n).
func KthToLastRecursive(head *library.LinkedListNode, k int) *library.LinkedListNode {
	return kthToLastRecursive(head, k, &indexAndNode{}).node
}

func kthToLastRecursive(head *library.LinkedListNode, k int, state *indexAndNode) *indexAndNode {
	if head == nil {
		return state
	}
	kthToLastRecursive(head.Next, k, state)
	state.index++
	if state.index == k {
		state.node = head
	}
	return state
}

func RunQ202() {
	head := library.CreateLinkedListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(head.PrintForward())
	const k = 3
	iterative := KthToLast(head, k)
	recursive := KthToLastRecursive(head, k)
	fmt.Printf("%dth to last: %d / %d\n", k, iterative.Data, recursive.Data)
}
