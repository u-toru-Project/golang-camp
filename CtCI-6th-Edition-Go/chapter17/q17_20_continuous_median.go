package chapter17

import (
	"container/heap"
	"fmt"
)

type intHeap struct {
	data []int
	less func(a, b int) bool
}

func (h intHeap) Len() int           { return len(h.data) }
func (h intHeap) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h intHeap) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *intHeap) Push(x any)        { h.data = append(h.data, x.(int)) }
func (h *intHeap) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[:n-1]
	return x
}
func (h *intHeap) Peek() int { return h.data[0] }

type MedianFinder struct {
	low  *intHeap
	high *intHeap
}

func NewMedianFinder() *MedianFinder {
	return &MedianFinder{
		low:  &intHeap{less: func(a, b int) bool { return a > b }},
		high: &intHeap{less: func(a, b int) bool { return a < b }},
	}
}

func (f *MedianFinder) AddNum(value int) {
	if f.low.Len() > 0 && value > f.low.Peek() {
		heap.Push(f.high, value)
	} else {
		heap.Push(f.low, value)
	}
	if f.low.Len() < f.high.Len() {
		moved := heap.Pop(f.high).(int)
		heap.Push(f.low, moved)
	} else if f.low.Len() > f.high.Len()+1 {
		moved := heap.Pop(f.low).(int)
		heap.Push(f.high, moved)
	}
}

func (f *MedianFinder) Median() float64 {
	if f.low.Len() == 0 && f.high.Len() == 0 {
		panic("no numbers added yet")
	}
	if f.low.Len() == f.high.Len() {
		return float64(f.low.Peek()+f.high.Peek()) / 2.0
	}
	return float64(f.low.Peek())
}

func MediansAfterEach(values []int) []float64 {
	finder := NewMedianFinder()
	medians := make([]float64, len(values))
	for i, v := range values {
		finder.AddNum(v)
		medians[i] = finder.Median()
	}
	return medians
}

func RunQ1720() {
	fmt.Println(MediansAfterEach([]int{5, 15, 1}))
}
