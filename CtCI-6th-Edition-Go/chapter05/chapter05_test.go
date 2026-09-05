package chapter05

import (
	"math"
	"strings"
	"testing"
)

func TestUpdateBits(t *testing.T) {
	cases := []struct {
		n, m, i, j, expected int
	}{
		{0b10000000000, 0b10011, 2, 6, 0b10001001100},
		{1024, 19, 2, 6, 1100},
		{0b10000000000, 0b10011, 0, 4, 0b10000010011},
		{0b1011001, 0b011, 2, 4, 0b1001101},
		{-1, 0, 31, 31, math.MaxInt32},
		{0, 1, 31, 31, math.MinInt32},
		{0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0},
	}
	for _, tc := range cases {
		if got := UpdateBits(tc.n, tc.m, tc.i, tc.j); got != tc.expected {
			t.Fatalf("UpdateBits(%d,%d,%d,%d)=%d want %d", tc.n, tc.m, tc.i, tc.j, got, tc.expected)
		}
	}
}

func TestPrintBinary(t *testing.T) {
	cases := []struct {
		number   float64
		expected string
	}{
		{0.5, ".1"},
		{0.25, ".01"},
		{0.125, ".001"},
		{0.625, ".101"},
		{0, "ERROR"},
		{1, "ERROR"},
		{-0.5, "ERROR"},
		{1.5, "ERROR"},
		{0.72, "ERROR"},
		{0.1, "ERROR"},
	}
	for _, tc := range cases {
		if PrintBinary(tc.number) != tc.expected || PrintBinary2(tc.number) != tc.expected {
			t.Fatalf("%v: %s / %s want %s", tc.number, PrintBinary(tc.number), PrintBinary2(tc.number), tc.expected)
		}
	}
}

func TestFlipBitToWin(t *testing.T) {
	cases := []struct {
		number, expected int
	}{
		{0b0, 1},
		{0b111, 4},
		{0b10011100111, 4},
		{0b10110110111, 6},
		{0b11011101111, 8},
		{0b1, 2},
		{0b10, 2},
		{0b101, 3},
	}
	for _, tc := range cases {
		a, err := FlipBitToWin(tc.number)
		b, err2 := FlipBitToWinAlt(tc.number)
		if err != nil || err2 != nil || a != tc.expected || b != tc.expected {
			t.Fatalf("%b: %d / %d want %d", tc.number, a, b, tc.expected)
		}
	}
	if _, err := FlipBitToWin(-1); err == nil {
		t.Fatal("negative")
	}
	if _, err := FlipBitToWinAlt(-1); err == nil {
		t.Fatal("negative alt")
	}
}

func TestNextNumber(t *testing.T) {
	nextCases := [][2]int{{13948, 13967}, {6, 9}, {1, 2}}
	for _, tc := range nextCases {
		if GetNext(tc[0]) != tc[1] || GetNextArith(tc[0]) != tc[1] || GetNextSlow(tc[0]) != tc[1] {
			t.Fatalf("next %d", tc[0])
		}
		if CountOnes(tc[0]) != CountOnes(tc[1]) || tc[1] <= tc[0] {
			t.Fatal("next popcount")
		}
	}
	prevCases := [][2]int{{13948, 13946}, {6, 5}, {2, 1}}
	for _, tc := range prevCases {
		if GetPrev(tc[0]) != tc[1] || GetPrevArith(tc[0]) != tc[1] || GetPrevSlow(tc[0]) != tc[1] {
			t.Fatalf("prev %d got %d / %d / %d", tc[0], GetPrev(tc[0]), GetPrevArith(tc[0]), GetPrevSlow(tc[0]))
		}
	}
	for i := 0; i <= 200; i++ {
		if GetNext(i) != GetNextArith(i) || GetNext(i) != GetNextSlow(i) {
			t.Fatalf("next disagree %d", i)
		}
		if GetPrev(i) != GetPrevArith(i) || GetPrev(i) != GetPrevSlow(i) {
			t.Fatalf("prev disagree %d: %d %d %d", i, GetPrev(i), GetPrevArith(i), GetPrevSlow(i))
		}
	}
	if GetNext(0) != -1 || GetNextArith(0) != -1 || GetNextSlow(0) != -1 || HasValidNext(0) {
		t.Fatal("no next for 0")
	}
	if GetNext(math.MaxInt32) != -1 {
		t.Fatal("no next for max")
	}
	if GetPrev(1) != -1 || GetPrev(3) != -1 || GetPrev(7) != -1 || HasValidPrev(1) || HasValidPrev(-1) {
		t.Fatal("no prev")
	}
	if CountOnes(0) != 0 || CountOnes(1) != 1 || CountOnes(6) != 2 || CountOnes(-1) != 32 {
		t.Fatal("count ones")
	}
	if CountZeros(-1) != 0 || CountZeros(0) != 32 {
		t.Fatal("count zeros")
	}
	if !strings.Contains(ToFullBinaryString(1), "1") || len(ToFullBinaryString(1)) != 32 {
		t.Fatal(ToFullBinaryString(1))
	}
}

