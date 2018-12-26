package main

import "fmt"

// Assumes N > M and (i,j) > bits(M).
func RealToBinary(num float32) string {
	if num >= 1.0 || num <= 0.0 {
		return "ERROR"
	}

	binary  := make([]byte, 0)
	binary = append(binary, '0')
	binary = append(binary, '.')

	for num > 0 {
		if len(binary) >= 32 {
			return "ERROR"
		}

		num *= 2
		if num >= 1 {
			binary = append(binary, '1')
			num -= 1
		} else {
			binary = append(binary, '0')
		}
	}

	return string(binary)
}

func main() {
	var num float32 = 0.72
	fmt.Printf("binary(%v) = %v\n", num, RealToBinary(num))
}
