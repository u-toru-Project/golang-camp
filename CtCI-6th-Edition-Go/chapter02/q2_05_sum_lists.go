package chapter02

import (
	"fmt"

	"ctci/library"
)

// AddListsReverse adds two numbers stored least-significant digit first.
// Time O(n), space O(n).
func AddListsReverse(first, second *library.LinkedListNode, carry ...int) *library.LinkedListNode {
	c := 0
	if len(carry) > 0 {
		c = carry[0]
	}
	if first == nil && second == nil && c == 0 {
		return nil
	}
	value := c
	if first != nil {
		value += first.Data
	}
	if second != nil {
		value += second.Data
	}
	result := library.NewLinkedListNode(value%10, nil, nil)
	var firstNext, secondNext *library.LinkedListNode
	if first != nil {
		firstNext = first.Next
	}
	if second != nil {
		secondNext = second.Next
	}
	result.SetNext(AddListsReverse(firstNext, secondNext, value/10))
	return result
}

type partialSum struct {
	node  *library.LinkedListNode
	carry int
}

// AddListsForward adds two numbers stored most-significant digit first.
// Time O(n), space O(n).
func AddListsForward(first, second *library.LinkedListNode) *library.LinkedListNode {
	length1, length2 := listLength(first), listLength(second)
	if length1 < length2 {
		first = padLeft(first, length2-length1)
	} else if length2 < length1 {
		second = padLeft(second, length1-length2)
	}
	sum := addForward(first, second)
	if sum.carry == 0 {
		return sum.node
	}
	return insertBefore(sum.node, sum.carry)
}

func listLength(node *library.LinkedListNode) int {
	length := 0
	for node != nil {
		length++
		node = node.Next
	}
	return length
}

func padLeft(node *library.LinkedListNode, padding int) *library.LinkedListNode {
	head := node
	for range padding {
		head = insertBefore(head, 0)
	}
	return head
}

func insertBefore(list *library.LinkedListNode, data int) *library.LinkedListNode {
	node := library.NewLinkedListNode(data, nil, nil)
	if list != nil {
		list.Prev = node
		node.Next = list
	}
	return node
}

func addForward(first, second *library.LinkedListNode) partialSum {
	if first == nil && second == nil {
		return partialSum{}
	}
	var firstNext, secondNext *library.LinkedListNode
	if first != nil {
		firstNext = first.Next
	}
	if second != nil {
		secondNext = second.Next
	}
	sum := addForward(firstNext, secondNext)
	value := sum.carry
	if first != nil {
		value += first.Data
	}
	if second != nil {
		value += second.Data
	}
	sum.node = insertBefore(sum.node, value%10)
	sum.carry = value / 10
	return sum
}

func RunQ205() {
	reverseA := library.CreateLinkedListFromArray([]int{9, 9, 9})
	reverseB := library.CreateLinkedListFromArray([]int{1, 0, 0})
	reverseSum := AddListsReverse(reverseA, reverseB)
	fmt.Printf("%s + %s = %s\n", reverseA.PrintForward(), reverseB.PrintForward(), reverseSum.PrintForward())

	forwardA := library.CreateLinkedListFromArray([]int{3, 1})
	forwardB := library.CreateLinkedListFromArray([]int{5, 9, 1})
	forwardSum := AddListsForward(forwardA, forwardB)
	fmt.Printf("%s + %s = %s\n", forwardA.PrintForward(), forwardB.PrintForward(), forwardSum.PrintForward())
}
