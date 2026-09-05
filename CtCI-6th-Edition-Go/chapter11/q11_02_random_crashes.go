package chapter11

import (
	"fmt"
	"sync"
)

type UnsafeCounter struct {
	Value int
}

func (c *UnsafeCounter) Increment(times int) {
	for range times {
		c.Value++
	}
}

type SafeCounter struct {
	gate  sync.Mutex
	Value int
}

func (c *SafeCounter) Increment(times int) {
	for range times {
		c.gate.Lock()
		c.Value++
		c.gate.Unlock()
	}
}

func RunUnsafe(threads, perThread int) int {
	counter := &UnsafeCounter{}
	var wg sync.WaitGroup
	wg.Add(threads)
	for range threads {
		go func() {
			defer wg.Done()
			counter.Increment(perThread)
		}()
	}
	wg.Wait()
	return counter.Value
}

func RunSafe(threads, perThread int) int {
	counter := &SafeCounter{}
	var wg sync.WaitGroup
	wg.Add(threads)
	for range threads {
		go func() {
			defer wg.Done()
			counter.Increment(perThread)
		}()
	}
	wg.Wait()
	return counter.Value
}

func RunQ1102() {
	fmt.Printf("unsafe=%d expected<=4000\n", RunUnsafe(8, 500))
	fmt.Printf("safe=%d\n", RunSafe(8, 500))
}
