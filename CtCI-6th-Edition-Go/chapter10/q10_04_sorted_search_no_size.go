package chapter10

import "fmt"

type Listy struct {
	data []int
}

func NewListy(data ...int) *Listy {
	return &Listy{data: append([]int{}, data...)}
}

func (l *Listy) ElementAt(index int) int {
	if index < 0 || index >= len(l.data) {
		return -1
	}
	return l.data[index]
}

func SearchListy(listy *Listy, value int) int {
	if listy == nil {
		panic("listy is nil")
	}
	index := 1
	for listy.ElementAt(index) != -1 && listy.ElementAt(index) < value {
		index *= 2
	}
	return binarySearchListy(listy, value, index/2, index)
}

func binarySearchListy(listy *Listy, value, low, high int) int {
	for low <= high {
		middle := low + (high-low)/2
		at := listy.ElementAt(middle)
		if at == -1 || value < at {
			high = middle - 1
		} else if value > at {
			low = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

func RunQ1004() {
	listy := NewListy(1, 2, 3, 4, 5, 6, 7, 8, 9)
	for _, value := range []int{0, 1, 5, 9, 10} {
		fmt.Printf("%d: %d\n", value, SearchListy(listy, value))
	}
}
