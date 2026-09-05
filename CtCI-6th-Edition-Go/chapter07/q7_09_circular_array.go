package chapter07

import "fmt"

type CircularArray[T any] struct {
	data  []T
	start int
	size  int
}

func NewCircularArray[T any](capacity int) *CircularArray[T] {
	if capacity < 1 {
		panic("capacity must be positive")
	}
	return &CircularArray[T]{data: make([]T, capacity)}
}

func (a *CircularArray[T]) Count() int { return a.size }

func (a *CircularArray[T]) Append(item T) {
	if a.size < len(a.data) {
		index := (a.start + a.size) % len(a.data)
		a.data[index] = item
		a.size++
		return
	}
	a.data[a.start] = item
	a.start = (a.start + 1) % len(a.data)
}

func (a *CircularArray[T]) ToList() []T {
	result := make([]T, a.size)
	for i := 0; i < a.size; i++ {
		result[i] = a.data[(a.start+i)%len(a.data)]
	}
	return result
}

func (a *CircularArray[T]) Get(index int) T {
	if index < 0 || index >= a.size {
		panic("index out of range")
	}
	return a.data[(a.start+index)%len(a.data)]
}

func RunQ709() {
	buf := NewCircularArray[int](3)
	for _, x := range []int{1, 2, 3, 4} {
		buf.Append(x)
	}
	fmt.Println(buf.ToList())
}
