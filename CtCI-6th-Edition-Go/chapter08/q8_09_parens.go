package chapter08

import "fmt"

func NextPermutation(arr []rune) bool {
	i := len(arr) - 1
	for i > 0 && arr[i-1] >= arr[i] {
		i--
	}
	if i <= 0 {
		return false
	}
	j := len(arr) - 1
	for arr[j] <= arr[i-1] {
		j--
	}
	arr[i-1], arr[j] = arr[j], arr[i-1]
	for l, r := i, len(arr)-1; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}
	return true
}

func IsMatchedParentheses(sequence []rune) bool {
	depth := 0
	for _, character := range sequence {
		if character == '(' {
			depth++
		} else if character == ')' {
			if depth < 1 {
				return false
			}
			depth--
		} else {
			return false
		}
	}
	return depth == 0
}

func GenerateParenthesesPermutationsBruteForce(numberOfPairs int) []string {
	if numberOfPairs < 0 {
		panic("number_of_pairs must be non-negative")
	}
	if numberOfPairs == 0 {
		return []string{""}
	}
	starting := make([]rune, 0, numberOfPairs*2)
	for range numberOfPairs {
		starting = append(starting, '(')
	}
	for range numberOfPairs {
		starting = append(starting, ')')
	}
	possibilities := []string{string(starting)}
	for NextPermutation(starting) {
		if IsMatchedParentheses(starting) {
			possibilities = append(possibilities, string(starting))
		}
	}
	return possibilities
}

func GenerateParenthesesPermutationsRecursive1(n int) []string {
	if n < 0 {
		panic("n must be non-negative")
	}
	result := make([]string, 0)
	var helper func(openRemaining, closedRemaining int, current string)
	helper = func(openRemaining, closedRemaining int, current string) {
		if len(current) == n*2 {
			result = append(result, current)
			return
		}
		if openRemaining > 0 {
			helper(openRemaining-1, closedRemaining, current+"(")
		}
		if closedRemaining > openRemaining {
			helper(openRemaining, closedRemaining-1, current+")")
		}
	}
	helper(n, n, "")
	return result
}

func AddParen(results *[]string, leftRem, rightRem int, buffer []rune, index int) {
	if leftRem < 0 || rightRem < leftRem {
		return
	}
	if leftRem == 0 && rightRem == 0 {
		*results = append(*results, string(buffer))
		return
	}
	buffer[index] = '('
	AddParen(results, leftRem-1, rightRem, buffer, index+1)
	buffer[index] = ')'
	AddParen(results, leftRem, rightRem-1, buffer, index+1)
}

func GenerateParenthesesPermutationsRecursive2(n int) []string {
	if n < 0 {
		panic("n must be non-negative")
	}
	results := make([]string, 0)
	buffer := make([]rune, n*2)
	for i := range buffer {
		buffer[i] = '*'
	}
	AddParen(&results, n, n, buffer, 0)
	return results
}

func RunQ809() {
	for _, n := range []int{0, 1, 2, 3} {
		fmt.Printf("%d: %v\n", n, GenerateParenthesesPermutationsRecursive1(n))
	}
}
