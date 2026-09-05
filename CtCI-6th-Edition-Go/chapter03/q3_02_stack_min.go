package chapter03

import (
	"fmt"
)

type MinStack struct {
	values []int
	mins   []int
}

func (s *MinStack) Push(value int) {
	s.values = append(s.values, value)
	if len(s.mins) == 0 || value <= s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, value)
	}
}

func (s *MinStack) Pop() (int, error) {
	if len(s.values) == 0 {
		return 0, fmt.Errorf("pop from empty stack")
	}
	value := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	if value == s.mins[len(s.mins)-1] {
		s.mins = s.mins[:len(s.mins)-1]
	}
	return value, nil
}

func (s *MinStack) Minimum() *int {
	if len(s.mins) == 0 {
		return nil
	}
	min := s.mins[len(s.mins)-1]
	return &min
}

func RunQ302() {
	stack := &MinStack{}
	for _, value := range []int{5, 6, 3, 7, 3} {
		stack.Push(value)
		fmt.Printf("push %d, min=%d\n", value, *stack.Minimum())
	}
	for stack.Minimum() != nil {
		value, _ := stack.Pop()
		min := stack.Minimum()
		if min == nil {
			fmt.Printf("pop %d, min=<nil>\n", value)
			continue
		}
		fmt.Printf("pop %d, min=%d\n", value, *min)
	}
}
