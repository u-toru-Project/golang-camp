package chapter15

import (
	"strings"
	"testing"
)

func TestSumSquares(t *testing.T) {
	if SumSquares(5) != 30 {
		t.Fatal(SumSquares(5))
	}
	if ParallelSumSquares([]int{3, 4}) != SumSquares(3)+SumSquares(4) {
		t.Fatal("parallel")
	}
	if ParallelSumSquares(nil) != 0 || SumSquares(0) != 0 {
		t.Fatal("empty")
	}
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("expected panic")
			}
		}()
		SumSquares(-1)
	}()
	note := DemoNote()
	if !strings.Contains(strings.ToLower(note), "thread") || !strings.Contains(strings.ToLower(note), "process") {
		t.Fatal(note)
	}
}

func TestCooperativeScheduler(t *testing.T) {
	log := CooperativeScheduler([]Enumerator{TaskA(), TaskB()})
	joined := strings.Join(log, " ")
	if !strings.Contains(joined, "t0") || !strings.Contains(joined, "t1") {
		t.Fatal(log)
	}
}

func TestDine(t *testing.T) {
	if Dine(5, 3) != 15 {
		t.Fatal(Dine(5, 3))
	}
}

func TestOrderedLocks(t *testing.T) {
	AcquireBothOrdersWithoutDeadlock()
	a := NewOrderedLock()
	b := NewOrderedLock()
	AcquireOrdered(a, b)
	ReleaseOrdered(a, b)
	AcquireOrdered(b, a)
	ReleaseOrdered(b, a)
	if a.Rank == b.Rank {
		t.Fatal("ranks should differ")
	}
}

func TestRunFooBar(t *testing.T) {
	if RunFooBar(3) != "foobarfoobarfoobar" {
		t.Fatal(RunFooBar(3))
	}
}

func TestRunHammer(t *testing.T) {
	if RunHammer(4, 1000) != 4000 {
		t.Fatal("hammer")
	}
	counter := &SyncCounter{}
	Hammer(counter, 7)
	if counter.Get() != 7 || counter.Increment() != 8 {
		t.Fatal("single thread")
	}
}

func TestRunFizzBuzz(t *testing.T) {
	result := RunFizzBuzz(15)
	if len(result) != 15 || result[2] != "fizz" || result[4] != "buzz" || result[14] != "fizzbuzz" {
		t.Fatal(result)
	}
}
