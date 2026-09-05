package chapter08

import (
	"fmt"
	"sort"
)

type Box struct {
	Height int
	Width  int
	Depth  int
}

func NewBox(height, width, depth int) Box {
	if height <= 0 || width <= 0 || depth <= 0 {
		panic("dimensions must be a positive int")
	}
	return Box{Height: height, Width: width, Depth: depth}
}

func (b Box) compare(other Box) int {
	if b.Height != other.Height {
		return b.Height - other.Height
	}
	if b.Width != other.Width {
		return b.Width - other.Width
	}
	return b.Depth - other.Depth
}

func TallestStack(boxes []Box) int {
	if len(boxes) == 0 {
		return 0
	}
	sorted := append([]Box{}, boxes...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].compare(sorted[j]) > 0 })
	largestHeight := 0
	var tallestForBottom func(current []Box, currentBoxIndex int) int
	tallestForBottom = func(current []Box, currentBoxIndex int) int {
		if currentBoxIndex == len(sorted) {
			sum := 0
			for _, box := range current {
				sum += box.Height
			}
			return sum
		}
		candidate := sorted[currentBoxIndex]
		top := current[len(current)-1]
		if top.Height > candidate.Height && top.Width > candidate.Width && top.Depth > candidate.Depth {
			next := append(append([]Box{}, current...), candidate)
			return tallestForBottom(next, currentBoxIndex+1)
		}
		return tallestForBottom(current, currentBoxIndex+1)
	}
	for i := range sorted {
		h := tallestForBottom([]Box{sorted[i]}, i+1)
		if h > largestHeight {
			largestHeight = h
		}
	}
	return largestHeight
}

func RunQ813() {
	boxes := []Box{NewBox(3, 2, 1), NewBox(6, 5, 4)}
	fmt.Printf("tallest=%d\n", TallestStack(boxes))
}
