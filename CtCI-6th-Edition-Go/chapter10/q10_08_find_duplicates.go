package chapter10

import "fmt"

const MaxValue = 32000

type BitSet struct {
	bits []int
}

func NewBitSet(size int) *BitSet {
	return &BitSet{bits: make([]int, (size>>5)+1)}
}

func (b *BitSet) Get(position int) bool {
	wordNumber := position >> 5
	bitNumber := position & 0x1F
	return b.bits[wordNumber]&(1<<bitNumber) != 0
}

func (b *BitSet) Set(position int) {
	wordNumber := position >> 5
	bitNumber := position & 0x1F
	b.bits[wordNumber] |= 1 << bitNumber
}

func FindDuplicates(values []int) []int {
	seen := NewBitSet(MaxValue)
	duplicates := make([]int, 0)
	for _, value := range values {
		index := value - 1
		if seen.Get(index) {
			duplicates = append(duplicates, value)
		} else {
			seen.Set(index)
		}
	}
	return duplicates
}

func RunQ1008() {
	values := []int{1, 2, 3, 2, 4, 3, 3}
	fmt.Printf("Duplicates: %v\n", FindDuplicates(values))
}
