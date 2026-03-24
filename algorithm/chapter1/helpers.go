package chapter1

func countRunes(input string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range input {
		counts[r]++
	}
	return counts
}
