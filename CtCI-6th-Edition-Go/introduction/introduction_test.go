package introduction

import "testing"

func TestCompareAndSwap(t *testing.T) {
	if !CompareBinToHex("111001011", "1CB") || CompareBinToHex("100", "3") {
		t.Fatal("compare")
	}
	if ConvertToBase("10", 1) != -1 || DigitToValue('G') != -1 {
		t.Fatal("invalid")
	}
	array := []int{9, 3, 7, 1, 8}
	SwapMinMax(array)
	if array[0] != 1 || array[3] != 9 {
		t.Fatalf("%v", array)
	}
}
