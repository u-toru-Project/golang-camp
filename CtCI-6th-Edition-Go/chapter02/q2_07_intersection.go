package chapter02

import (
	"fmt"

	"ctci/library"
)

type tailAndSize struct {
	tail *library.LinkedListNode
	size int
}

// FindIntersection returns the first shared node by reference, if any.
// Time O(n + m), space O(1).
func FindIntersection(first, second *library.LinkedListNode) *library.LinkedListNode {
	firstTail := getTailAndSize(first)
	secondTail := getTailAndSize(second)
	if firstTail == nil || secondTail == nil || firstTail.tail != secondTail.tail {
		return nil
	}
	shorter, longer := first, second
	if firstTail.size < secondTail.size {
		shorter, longer = first, second
	} else {
		shorter, longer = second, first
	}
	diff := firstTail.size - secondTail.size
	if diff < 0 {
		diff = -diff
	}
	longer = getKthNode(longer, diff)
	for shorter != longer {
		shorter = shorter.Next
		longer = longer.Next
	}
	return longer
}

func getTailAndSize(list *library.LinkedListNode) *tailAndSize {
	if list == nil {
		return nil
	}
	size := 1
	current := list
	for current.Next != nil {
		size++
		current = current.Next
	}
	return &tailAndSize{tail: current, size: size}
}

func getKthNode(head *library.LinkedListNode, k int) *library.LinkedListNode {
	current := head
	for k > 0 && current != nil {
		current = current.Next
		k--
	}
	return current
}

func RunQ207() {
	first := library.CreateLinkedListFromArray([]int{-1, -2, 0, 1, 2, 3, 4, 5, 6, 7, 8})
	second := library.CreateLinkedListFromArray([]int{12, 14, 15})
	second.Next.Next = first.Next.Next.Next.Next
	fmt.Println(first.PrintForward())
	fmt.Println(second.PrintForward())
	fmt.Println(FindIntersection(first, second).PrintForward())
}
