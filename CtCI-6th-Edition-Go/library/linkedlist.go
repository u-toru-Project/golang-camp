package library

import "fmt"

type LinkedListNode struct {
	Next *LinkedListNode
	Prev *LinkedListNode
	Last *LinkedListNode
	Data int
}

func NewLinkedListNode(data int, next, prev *LinkedListNode) *LinkedListNode {
	node := &LinkedListNode{Data: data}
	if next != nil || prev != nil {
		node.SetNext(next)
		node.SetPrevious(prev)
	}
	return node
}

func (n *LinkedListNode) SetNext(node *LinkedListNode) {
	n.Next = node
	if n == n.Last {
		n.Last = node
	}
	if node != nil && node.Prev != n {
		node.SetPrevious(n)
	}
}

func (n *LinkedListNode) SetPrevious(node *LinkedListNode) {
	n.Prev = node
	if node != nil && node.Next != n {
		node.SetNext(n)
	}
}

func (n *LinkedListNode) PrintForward() string {
	if n.Next == nil {
		return fmt.Sprintf("%d", n.Data)
	}
	return fmt.Sprintf("%d->%s", n.Data, n.Next.PrintForward())
}

func (n *LinkedListNode) Clone() *LinkedListNode {
	var next *LinkedListNode
	if n.Next != nil {
		next = n.Next.Clone()
	}
	return NewLinkedListNode(n.Data, next, nil)
}

func CreateLinkedListFromArray(values []int) *LinkedListNode {
	if len(values) == 0 {
		return nil
	}
	head := NewLinkedListNode(values[0], nil, nil)
	current := head
	for i := 1; i < len(values); i++ {
		current = NewLinkedListNode(values[i], nil, current)
	}
	return head
}

func LinkedListToArray(head *LinkedListNode) []int {
	values := make([]int, 0)
	for node, guard := head, 0; node != nil && guard < 10_000; guard++ {
		values = append(values, node.Data)
		node = node.Next
	}
	return values
}
