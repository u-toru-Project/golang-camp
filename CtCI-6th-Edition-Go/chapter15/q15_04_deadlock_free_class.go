package chapter15

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var orderedLockSeq atomic.Int64

type OrderedLock struct {
	Rank int
	mu   sync.Mutex
}

func NewOrderedLock() *OrderedLock {
	return &OrderedLock{Rank: int(orderedLockSeq.Add(1))}
}

func (l *OrderedLock) Acquire() {
	l.mu.Lock()
}

func (l *OrderedLock) Release() {
	l.mu.Unlock()
}

func AcquireOrdered(a, b *OrderedLock) {
	first, second := a, b
	if a.Rank > b.Rank {
		first, second = b, a
	}
	first.Acquire()
	second.Acquire()
}

func ReleaseOrdered(a, b *OrderedLock) {
	a.Release()
	b.Release()
}

func AcquireBothOrdersWithoutDeadlock() {
	a := NewOrderedLock()
	b := NewOrderedLock()

	AcquireOrdered(a, b)
	ReleaseOrdered(a, b)
	AcquireOrdered(b, a)
	ReleaseOrdered(b, a)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		AcquireOrdered(a, b)
		ReleaseOrdered(a, b)
	}()
	go func() {
		defer wg.Done()
		AcquireOrdered(b, a)
		ReleaseOrdered(b, a)
	}()
	wg.Wait()
}

func RunQ1504() {
	AcquireBothOrdersWithoutDeadlock()
	fmt.Println("ordered locks completed")
}
