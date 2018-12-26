package main

import "fmt"

// Starting coordinate is (x=0, y=0).
func DrawLine(screen []byte, width, x1, x2, y uint) (line []byte) {
	line = make([]byte, 0)

	var heightOffset uint = width / 8 * y
	var indexStart, indexEnd uint = x1/8+heightOffset, x2/8+heightOffset

	var bitStart, bitEnd uint = x1 % 8, x2 % 8
	var maskFirst, maskLast byte = 0xFF >> bitStart, 0xFF << (7-bitEnd)

	if indexStart == indexEnd {
		// The start and end positions are in the same byte.
		line = append(line, screen[indexStart] & maskFirst & maskLast)
	} else {
		// Prepare the starting byte.
		line = append(line, screen[indexStart] & maskFirst)

		// Copy the intermediate bytes as-is.
		for s := indexStart+1; s < indexEnd; s++ {
			line = append(line, screen[s])
		}

		// Prepare the ending byte.
		line = append(line, screen[indexEnd] & maskLast)
	}

	return
}

func main() {
	var width, x1, x2, y uint

	screen := []byte{20, 12, 15, 5, 8, 9, 30, 100, 75}
	fmt.Printf("Screen(s=%v): %b\n", 8*len(screen), screen)

	width, x1, x2, y = 24, 10, 22, 2
	fmt.Printf("(%2v,%2v,%2v,%v): %b\n",
		width, x1, x2, y,DrawLine(screen, width, x1, x2, y))

	width, x1, x2, y = 24, 5, 20, 1
	fmt.Printf("(%2v,%2v,%2v,%v): %b\n",
		width, x1, x2, y,DrawLine(screen, width, x1, x2, y))
}
