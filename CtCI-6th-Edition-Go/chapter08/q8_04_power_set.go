package chapter08

import "slices"

import "fmt"

func GetSubsetsA(set []int, index ...int) [][]int {
	idx := len(set) - 1
	if len(index) > 0 {
		idx = index[0]
	}
	if idx == -1 {
		return [][]int{{}}
	}
	oldSubs := GetSubsetsA(set, idx-1)
	newSubs := make([][]int, 0)
	item := set[idx]
	for _, subset := range oldSubs {
		copySubset := append([]int{}, subset...)
		newSubs = append(newSubs, copySubset)
		entry := append(append([]int{}, subset...), item)
		newSubs = append(newSubs, entry)
	}
	return newSubs
}

func GetSubsetsB(set []int) [][]int {
	all := make([][]int, 0)
	maxN := 1 << len(set)
	for k := range maxN {
		all = append(all, ConvertIntToSet(k, set))
	}
	return all
}

func ConvertIntToSet(value int, set []int) []int {
	subset := make([]int, 0)
	index := 0
	k := value
	for k > 0 {
		if k&1 == 1 && !containsInt(subset, set[index]) {
			subset = append(subset, set[index])
		}
		index++
		k >>= 1
	}
	return subset
}

func GetSubsetsC(set []int) [][]int {
	subsets := [][]int{{}}
	var recurse func(current []int, remaining []int)
	recurse = func(current []int, remaining []int) {
		if len(remaining) == 0 {
			return
		}
		for i := range remaining {
			candidate := append(append([]int{}, current...), remaining[i])
			if !containsSubset(subsets, candidate) {
				subsets = append(subsets, candidate)
				recurse(candidate, remaining[i+1:])
			}
		}
	}
	recurse(nil, set)
	return subsets
}

func containsInt(values []int, target int) bool {
	return slices.Contains(values, target)
}

func containsSubset(subsets [][]int, candidate []int) bool {
	for _, existing := range subsets {
		if len(existing) != len(candidate) {
			continue
		}
		same := true
		for i := range existing {
			if existing[i] != candidate[i] {
				same = false
				break
			}
		}
		if same {
			return true
		}
	}
	return false
}

func RunQ804() {
	set := []int{1, 2, 3}
	fmt.Printf("A=%v\n", GetSubsetsA(set))
	fmt.Printf("B=%v\n", GetSubsetsB(set))
	fmt.Printf("C=%v\n", GetSubsetsC(set))
}
