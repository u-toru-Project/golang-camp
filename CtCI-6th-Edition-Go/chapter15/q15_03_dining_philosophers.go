package chapter15

import (
	"fmt"
	"sync"
)

type Chopstick struct {
	mu sync.Mutex
}

func (c *Chopstick) Acquire() {
	c.mu.Lock()
}

func (c *Chopstick) Release() {
	c.mu.Unlock()
}

func Dine(philosophers, meals int) int {
	sticks := make([]*Chopstick, philosophers)
	for i := range sticks {
		sticks[i] = &Chopstick{}
	}
	eaten := 0
	var countMu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(philosophers)
	for pid := range philosophers {
		go func(pid int) {
			defer wg.Done()
			left := sticks[pid]
			right := sticks[(pid+1)%philosophers]
			first, second := left, right
			if pid == philosophers-1 {
				first, second = right, left
			}
			for range meals {
				first.Acquire()
				second.Acquire()
				countMu.Lock()
				eaten++
				countMu.Unlock()
				second.Release()
				first.Release()
			}
		}(pid)
	}
	wg.Wait()
	return eaten
}

func RunQ1503() {
	fmt.Printf("eaten=%d\n", Dine(5, 3))
}
