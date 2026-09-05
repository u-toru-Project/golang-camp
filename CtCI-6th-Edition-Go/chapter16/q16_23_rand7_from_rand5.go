package chapter16

import (
	"fmt"
	"math/rand"
)

func Rand5() int {
	return rand.Intn(5)
}

func Rand7(rand5 func() int) int {
	rng := rand5
	if rng == nil {
		rng = Rand5
	}
	for {
		value := rng()*5 + rng()
		if value < 21 {
			return value % 7
		}
	}
}

func RunQ1623() {
	samples := make([]int, 10)
	for i := range samples {
		samples[i] = Rand7(nil)
	}
	fmt.Printf("Rand7 samples: %v\n", samples)
}
