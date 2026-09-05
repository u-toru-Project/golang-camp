package chapter15

import (
	"fmt"
	"strings"
	"sync"
)

type OrderedPrinter struct {
	mu   sync.Mutex
	cond *sync.Cond
	turn int
}

func NewOrderedPrinter() *OrderedPrinter {
	p := &OrderedPrinter{}
	p.cond = sync.NewCond(&p.mu)
	return p
}

func (p *OrderedPrinter) PrintFoo(foo string, output *strings.Builder) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for p.turn != 0 {
		p.cond.Wait()
	}
	output.WriteString(foo)
	p.turn = 1
	p.cond.Broadcast()
}

func (p *OrderedPrinter) PrintBar(bar string, output *strings.Builder) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for p.turn != 1 {
		p.cond.Wait()
	}
	output.WriteString(bar)
	p.turn = 0
	p.cond.Broadcast()
}

func RunFooBar(n int) string {
	var output strings.Builder
	printer := NewOrderedPrinter()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for range n {
			printer.PrintFoo("foo", &output)
		}
	}()
	go func() {
		defer wg.Done()
		for range n {
			printer.PrintBar("bar", &output)
		}
	}()
	wg.Wait()
	return output.String()
}

func RunQ1505() {
	fmt.Println(RunFooBar(3))
}
