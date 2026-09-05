package chapter03

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Count() int { return len(s.items) }

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() T {
	n := len(s.items) - 1
	value := s.items[n]
	s.items = s.items[:n]
	return value
}

func (s *Stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}

type MyQueue[T any] struct {
	NewStack Stack[T]
	OldStack Stack[T]
}

func (q *MyQueue[T]) Count() int {
	return q.NewStack.Count() + q.OldStack.Count()
}

func (q *MyQueue[T]) IsEmpty() bool {
	return q.Count() == 0
}

func (q *MyQueue[T]) Add(value T) {
	q.NewStack.Push(value)
}

func (q *MyQueue[T]) Peek() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, fmt.Errorf("peek from empty queue")
	}
	q.ShiftStacks()
	return q.OldStack.Peek(), nil
}

func (q *MyQueue[T]) Remove() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, fmt.Errorf("remove from empty queue")
	}
	q.ShiftStacks()
	return q.OldStack.Pop(), nil
}

func (q *MyQueue[T]) ShiftStacks() {
	if q.OldStack.Count() > 0 {
		return
	}
	for q.NewStack.Count() > 0 {
		q.OldStack.Push(q.NewStack.Pop())
	}
}

func RunQ304() {
	queue := &MyQueue[int]{}
	queue.Add(4)
	queue.Add(6)
	queue.Add(101)
	peek, _ := queue.Peek()
	a, _ := queue.Remove()
	b, _ := queue.Remove()
	c, _ := queue.Remove()
	fmt.Printf("peek=%d, remove=%d, remove=%d, remove=%d\n", peek, a, b, c)
}
