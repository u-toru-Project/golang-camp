package chapter06

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

type Gender int

const (
	Boy Gender = iota
	Girl
)

type Family struct {
	NumBoys  int
	NumGirls int
}

func NewFamily(numBoys, numGirls int) Family {
	if numBoys < 0 || numGirls < 0 {
		panic("counts must be non-negative")
	}
	return Family{NumBoys: numBoys, NumGirls: numGirls}
}

func SimulateChild(r *rand.Rand) Gender {
	r = rngOrDefault(r)
	if r.Intn(2) == 0 {
		return Boy
	}
	return Girl
}

func SimulateFamily(r *rand.Rand) Family {
	r = rngOrDefault(r)
	numBoys := 0
	for SimulateChild(r) == Boy {
		numBoys++
	}
	return NewFamily(numBoys, 1)
}

func SimulateApocalypse(numFamilies int, r *rand.Rand) float64 {
	if numFamilies <= 0 {
		panic("num_families must be positive")
	}
	if numFamilies > 1_000_000 {
		panic("num_families exceeds safety bound")
	}
	r = rngOrDefault(r)
	totalBoys := 0
	totalGirls := 0
	for range numFamilies {
		family := SimulateFamily(r)
		totalBoys += family.NumBoys
		totalGirls += family.NumGirls
	}
	return float64(totalBoys) / float64(totalBoys+totalGirls)
}

func ReadPositiveInt(prompt string, maxAttempts, maxValue int, inputFn func(string) string) int {
	if maxAttempts < 1 {
		panic("max_attempts must be >= 1")
	}
	if inputFn == nil {
		inputFn = func(p string) string {
			fmt.Print(p)
			var text string
			fmt.Scanln(&text)
			return text
		}
	}
	lastError := "invalid input"
	for range maxAttempts {
		text := strings.TrimSpace(inputFn(prompt))
		if text == "" || len(text) > 12 || !allDigits(text) {
			lastError = "enter a positive integer"
			continue
		}
		var value int
		fmt.Sscanf(text, "%d", &value)
		if value < 1 || value > maxValue {
			lastError = fmt.Sprintf("value must be in 1..%d", maxValue)
			continue
		}
		return value
	}
	panic("failed to read positive int: " + lastError)
}

func allDigits(text string) bool {
	for _, r := range text {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func rngOrDefault(r *rand.Rand) *rand.Rand {
	if r != nil {
		return r
	}
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RunQ607() {
	ratio := SimulateApocalypse(200, rand.New(rand.NewSource(0)))
	fmt.Printf("Proportion of boys: %v\n", ratio)
}
