package chapter06

import (
	"fmt"
	"math"
)

func ProbAheadAfterShots(p float64, shots, needAhead int) float64 {
	if shots < 0 || needAhead < 0 {
		panic("shots and need_ahead must be non-negative")
	}
	if p < 0.0 || p > 1.0 {
		panic("p must be in [0, 1]")
	}
	offset := shots
	dp := make([]float64, 2*shots+1)
	dp[offset] = 1.0
	for range shots {
		next := make([]float64, 2*shots+1)
		for k := -shots; k <= shots; k++ {
			prob := dp[k+offset]
			if prob == 0.0 {
				continue
			}
			next[k+1+offset] += prob * p
			next[k-1+offset] += prob * (1.0 - p)
		}
		dp = next
	}
	total := 0.0
	for k := needAhead; k <= shots; k++ {
		total += dp[k+offset]
	}
	return total
}

func ProbWinRaceToN(p float64, n int) float64 {
	if n < 1 {
		panic("n must be positive")
	}
	maxShots := int(math.Min(200, math.Max(20, float64(4*n*n))))
	return ProbAheadAfterShots(p, maxShots, n)
}

func ChooseGame(p float64) string {
	if p < 0.0 || p > 1.0 {
		panic("p must be in [0, 1]")
	}
	game1 := p
	game2 := 3*p*p*(1-p) + p*p*p
	if game1 > game2 {
		return "game1"
	}
	if game2 > game1 {
		return "game2"
	}
	return "tie"
}

func RunQ602() {
	for _, p := range []float64{0.1, 0.5, 0.9} {
		fmt.Printf("p=%v: %s\n", p, ChooseGame(p))
	}
}
