package chapter08

import (
	"fmt"
	"regexp"
)

var expressionPattern = regexp.MustCompile(`^[01]([&|^][01])*$`)

func StringToBool(value string) bool {
	if value != "0" && value != "1" {
		panic("invalid boolean literal: '" + value + "'")
	}
	return value == "1"
}

func CountWays(exp string, result bool, memo map[string]int) int {
	if len(exp) == 0 {
		return 0
	}
	if len(exp) == 1 {
		if StringToBool(exp) == result {
			return 1
		}
		return 0
	}
	key := exp + fmt.Sprintf("%t", result)
	if cached, ok := memo[key]; ok {
		return cached
	}
	ways := 0
	for i := 1; i < len(exp); i += 2 {
		left := exp[:i]
		right := exp[i+1:]
		op := exp[i]
		leftTrue := CountWays(left, true, memo)
		leftFalse := CountWays(left, false, memo)
		rightTrue := CountWays(right, true, memo)
		rightFalse := CountWays(right, false, memo)
		total := (leftTrue + leftFalse) * (rightTrue + rightFalse)
		totalTrue := 0
		switch op {
		case '|':
			totalTrue = leftTrue*rightTrue + leftFalse*rightTrue + leftTrue*rightFalse
		case '&':
			totalTrue = leftTrue * rightTrue
		case '^':
			totalTrue = leftTrue*rightFalse + leftFalse*rightTrue
		default:
			panic("unsupported operator")
		}
		if result {
			ways += totalTrue
		} else {
			ways += total - totalTrue
		}
	}
	memo[key] = ways
	return ways
}

func Evaluate(exp string, result bool) int {
	if !expressionPattern.MatchString(exp) {
		panic("exp must look like 0/1 with &|^ operators")
	}
	if len(exp) > 31 {
		panic("expression exceeds safety bound")
	}
	return CountWays(exp, result, map[string]int{})
}

func RunQ814() {
	fmt.Printf("1^0|0|1 false=%d\n", Evaluate("1^0|0|1", false))
	fmt.Printf("0&0&0&1^1|0 true=%d\n", Evaluate("0&0&0&1^1|0", true))
}
