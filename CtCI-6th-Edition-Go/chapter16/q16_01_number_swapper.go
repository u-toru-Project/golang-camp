package chapter16

import "fmt"

func Swap(a, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
}

func RunQ1601() {
	a, b := 5, 10
	fmt.Printf("Before %d %d\n", a, b)
	Swap(&a, &b)
	fmt.Printf("After  %d %d\n", a, b)
}
