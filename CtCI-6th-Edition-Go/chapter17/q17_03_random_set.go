package chapter17

import (
	"fmt"
	"math/rand"
)

type RandomSet struct {
	values []int
	index  map[int]int
}

func NewRandomSet() *RandomSet {
	return &RandomSet{index: make(map[int]int)}
}

func (s *RandomSet) Count() int {
	return len(s.values)
}

func (s *RandomSet) Insert(value int) bool {
	if _, ok := s.index[value]; ok {
		return false
	}
	s.index[value] = len(s.values)
	s.values = append(s.values, value)
	return true
}

func (s *RandomSet) Remove(value int) bool {
	slot, ok := s.index[value]
	if !ok {
		return false
	}
	lastValue := s.values[len(s.values)-1]
	s.values[slot] = lastValue
	s.index[lastValue] = slot
	s.values = s.values[:len(s.values)-1]
	delete(s.index, value)
	return true
}

func (s *RandomSet) GetRandom() int {
	if len(s.values) == 0 {
		panic("set is empty")
	}
	return s.values[rand.Intn(len(s.values))]
}

func (s *RandomSet) Contains(value int) bool {
	_, ok := s.index[value]
	return ok
}

func RunQ1703() {
	set := NewRandomSet()
	for _, value := range []int{1, 2, 3, 2} {
		fmt.Printf("insert %d: %t\n", value, set.Insert(value))
	}
	fmt.Printf("remove 2: %t\n", set.Remove(2))
	fmt.Printf("count=%d random=%d\n", set.Count(), set.GetRandom())
}
