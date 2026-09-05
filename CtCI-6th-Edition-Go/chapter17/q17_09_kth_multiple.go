package chapter17

import (
	"container/heap"
	"fmt"
	"math"
)

type longHeap []int64

func (h longHeap) Len() int           { return len(h) }
func (h longHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h longHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *longHeap) Push(x any)        { *h = append(*h, x.(int64)) }
func (h *longHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func GetKthMultiple(k int) int64 {
	ensurePositiveK(k)
	numbers := []int64{1, 3, 5, 7}
	seen := map[int64]struct{}{1: {}, 3: {}, 5: {}, 7: {}}
	if k <= 3 {
		return numbers[k]
	}
	for i := 0; i < k-3; i++ {
		next := int64(math.MaxInt64)
		for _, value := range numbers {
			considerMultiple(value*3, seen, &next)
			considerMultiple(value*5, seen, &next)
			considerMultiple(value*7, seen, &next)
		}
		numbers = append(numbers, next)
		seen[next] = struct{}{}
	}
	return numbers[len(numbers)-1]
}

func GetKthMultipleViaHeap(k int) int64 {
	ensurePositiveK(k)
	seen := make(map[int64]struct{})
	h := &longHeap{3, 5, 7}
	heap.Init(h)
	var result int64
	for range k {
		result = heap.Pop(h).(int64)
		for _, factor := range []int64{3, 5, 7} {
			candidate := result * factor
			if _, ok := seen[candidate]; !ok {
				seen[candidate] = struct{}{}
				heap.Push(h, candidate)
			}
		}
	}
	return result
}

func considerMultiple(candidate int64, seen map[int64]struct{}, next *int64) {
	if _, ok := seen[candidate]; !ok && candidate < *next {
		*next = candidate
	}
}

func ensurePositiveK(k int) {
	if k < 1 {
		panic("k must be >= 1")
	}
}

func RunQ1709() {
	for k := 1; k <= 5; k++ {
		fmt.Printf("k=%d: %d / %d\n", k, GetKthMultiple(k), GetKthMultipleViaHeap(k))
	}
}
