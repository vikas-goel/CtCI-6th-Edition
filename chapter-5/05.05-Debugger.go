package main

import "fmt"

func PowerOf2(number int) bool {
	return number & (number-1) == 0
}

func main() {
	for _, n := range []int{1, 2, 3, 4, 5, 10, 16, 20, 32, 36, 40} {
		fmt.Println(n, "is power of 2?", PowerOf2(n))
	}
}
