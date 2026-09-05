package chapter16

import (
	"fmt"
	"strings"
)

var smalls = []string{
	"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
	"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen",
	"Seventeen", "Eighteen", "Nineteen",
}

var tens = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

var bigs = []string{"", "Thousand", "Million", "Billion"}

func ConvertIntToEnglish(number int) string {
	if number == 0 {
		return smalls[0]
	}
	if number < 0 {
		return "Negative " + convertPositive(-int64(number))
	}
	return convertPositive(int64(number))
}

func convertPositive(number int64) string {
	parts := make([]string, 0)
	chunkCount := 0
	for number > 0 {
		if number%1000 != 0 {
			chunk := convertChunk(int(number % 1000))
			scale := bigs[chunkCount]
			if scale == "" {
				parts = append([]string{chunk}, parts...)
			} else {
				parts = append([]string{chunk + " " + scale}, parts...)
			}
		}
		number /= 1000
		chunkCount++
	}
	return strings.Join(parts, " ")
}

func convertChunk(number int) string {
	parts := make([]string, 0)
	if number >= 100 {
		parts = append(parts, smalls[number/100], "Hundred")
		number %= 100
	}
	if number >= 10 && number <= 19 {
		parts = append(parts, smalls[number])
	} else if number >= 20 {
		parts = append(parts, tens[number/10])
		number %= 10
	}
	if number >= 1 && number <= 9 {
		parts = append(parts, smalls[number])
	}
	return strings.Join(parts, " ")
}

func RunQ1608() {
	value := 832787436
	fmt.Printf("%d in English is %s\n", value, ConvertIntToEnglish(value))
}