func TestDebugger(t *testing.T) {
	if !IsPowerOfTwo(1) || !IsPowerOfTwo(2) || !IsPowerOfTwo(8) || !IsPowerOfTwo(1<<30) {
		t.Fatal("powers")
	}
	if IsPowerOfTwo(0) || IsPowerOfTwo(3) || IsPowerOfTwo(6) || IsPowerOfTwo(-8) {
		t.Fatal("non powers")
	}
	demo := DemonstrateMismatch(0.625)
	if demo.DecimalStr != "0.625" || !strings.HasPrefix(demo.BinaryFraction, ".101") || demo.DecimalStr == demo.BinaryFraction {
		t.Fatal(demo)
	}
	hidden := DemonstrateMismatch(0.1)
	if !strings.HasPrefix(hidden.DecimalStr, "0.1") || hidden.BinaryFraction != "ERROR" {
		t.Fatal(hidden)
	}
	if demo.Ieee754Hex == demo.BinaryFraction || demo.Ieee754Hex != "3fe4000000000000" {
		t.Fatal(demo.Ieee754Hex)
	}
	for _, value := range []float64{0.5, 0.625, 0.25} {
		if BinaryStringCorrect(value) != PrintBinary(value) {
			t.Fatal(value)
		}
	}
	if DecimalStringWrong(0.625) != demo.DecimalStr || DoubleToHexBits(0.625) != demo.Ieee754Hex {
		t.Fatal("mismatch helpers")
	}
	if DecimalStringWrong(0.625) != "0.625" || DoubleToHexBits(0.625) != "3fe4000000000000" {
		t.Fatal(DecimalStringWrong(0.625), DoubleToHexBits(0.625))
	}
}

func TestConversion(t *testing.T) {
	cases := [][3]int{{29, 15, 2}, {15, 29, 2}, {0, 0, 0}, {42, 42, 0}, {-1, -1, 0}, {-1, 0, 32}, {math.MinInt32, 0, 1}, {7, 8, 4}}
	for _, tc := range cases {
		if BitSwapRequired(tc[0], tc[1]) != tc[2] || BitSwapRequired2(tc[0], tc[1]) != tc[2] {
			t.Fatalf("%d %d -> %d / %d want %d", tc[0], tc[1], BitSwapRequired(tc[0], tc[1]), BitSwapRequired2(tc[0], tc[1]), tc[2])
		}
	}
	if BitSwapRequired(23432, 512132) != BitSwapRequired2(23432, 512132) {
		t.Fatal("random pair")
	}
}

func TestSwapOddEvenBits(t *testing.T) {
	if SwapOddEvenBits(0b10) != 0b01 || SwapOddEvenBits(0b01) != 0b10 || SwapOddEvenBits(0b1010) != 0b0101 || SwapOddEvenBits(0b0101) != 0b1010 {
		t.Fatal("pairs")
	}
	aaaaBits := uint32(0xaaaaaaaa)
	fiveBits := uint32(0x55555555)
	aaaa := int(int32(aaaaBits))
	fives := int(int32(fiveBits))
	if SwapOddEvenBits(aaaa) != fives || SwapOddEvenBits(fives) != aaaa {
		t.Fatal("masks")
	}
	for _, value := range []int{0, 1, 2, 103217, -1, math.MinInt32, math.MaxInt32} {
		if SwapOddEvenBits(SwapOddEvenBits(value)) != value {
			t.Fatal(value)
		}
	}
	if SwapOddEvenBits(0) != 0 {
		t.Fatal("zero")
	}
}

func TestDrawLine(t *testing.T) {
	const width = 32
	screen := make([]byte, width*15/8)
	DrawLine(screen, width, 8, 10, 2)
	if screen[ComputeByteNum(width, 8, 2)] != 0xE0 {
		t.Fatalf("%x", screen[ComputeByteNum(width, 8, 2)])
	}
	bytesPerRow := width / 8
	for row := range 15 {
		if row == 2 {
			continue
		}
		for b := range bytesPerRow {
			if screen[row*bytesPerRow+b] != 0 {
				t.Fatalf("row %d byte %d", row, b)
			}
		}
	}
	full := make([]byte, width/8)
	DrawLine(full, width, 0, 7, 0)
	if full[0] != 0xFF || full[1] != 0 {
		t.Fatal(full)
	}
	partial := make([]byte, width/8)
	DrawLine(partial, width, 3, 20, 0)
	if partial[0] != 0x1F || partial[1] != 0xFF || partial[2] != 0xF8 || partial[3] != 0 {
		t.Fatal(partial)
	}
	single := make([]byte, 16/8)
	DrawLine(single, 16, 0, 0, 0)
	if single[0] != 0x80 || single[1] != 0 {
		t.Fatal(single)
	}
	if ComputeByteNum(32, 0, 0) != 0 || ComputeByteNum(32, 8, 0) != 1 || ComputeByteNum(32, 8, 2) != 9 {
		t.Fatal("byte num")
	}
}
