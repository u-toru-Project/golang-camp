package chapter11

import "fmt"

type Pen struct {
	remaining    float64
	writtenChars int
}

func NewPen(inkMl float64) *Pen {
	if inkMl <= 0 {
		panic("ink must be positive")
	}
	return &Pen{remaining: inkMl}
}

func (p *Pen) RemainingInk() float64 {
	return p.remaining
}

func (p *Pen) WrittenChars() int {
	return p.writtenChars
}

func (p *Pen) Write(text string) string {
	if text == "" {
		return ""
	}
	cost := float64(len(text)) * 0.001
	if cost > p.remaining {
		affordable := int(p.remaining / 0.001)
		text = text[:affordable]
		cost = float64(len(text)) * 0.001
	}
	p.remaining -= cost
	p.writtenChars += len(text)
	return text
}

func (p *Pen) IsEmpty() bool {
	return p.remaining <= 0
}

func RunQ1105() {
	pen := NewPen(0.002)
	fmt.Printf("wrote '%s' empty=%t\n", pen.Write("hello"), pen.IsEmpty())
}
