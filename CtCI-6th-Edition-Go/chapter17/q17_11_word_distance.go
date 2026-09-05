package chapter17

import (
	"fmt"
	"math"
)

func BuildWordIndex(words []string) map[string][]int {
	index := make(map[string][]int)
	for i, word := range words {
		index[word] = append(index[word], i)
	}
	return index
}

func ShortestWordDistance(wordA, wordB string, index map[string][]int) int {
	positionsA, okA := index[wordA]
	positionsB, okB := index[wordB]
	if !okA || !okB {
		panic("both words must appear in the corpus")
	}
	i, j := 0, 0
	best := math.MaxInt32
	for i < len(positionsA) && j < len(positionsB) {
		diff := positionsA[i] - positionsB[j]
		if diff < 0 {
			diff = -diff
		}
		if diff < best {
			best = diff
		}
		if positionsA[i] <= positionsB[j] {
			i++
		} else {
			j++
		}
	}
	return best
}

func ShortestDistancesForQueries(words []string, queries [][2]string) []int {
	index := BuildWordIndex(words)
	distances := make([]int, len(queries))
	for i, q := range queries {
		distances[i] = ShortestWordDistance(q[0], q[1], index)
	}
	return distances
}

func RunQ1711() {
	words := []string{"the", "quick", "brown", "fox", "quick"}
	queries := [][2]string{{"quick", "fox"}, {"the", "fox"}}
	fmt.Println(ShortestDistancesForQueries(words, queries))
}
