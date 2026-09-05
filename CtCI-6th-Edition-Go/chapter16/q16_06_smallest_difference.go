package chapter16

import (
	"fmt"
	"math"
	"sort"
)

type IntPair struct {
	First, Second int
}

func FindSmallestDifference(first, second []int) int {
	pair := GetSmallestDifferencePair(first, second)
	if pair == nil {
		return math.MaxInt32
	}
	diff := pair.First - pair.Second
	if diff < 0 {
		return -diff
	}
	return diff
}

func GetSmallestDifferencePair(first, second []int) *IntPair {
	if len(first) == 0 || len(second) == 0 {
		return nil
	}
	a := append([]int(nil), first...)
	b := append([]int(nil), second...)
	sort.Ints(a)
	sort.Ints(b)
	index1, index2 := 0, 0
	minDifference := math.MaxInt32
	best := IntPair{First: a[0], Second: b[0]}
	for index1 < len(a) && index2 < len(b) {
		value1, value2 := a[index1], b[index2]
		difference := value1 - value2
		if difference < 0 {
			difference = -difference
		}
		if difference < minDifference {
			minDifference = difference
			best = IntPair{First: value1, Second: value2}
		}
		if value1 < value2 {
			index1++
		} else {
			index2++
		}
	}
	return &best
}

func RunQ1606() {
	first := []int{1, 3, 15, 11, 2}
	second := []int{23, 127, 235, 19, 8}
	pair := GetSmallestDifferencePair(first, second)
	fmt.Printf("Pair %d, %d\n", pair.First, pair.Second)
}
