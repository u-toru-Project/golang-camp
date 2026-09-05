package introduction

import "fmt"

func DigitToValue(character rune) int {
	if character >= '0' && character <= '9' {
		return int(character - '0')
	}
	if character >= 'A' && character <= 'F' {
		return 10 + int(character-'A')
	}
	if character >= 'a' && character <= 'f' {
		return 10 + int(character-'a')
	}
	return -1
}

func ConvertToBase(number string, baseValue int) int {
	if baseValue < 2 || (baseValue > 10 && baseValue != 16) {
		return -1
	}
	value := 0
	for _, character := range number {
		digit := DigitToValue(character)
		if digit < 0 || digit >= baseValue {
			return -1
		}
		value = value*baseValue + digit
	}
	return value
}

func CompareBinToHex(binary, hex string) bool {
	n1 := ConvertToBase(binary, 2)
	n2 := ConvertToBase(hex, 16)
	return n1 >= 0 && n2 >= 0 && n1 == n2
}

func GetMinIndex(array []int) int {
	if len(array) == 0 {
		panic("array must not be empty")
	}
	minIndex := 0
	for i := 1; i < len(array); i++ {
		if array[i] < array[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func GetMaxIndex(array []int) int {
	if len(array) == 0 {
		panic("array must not be empty")
	}
	maxIndex := 0
	for i := 1; i < len(array); i++ {
		if array[i] > array[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

func SwapMinMax(array []int) {
	if len(array) < 2 {
		return
	}
	minIndex, maxIndex := GetMinIndex(array), GetMaxIndex(array)
	array[minIndex], array[maxIndex] = array[maxIndex], array[minIndex]
}

func RunIntroduction() {
	fmt.Printf("111001011 vs 1CB: %t\n", CompareBinToHex("111001011", "1CB"))
	array := []int{9, 3, 7, 1, 8}
	SwapMinMax(array)
	fmt.Println(array)
}
