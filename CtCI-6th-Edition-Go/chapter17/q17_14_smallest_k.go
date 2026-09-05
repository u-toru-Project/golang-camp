package chapter17

import (
	"container/heap"
	"fmt"
	"sort"
)

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func SmallestKHeap(arr []int, k int) []int {
	if k < 0 {
		panic("k must be >= 0")
	}
	if k == 0 {
		return []int{}
	}
	if k > len(arr) {
		panic("k cannot exceed array length.")
	}
	h := &maxHeap{}
	for i := range k {
		heap.Push(h, arr[i])
	}
	for i := k; i < len(arr); i++ {
		if arr[i] < (*h)[0] {
			heap.Pop(h)
			heap.Push(h, arr[i])
		}
	}
	result := make([]int, k)
	for i := range k {
		result[i] = heap.Pop(h).(int)
	}
	sort.Ints(result)
	return result
}

func SmallestKQuickselect(arr []int, k int) []int {
	if k < 0 {
		panic("k must be >= 0")
	}
	if k == 0 {
		return []int{}
	}
	if k > len(arr) {
		panic("k cannot exceed array length.")
	}
	data := append([]int(nil), arr...)
	quickselect(data, 0, len(data)-1, k-1)
	result := append([]int(nil), data[:k]...)
	sort.Ints(result)
	return result
}

func quickselect(data []int, left, right, kIndex int) {
	if left >= right {
		return
	}
	pivot := data[right]
	store := left
	for i := left; i < right; i++ {
		if data[i] <= pivot {
			data[store], data[i] = data[i], data[store]
			store++
		}
	}
	data[store], data[right] = data[right], data[store]
	if store == kIndex {
		return
	}
	if kIndex < store {
		quickselect(data, left, store-1, kIndex)
	} else {
		quickselect(data, store+1, right, kIndex)
	}
}

func RunQ1714() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("heap: %v\n", SmallestKHeap(arr, 4))
	fmt.Printf("quickselect: %v\n", SmallestKQuickselect(arr, 4))
}
