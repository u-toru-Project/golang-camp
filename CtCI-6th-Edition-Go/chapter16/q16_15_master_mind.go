package chapter16

import "fmt"

func ScoreGuess(solution, guess string) (int, int) {
	if len(solution) != len(guess) {
		panic("Solution and guess must have equal length.")
	}
	if len(solution) == 0 {
		return 0, 0
	}
	hits := 0
	solutionOther := make(map[byte]int)
	guessOther := make(map[byte]int)
	for i := 0; i < len(solution); i++ {
		if solution[i] == guess[i] {
			hits++
			continue
		}
		solutionOther[solution[i]]++
		guessOther[guess[i]]++
	}
	pseudo := 0
	for ch, guessCount := range guessOther {
		solCount := solutionOther[ch]
		if solCount < guessCount {
			pseudo += solCount
		} else {
			pseudo += guessCount
		}
	}
	return hits, pseudo
}

func RunQ1615() {
	hits, pseudo := ScoreGuess("RGBY", "GGRR")
	fmt.Printf("Hits=%d pseudo=%d\n", hits, pseudo)
}
