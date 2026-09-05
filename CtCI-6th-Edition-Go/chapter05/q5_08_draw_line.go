package chapter05

import "fmt"

func ComputeByteNum(width, x, y int) int {
	return (width*y + x) / 8
}

// DrawLine draws a horizontal line from (x1, y) to (x2, y) on a bit-packed screen.
func DrawLine(screen []byte, width, x1, x2, y int) {
	startOffset := x1 % 8
	firstFullByte := x1 / 8
	if startOffset != 0 {
		firstFullByte++
	}
	endOffset := x2 % 8
	lastFullByte := x2 / 8
	if endOffset != 7 {
		lastFullByte--
	}
	for b := firstFullByte; b <= lastFullByte; b++ {
		screen[(width/8)*y+b] = 0xFF
	}
	startMask := byte(0xFF >> startOffset)
	endMask := byte(^(0xFF >> (endOffset + 1)))
	if x1/8 == x2/8 {
		mask := startMask & endMask
		screen[(width/8)*y+x1/8] |= mask
	} else {
		if startOffset != 0 {
			byteNumber := (width/8)*y + firstFullByte - 1
			screen[byteNumber] |= startMask
		}
		if endOffset != 7 {
			byteNumber := (width/8)*y + lastFullByte + 1
			screen[byteNumber] |= endMask
		}
	}
}

func PrintByte(value byte) {
	for i := 7; i >= 0; i-- {
		fmt.Print((value >> i) & 1)
	}
}

func PrintScreen(screen []byte, width int) {
	height := len(screen) * 8 / width
	for row := range height {
		for column := 0; column < width; column += 8 {
			PrintByte(screen[ComputeByteNum(width, column, row)])
		}
		fmt.Println()
	}
}

func RunQ508() {
	const width = 8 * 4
	const height = 15
	screen := make([]byte, width*height/8)
	DrawLine(screen, width, 8, 10, 2)
	PrintScreen(screen, width)
}
