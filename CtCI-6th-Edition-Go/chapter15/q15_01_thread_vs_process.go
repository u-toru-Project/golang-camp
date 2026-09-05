package chapter15

import "fmt"

func SumSquares(n int) int64 {
	if n < 0 {
		panic("n must be non-negative")
	}
	var sum int64
	for i := range n {
		sum += int64(i) * int64(i)
	}
	return sum
}

func ParallelSumSquares(values []int) int64 {
	var total int64
	for _, n := range values {
		total += SumSquares(n)
	}
	return total
}

func DemoNote() string {
	return "In Go, goroutines (lightweight threads) can run CPU work in parallel on multiple cores. Processes isolate memory and crash domains; use them when you need isolation, not just throughput."
}

func RunQ1501() {
	fmt.Println(SumSquares(5))
	fmt.Println(ParallelSumSquares([]int{3, 4}))
	fmt.Println(DemoNote())
}
