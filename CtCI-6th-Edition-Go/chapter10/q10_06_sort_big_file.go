package chapter10

import (
	"fmt"
	"sort"
)

func ExternalMergeSort(data []int, chunkSize ...int) []int {
	if data == nil {
		panic("data is nil")
	}
	size := 4
	if len(chunkSize) > 0 {
		size = chunkSize[0]
	}
	if size < 1 {
		panic("chunk_size must be positive")
	}
	if len(data) <= 1 {
		return append([]int{}, data...)
	}
	runs := make([][]int, 0)
	for start := 0; start < len(data); start += size {
		end := min(start+size, len(data))
		chunk := append([]int{}, data[start:end]...)
		sort.Ints(chunk)
		runs = append(runs, chunk)
	}
	for len(runs) > 1 {
		merged := make([][]int, 0)
		for i := 0; i < len(runs); i += 2 {
			if i+1 < len(runs) {
				merged = append(merged, mergeSortedRuns(runs[i], runs[i+1]))
			} else {
				merged = append(merged, runs[i])
			}
		}
		runs = merged
	}
	return runs[0]
}

func mergeSortedRuns(first, second []int) []int {
	merged := make([]int, 0, len(first)+len(second))
	i, j := 0, 0
	for i < len(first) && j < len(second) {
		if first[i] <= second[j] {
			merged = append(merged, first[i])
			i++
		} else {
			merged = append(merged, second[j])
			j++
		}
	}
	merged = append(merged, first[i:]...)
	merged = append(merged, second[j:]...)
	return merged
}

func RunQ1006() {
	data := []int{10, -1, 7, 7, 2, 5, 4, 3}
	fmt.Println(ExternalMergeSort(data, 2))
}
