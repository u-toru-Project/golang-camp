package chapter02

import (
	"fmt"

	"ctci/library"
)

// DeleteDupsWithSet removes duplicate values, keeping the first occurrence.
// Time O(n), space O(n).
func DeleteDupsWithSet(node *library.LinkedListNode) {
	seen := make(map[int]struct{})
	var previous *library.LinkedListNode
	for node != nil {
		if _, ok := seen[node.Data]; ok {
			if previous != nil {
				previous.Next = node.Next
			}
		} else {
			seen[node.Data] = struct{}{}
			previous = node
		}
		node = node.Next
	}
}

// DeleteDupsWithRunner removes duplicates with a runner pointer.
// Time O(n^2), space O(1).
func DeleteDupsWithRunner(head *library.LinkedListNode) {
	current := head
	for current != nil {
		runner := current
		for runner.Next != nil {
			if runner.Next.Data == current.Data {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}
		current = current.Next
	}
}

func RunQ201() {
	head := library.CreateLinkedListFromArray([]int{0, 1, 0, 1, 0, 1, 0, 1})
	withSet := head.Clone()
	withRunner := head.Clone()
	DeleteDupsWithSet(withSet)
	DeleteDupsWithRunner(withRunner)
	fmt.Println(head.PrintForward())
	fmt.Println(withSet.PrintForward())
	fmt.Println(withRunner.PrintForward())
}
