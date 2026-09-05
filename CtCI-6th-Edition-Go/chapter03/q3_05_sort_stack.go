package chapter03

import "fmt"

type SortedStack struct {
	values []int
	temp   []int
}

func (s *SortedStack) Count() int { return len(s.values) }

func (s *SortedStack) Push(item int) {
	if len(s.values) == 0 || item < s.values[len(s.values)-1] {
		s.values = append(s.values, item)
		return
	}
	for len(s.values) > 0 && item > s.values[len(s.values)-1] {
		s.temp = append(s.temp, s.values[len(s.values)-1])
		s.values = s.values[:len(s.values)-1]
	}
	s.values = append(s.values, item)
	for len(s.temp) > 0 {
		s.values = append(s.values, s.temp[len(s.temp)-1])
		s.temp = s.temp[:len(s.temp)-1]
	}
}

func (s *SortedStack) Pop() (int, error) {
	if len(s.values) == 0 {
		return 0, fmt.Errorf("pop from empty stack")
	}
	value := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return value, nil
}

func RunQ305() {
	stack := &SortedStack{}
	for _, value := range []int{3, 2, 1, 4} {
		stack.Push(value)
	}
	for stack.Count() > 0 {
		value, _ := stack.Pop()
		fmt.Println(value)
	}
}
