package chapter1

func ArePermutations(input1, input2 string) bool {
	if len(input1) != len(input2) {
		return false
	}

	counts := countRunes(input1)

	for _, r := range input2 {
		counts[r]--
		if counts[r] < 0 {
			return false
		}
	}
	return true
}
