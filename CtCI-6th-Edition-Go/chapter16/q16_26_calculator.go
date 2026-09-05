package chapter16

import (
	"fmt"
	"strconv"
	"unicode"
)

var calcPriority = map[rune]int{'+': 1, '-': 1, '*': 2, '/': 2}

func Calculate(equation string) float64 {
	terms := ParseEquation(equation)
	if len(terms) == 0 {
		panic("Empty equation.")
	}
	ops := make([]rune, 0)
	nums := make([]float64, 0)
	for _, token := range terms {
		if op, ok := token.(rune); ok {
			for len(ops) > 0 && calcPriority[op] <= calcPriority[ops[len(ops)-1]] {
				if len(nums) < 2 {
					panic("Malformed equation.")
				}
				a := nums[len(nums)-1]
				b := nums[len(nums)-2]
				nums = nums[:len(nums)-2]
				nums = append(nums, CollapseRune(a, b, ops[len(ops)-1]))
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, op)
			continue
		}
		nums = append(nums, token.(float64))
	}
	for len(ops) > 0 {
		if len(nums) < 2 {
			panic("Malformed equation.")
		}
		a := nums[len(nums)-1]
		b := nums[len(nums)-2]
		nums = nums[:len(nums)-2]
		nums = append(nums, CollapseRune(a, b, ops[len(ops)-1]))
		ops = ops[:len(ops)-1]
	}
	if len(nums) != 1 {
		panic("Malformed equation.")
	}
	return nums[0]
}

func ParseEquation(equation string) []any {
	terms := make([]any, 0)
	i := 0
	for i < len(equation) {
		ch := rune(equation[i])
		if ch == ' ' {
			i++
			continue
		}
		if _, ok := calcPriority[ch]; ok {
			terms = append(terms, ch)
			i++
			continue
		}
		if unicode.IsDigit(ch) || ch == '.' {
			start := i
			i++
			for i < len(equation) && (unicode.IsDigit(rune(equation[i])) || equation[i] == '.') {
				i++
			}
			terms = append(terms, parseNumber(equation[start:i]))
			continue
		}
		panic("Unexpected character in equation: '" + string(ch) + "'.")
	}
	return terms
}

func Collapse(a, b float64, operation string) float64 {
	if len(operation) != 1 {
		panic("Unsupported operation: '" + operation + "'.")
	}
	return CollapseRune(a, b, rune(operation[0]))
}

func CollapseRune(a, b float64, operation rune) float64 {
	switch operation {
	case '+':
		return b + a
	case '-':
		return b - a
	case '*':
		return b * a
	case '/':
		if a == 0 {
			panic("Division by zero.")
		}
		return b / a
	default:
		panic("Unsupported operation: '" + string(operation) + "'.")
	}
}

func parseNumber(token string) float64 {
	dots := 0
	if token == "" {
		panic("Invalid numeric token: '" + token + "'.")
	}
	for _, ch := range token {
		if ch == '.' {
			dots++
		} else if !unicode.IsDigit(ch) {
			panic("Invalid numeric token: '" + token + "'.")
		}
	}
	if dots > 1 {
		panic("Invalid numeric token: '" + token + "'.")
	}
	value, err := strconv.ParseFloat(token, 64)
	if err != nil {
		panic("Invalid numeric token: '" + token + "'.")
	}
	return value
}

func RunQ1626() {
	for _, sample := range []string{"2 -6 - 7 * 8 / 2 + 5", "2*3+5/6*3+15"} {
		fmt.Printf("%s => %v\n", sample, Calculate(sample))
	}
}
