package chapter16

import "fmt"

func SumSwap(first, second []int) *IntPair {
	if len(first) == 0 || len(second) == 0 {
		return nil
	}
	sum1, sum2 := 0, 0
	for _, v := range first {
		sum1 += v
	}
	for _, v := range second {
		sum2 += v
	}
	diff := sum1 - sum2
	if diff%2 != 0 {
		return nil
	}
	target := diff / 2
	secondValues := make(map[int]struct{}, len(second))
	for _, v := range second {
		secondValues[v] = struct{}{}
	}
	for _, value := range first {
		needed := value - target
		if _, ok := secondValues[needed]; ok {
			return &IntPair{First: value, Second: needed}
		}
	}
	return nil
}

func RunQ1621() {
	pair := SumSwap([]int{4, 1, 2, 1, 1, 2}, []int{3, 6, 3, 3})
	if pair == nil {
		fmt.Println("None")
		return
	}
	fmt.Printf("%d, %d\n", pair.First, pair.Second)
}
