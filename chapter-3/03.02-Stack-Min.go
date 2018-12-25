package main

import (
	"math/rand"
	"./stack"
)

func main() {
	capacity := 11
	s := stack.Stack{}.New(capacity, stack.Integers{})
	for i := 0; i < capacity; i++ {
		s.Push(rand.Intn(100))
	}

	s.PopMinAll()
}
