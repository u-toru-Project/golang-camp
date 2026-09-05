package chapter03

import (
	"fmt"
)

type MultiStackError struct {
	Message string
}

func (e *MultiStackError) Error() string { return e.Message }

type StackFullError struct{ MultiStackError }

type StackEmptyError struct{ MultiStackError }

type StackDoesNotExistError struct {
	StackNum int
}

func (e *StackDoesNotExistError) Error() string {
	return fmt.Sprintf("Stack #%d does not exist", e.StackNum)
}

type MultiStack struct {
	values         []int
	sizes          []int
	stackSize      int
	numberOfStacks int
}

func NewMultiStack(stackSize, numberOfStacks int) (*MultiStack, error) {
	if stackSize < 1 {
		return nil, fmt.Errorf("stackSize must be >= 1")
	}
	if numberOfStacks < 1 {
		return nil, fmt.Errorf("numberOfStacks must be >= 1")
	}
	return &MultiStack{
		values:         make([]int, stackSize*numberOfStacks),
		sizes:          make([]int, numberOfStacks),
		stackSize:      stackSize,
		numberOfStacks: numberOfStacks,
	}, nil
}

func (s *MultiStack) Push(value, stackNum int) error {
	if err := s.assertValidStackNum(stackNum); err != nil {
		return err
	}
	if s.IsFull(stackNum) {
		return &StackFullError{MultiStackError{Message: fmt.Sprintf("Push failed: stack #%d is full", stackNum)}}
	}
	s.sizes[stackNum]++
	s.values[s.indexOfTop(stackNum)] = value
	return nil
}

func (s *MultiStack) Pop(stackNum int) (int, error) {
	if err := s.assertValidStackNum(stackNum); err != nil {
		return 0, err
	}
	if s.IsEmpty(stackNum) {
		return 0, &StackEmptyError{MultiStackError{Message: fmt.Sprintf("Cannot pop from empty stack #%d", stackNum)}}
	}
	index := s.indexOfTop(stackNum)
	value := s.values[index]
	s.values[index] = 0
	s.sizes[stackNum]--
	return value, nil
}

func (s *MultiStack) Peek(stackNum int) (int, error) {
	if err := s.assertValidStackNum(stackNum); err != nil {
		return 0, err
	}
	if s.IsEmpty(stackNum) {
		return 0, &StackEmptyError{MultiStackError{Message: fmt.Sprintf("Cannot peek at empty stack #%d", stackNum)}}
	}
	return s.values[s.indexOfTop(stackNum)], nil
}

func (s *MultiStack) IsEmpty(stackNum int) bool {
	return s.sizes[stackNum] == 0
}

func (s *MultiStack) IsFull(stackNum int) bool {
	return s.sizes[stackNum] == s.stackSize
}

func (s *MultiStack) indexOfTop(stackNum int) int {
	return stackNum*s.stackSize + s.sizes[stackNum] - 1
}

func (s *MultiStack) assertValidStackNum(stackNum int) error {
	if stackNum < 0 || stackNum >= s.numberOfStacks {
		return &StackDoesNotExistError{StackNum: stackNum}
	}
	return nil
}

func RunQ301() {
	stacks, _ := NewMultiStack(3, 3)
	_ = stacks.Push(1, 0)
	_ = stacks.Push(2, 0)
	_ = stacks.Push(10, 1)
	p0, _ := stacks.Peek(0)
	p1, _ := stacks.Peek(1)
	fmt.Printf("stack0 peek=%d, stack1 peek=%d, stack2 empty=%t\n", p0, p1, stacks.IsEmpty(2))
	popped, _ := stacks.Pop(0)
	peek, _ := stacks.Peek(0)
	fmt.Printf("pop stack0=%d, then peek=%d\n", popped, peek)
}
