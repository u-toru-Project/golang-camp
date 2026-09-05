package chapter01

import (
	"fmt"
	"slices"
)

func IsPermutationBySorting(original, valueToTest string) bool {
	if len(original) !== len(valueToTest) {
		return  false
	}

