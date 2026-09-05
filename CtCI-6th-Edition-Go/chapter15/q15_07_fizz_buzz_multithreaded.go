package chapter15

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type FizzBuzz struct {
	n       int
	mu      sync.Mutex
	cond    *sync.Cond
	current int
}

func NewFizzBuzz(n int) *FizzBuzz {
	if n < 1 {
		panic("n must be positive")
	}
	fb := &FizzBuzz{n: n, current: 1}
	fb.cond = sync.NewCond(&fb.mu)
	return fb
}

func (f *FizzBuzz) Fizz(output *[]string) {
	f.runWorker(output, "fizz")
}

func (f *FizzBuzz) Buzz(output *[]string) {
	f.runWorker(output, "buzz")
}

func (f *FizzBuzz) Fizzbuzz(output *[]string) {
	f.runWorker(output, "fizzbuzz")
}

func (f *FizzBuzz) Number(output *[]string) {
	for {
		f.mu.Lock()
		for f.current <= f.n && !((f.current%3 != 0) && (f.current%5 != 0)) {
			f.cond.Wait()
		}
		if f.current > f.n {
			f.cond.Broadcast()
			f.mu.Unlock()
			return
		}
		*output = append(*output, strconv.Itoa(f.current))
		f.current++
		f.cond.Broadcast()
		f.mu.Unlock()
	}
}

func (f *FizzBuzz) runWorker(output *[]string, matches string) {
	for {
		f.mu.Lock()
		for f.current <= f.n && LabelFor(f.current) != matches {
			f.cond.Wait()
		}
		if f.current > f.n {
			f.cond.Broadcast()
			f.mu.Unlock()
			return
		}
		*output = append(*output, matches)
		f.current++
		f.cond.Broadcast()
		f.mu.Unlock()
	}
}

func LabelFor(value int) string {
	if value%15 == 0 {
		return "fizzbuzz"
	}
	if value%3 == 0 {
		return "fizz"
	}
	if value%5 == 0 {
		return "buzz"
	}
	return strconv.Itoa(value)
}

func RunFizzBuzz(n int) []string {
	fb := NewFizzBuzz(n)
	output := make([]string, 0, n)
	var wg sync.WaitGroup
	wg.Add(4)
	go func() { defer wg.Done(); fb.Fizz(&output) }()
	go func() { defer wg.Done(); fb.Buzz(&output) }()
	go func() { defer wg.Done(); fb.Fizzbuzz(&output) }()
	go func() { defer wg.Done(); fb.Number(&output) }()
	wg.Wait()
	return output
}

func RunQ1507() {
	fmt.Println(strings.Join(RunFizzBuzz(15), " "))
}
