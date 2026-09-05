package chapter10

import "fmt"

func SearchSparse(strings []string, value string) int {
	if strings == nil || value == "" {
		return -1
	}
	return searchSparse(strings, value, 0, len(strings)-1)
}

func SearchSparseIterative(strings []string, value string) int {
	if strings == nil || value == "" {
		return -1
	}
	first := 0
	last := len(strings) - 1
	for first <= last {
		mid := findNonEmptyMid(strings, first, last, first+(last-first)/2)
		if mid < 0 {
			return -1
		}
		if strings[mid] == value {
			return mid
		}
		if strings[mid] < value {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}
	return -1
}

func searchSparse(strings []string, value string, first, last int) int {
	if first > last {
		return -1
	}
	mid := findNonEmptyMid(strings, first, last, first+(last-first)/2)
	if mid < 0 {
		return -1
	}
	if strings[mid] == value {
		return mid
	}
	if strings[mid] < value {
		return searchSparse(strings, value, mid+1, last)
	}
	return searchSparse(strings, value, first, mid-1)
}

func findNonEmptyMid(strings []string, first, last, mid int) int {
	if strings[mid] != "" {
		return mid
	}
	left := mid - 1
	right := mid + 1
	for {
		if left < first && right > last {
			return -1
		}
		if right <= last && strings[right] != "" {
			return right
		}
		if left >= first && strings[left] != "" {
			return left
		}
		right++
		left--
	}
}

func RunQ1005() {
	strings := []string{"apple", "", "", "banana", "", "", "", "carrot", "duck", "", "", "eel", "", "flower"}
	for _, value := range []string{"apple", "banana", "carrot", "duck", "eel", "flower", "missing"} {
		fmt.Printf("%s: %d / %d\n", value, SearchSparse(strings, value), SearchSparseIterative(strings, value))
	}
}
