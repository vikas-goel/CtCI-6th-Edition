package main

import (
	"fmt"
	"math/rand"
	"./stack"
)

func main() {
	capacity := 11
	s := stack.Stack{}.New(capacity, stack.Integers{})

	fmt.Print("Pushing elements:")
	for i := 0; i < capacity; i++ {
		value := rand.Intn(100)
		s.Push(value)
		fmt.Print(" ", value)
	}
	fmt.Println()

	s.Sort()
	s.PopAll()
}
