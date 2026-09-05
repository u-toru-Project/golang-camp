package chapter08

import (
	"fmt"
	"strings"
)

type StackTooBigError struct {
	Message string
}

func (e *StackTooBigError) Error() string { return e.Message }

type Stack struct {
	items     []int
	StackSize int
}

func NewStack(stackSize int) *Stack {
	if stackSize < 1 {
		panic("stack_size must be positive")
	}
	return &Stack{StackSize: stackSize}
}

func (s *Stack) Items() []int { return append([]int{}, s.items...) }

func (s *Stack) Push(value int) {
	if len(s.items) == s.StackSize {
		panic(&StackTooBigError{Message: "stack already reached max size"})
	}
	s.items = append(s.items, value)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		panic("pop attempted from an empty stack")
	}
	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return value
}

func (s *Stack) Top() int {
	if len(s.items) == 0 {
		panic("top attempted from an empty stack")
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) String() string {
	rev := make([]string, len(s.items))
	for i := range s.items {
		rev[i] = fmt.Sprintf("%d", s.items[len(s.items)-1-i])
	}
	return strings.Join(rev, " ")
}

type MultiStack struct {
	stacks    []*Stack
	StackSize int
	NumStacks int
}

func NewMultiStack(stackSize int, numStacks ...int) *MultiStack {
	n := 3
	if len(numStacks) > 0 {
		n = numStacks[0]
	}
	if n < 1 {
		panic("num_stacks must be positive")
	}
	stacks := make([]*Stack, n)
	for i := range stacks {
		stacks[i] = NewStack(stackSize)
	}
	return &MultiStack{stacks: stacks, StackSize: stackSize, NumStacks: n}
}

func (m *MultiStack) GetStack(stackNum int) *Stack {
	if stackNum < 0 || stackNum >= m.NumStacks {
		panic("stack_num invalid")
	}
	return m.stacks[stackNum]
}

func (m *MultiStack) Push(stackNum, value int) { m.GetStack(stackNum).Push(value) }
func (m *MultiStack) Top(stackNum int) int     { return m.GetStack(stackNum).Top() }
func (m *MultiStack) Pop(stackNum int) int     { return m.GetStack(stackNum).Pop() }

type TowersOfHanoi struct {
	stacks    *MultiStack
	StackSize int
	Debug     bool
}

func NewTowersOfHanoi(stackSize int, debug ...bool) *TowersOfHanoi {
	if stackSize < 1 {
		panic("stack_size must be positive")
	}
	d := false
	if len(debug) > 0 {
		d = debug[0]
	}
	t := &TowersOfHanoi{stacks: NewMultiStack(stackSize), StackSize: stackSize, Debug: d}
	for value := stackSize; value >= 1; value-- {
		t.stacks.Push(0, value)
	}
	return t
}

func (t *TowersOfHanoi) Solve() { t.solve(t.StackSize, 0, 1, 2) }

func (t *TowersOfHanoi) solve(n, from, aux, to int) {
	if n <= 0 {
		return
	}
	t.solve(n-1, from, to, aux)
	t.stacks.Push(to, t.stacks.Pop(from))
	t.solve(n-1, aux, from, to)
}

func (t *TowersOfHanoi) GetStack(stackNum int) []int {
	return t.stacks.GetStack(stackNum).Items()
}

func RunQ806() {
	towers := NewTowersOfHanoi(3)
	fmt.Println("before:")
	fmt.Println(towers.stacks.GetStack(0), towers.stacks.GetStack(1), towers.stacks.GetStack(2))
	towers.Solve()
	fmt.Println("after:")
	fmt.Println(towers.GetStack(0), towers.GetStack(1), towers.GetStack(2))
}
