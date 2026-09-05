package chapter02

import (
	"fmt"

	"ctci/library"
)

// FindLoopStart returns the start of a cycle, or nil if there is none.
// Time O(n), space O(1).
func FindLoopStart(head *library.LinkedListNode) *library.LinkedListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}

func RunQ208() {
	nodes := make([]*library.LinkedListNode, 10)
	for i := range nodes {
		var prev *library.LinkedListNode
		if i > 0 {
			prev = nodes[i-1]
		}
		nodes[i] = library.NewLinkedListNode(i+1, nil, prev)
	}
	nodes[len(nodes)-1].Next = nodes[6]
	loop := FindLoopStart(nodes[0])
	if loop == nil {
		fmt.Println("No Cycle.")
		return
	}
	fmt.Printf("%d\n", loop.Data)
}
