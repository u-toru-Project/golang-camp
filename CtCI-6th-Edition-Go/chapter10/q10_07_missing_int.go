package chapter10

import "fmt"

func MissingIntBitVector(nums []int) int {
	if nums == nil {
		panic("nums is nil")
	}
	n := len(nums)
	if n < 1 {
		panic("nums must have length at least 1")
	}
	seenValues := map[int]struct{}{}
	for _, value := range nums {
		if value < 0 || value > n {
			panic("values must lie in 0..n")
		}
		if _, ok := seenValues[value]; ok {
			panic("nums must be unique")
		}
		seenValues[value] = struct{}{}
	}
	seen := make([]bool, n+1)
	for _, value := range nums {
		seen[value] = true
	}
	for candidate := 0; candidate <= n; candidate++ {
		if !seen[candidate] {
			return candidate
		}
	}
	panic("no missing value")
}

func MissingIntSum(nums []int) int {
	if nums == nil {
		panic("nums is nil")
	}
	n := len(nums)
	expected := int64(n) * int64(n+1) / 2
	actual := int64(0)
	for _, value := range nums {
		actual += int64(value)
	}
	return int(expected - actual)
}

func RunQ1007() {
	nums := []int{0, 1, 3}
	fmt.Printf("%d / %d\n", MissingIntBitVector(nums), MissingIntSum(nums))
}
