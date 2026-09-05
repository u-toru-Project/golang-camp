package chapter03

import "fmt"

type plateNode struct {
	value int
	above *plateNode
	below *plateNode
}

type plateStack struct {
	capacity int
	top      *plateNode
	bottom   *plateNode
	size     int
}

func (s *plateStack) isFull() bool  { return s.size == s.capacity }
func (s *plateStack) isEmpty() bool { return s.size == 0 }

func (s *plateStack) push(value int) {
	s.size++
	node := &plateNode{value: value}
	if s.size == 1 {
		s.bottom = node
	}
	joinPlate(node, s.top)
	s.top = node
}

func (s *plateStack) pop() int {
	top := s.top
	s.top = top.below
	if s.top != nil {
		s.top.above = nil
	}
	s.size--
	if s.size == 0 {
		s.bottom = nil
	}
	return top.value
}

func (s *plateStack) removeBottom() int {
	bottom := s.bottom
	s.bottom = bottom.above
	if s.bottom != nil {
		s.bottom.below = nil
	} else {
		s.top = nil
	}
	s.size--
	return bottom.value
}

func joinPlate(above, below *plateNode) {
	if below != nil {
		below.above = above
	}
	if above != nil {
		above.below = below
	}
}

type SetOfStacks struct {
	capacity int
	stacks   []*plateStack
}

func NewSetOfStacks(capacity int) (*SetOfStacks, error) {
	if capacity < 1 {
		return nil, fmt.Errorf("capacity must be >= 1")
	}
	return &SetOfStacks{capacity: capacity}, nil
}

func (s *SetOfStacks) Push(value int) {
	last := s.lastStack()
	if last != nil && !last.isFull() {
		last.push(value)
		return
	}
	stack := &plateStack{capacity: s.capacity}
	stack.push(value)
	s.stacks = append(s.stacks, stack)
}

func (s *SetOfStacks) Pop() *int {
	last := s.lastStack()
	if last == nil {
		return nil
	}
	value := last.pop()
	if last.size == 0 {
		s.stacks = s.stacks[:len(s.stacks)-1]
	}
	return &value
}

func (s *SetOfStacks) PopAt(index int) (int, error) {
	if index < 0 || index >= len(s.stacks) {
		return 0, fmt.Errorf("stack index out of range: %d", index)
	}
	return s.leftShift(index, true), nil
}

func (s *SetOfStacks) IsEmpty() bool {
	last := s.lastStack()
	return last == nil || last.isEmpty()
}

func (s *SetOfStacks) lastStack() *plateStack {
	if len(s.stacks) == 0 {
		return nil
	}
	return s.stacks[len(s.stacks)-1]
}

func (s *SetOfStacks) leftShift(index int, removeTop bool) int {
	stack := s.stacks[index]
	var removed int
	if removeTop {
		removed = stack.pop()
	} else {
		removed = stack.removeBottom()
	}
	if stack.isEmpty() {
		s.stacks = append(s.stacks[:index], s.stacks[index+1:]...)
	} else if len(s.stacks) > index+1 {
		value := s.leftShift(index+1, false)
		stack.push(value)
	}
	return removed
}

func RunQ303() {
	stacks, _ := NewSetOfStacks(3)
	for i := range 10 {
		stacks.Push(i)
	}
	popAt, _ := stacks.PopAt(0)
	pop := stacks.Pop()
	fmt.Printf("popAt(0)=%d, pop=%d, empty=%t\n", popAt, *pop, stacks.IsEmpty())
}
