package chapter15

import (
	"fmt"
	"sync"
)

type SyncCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SyncCounter) Increment() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	return c.value
}

func (c *SyncCounter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func Hammer(counter *SyncCounter, times int) {
	for range times {
		counter.Increment()
	}
}

func RunHammer(threads, times int) int {
	counter := &SyncCounter{}
	var wg sync.WaitGroup
	wg.Add(threads)
	for range threads {
		go func() {
			defer wg.Done()
			Hammer(counter, times)
		}()
	}
	wg.Wait()
	return counter.Get()
}

func RunQ1506() {
	fmt.Println(RunHammer(4, 1000))
}
