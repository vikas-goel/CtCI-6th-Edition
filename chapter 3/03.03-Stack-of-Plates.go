package main

import "fmt"
import (
	"math/rand"
	"./stack"
)

type StackOfPlates struct {
	threshold int
	dataType stack.Interface
	stacks []*stack.Stack
}

func (this StackOfPlates) New(t int, dt stack.Interface) *StackOfPlates {
	this.threshold = t
	this.dataType = dt
	this.stacks = make([]*stack.Stack, 0, 1)
	return &this
}

func (this *StackOfPlates) Empty() bool {
	if this == nil || this.stacks == nil || cap(this.stacks) == 0 {
		return true
	}

	return this.stacks[0].Empty()
}

func (this *StackOfPlates) Peek() (interface{}, bool) {
	if this.Empty() {
		return 0, false
	}

	// Peek the last stack in the slice.
	return this.stacks[len(this.stacks)-1].Pop()
}

func (this *StackOfPlates) Pop() (interface{}, bool) {
	if this.Empty() {
		return 0, false
	}

	// Pop the last stack in the slice.
	var tstack *stack.Stack = this.stacks[len(this.stacks)-1]
	value, ok := tstack.Pop()

	// If the last stack is empty after the operation, then delete the
	// stack from the slice.
	if tstack.Empty() {
		this.stacks = this.stacks[:len(this.stacks)-1]
	}

	return value, ok
}

func (this *StackOfPlates) PopAt(index int) (interface{}, bool) {
	if this.Empty() || len(this.stacks) <= index {
		return 0, false
	}

	value, ok := this.stacks[index].Pop()

	// If the stack becomes empty, then remove it from the slice.
	// Else, leave the stack in partial filled state.
	if this.stacks[index].Empty() {
		this.stacks = append(this.stacks[:index], this.stacks[index+1:]...)
	}

	return value, ok
}

func (this *StackOfPlates) Push(value interface{}) bool {
	if this == nil {
		return false
	}

	// Push to the last stack in the slice.
	var tstack *stack.Stack = this.stacks[len(this.stacks)-1]

	// If the last stack is full, then create a new stack and push to the
	// the new stack. Push the new stack to the slice as well.
	if tstack.Full() {
		tstack = stack.Stack{}.New(this.threshold, this.dataType)
		this.stacks = append(this.stacks, tstack)
	}

	return tstack.Push(value)
}

func (this *StackOfPlates) PopAll() {
	for _, s := range this.stacks {
		s.PopAll()
	}
}

func main() {
	threshold := 5
	s := StackOfPlates{}.New(threshold, stack.Integers{})
	fmt.Print("Pushing elements:")
	for i := 0; i < 30; i++ {
		value := rand.Intn(100)
		s.Push(value)
		fmt.Print(" ", value)
	}
	fmt.Println()

	fmt.Print("Popping stack# 1:")
	for i := 0; i < 2; i++ {
		value, _ := s.PopAt(1)
		fmt.Print(" ", value)
	}
	fmt.Println()

	fmt.Print("Popping elements:")
	for i := 0; i < 12; i++ {
		value, _ := s.Pop()
		fmt.Print(" ", value)
	}
	fmt.Println()

	fmt.Print("Pushing elements:")
	for i := 0; i < 15; i++ {
		value := rand.Intn(100)
		s.Push(value)
		fmt.Print(" ", value)
	}
	fmt.Println()

	fmt.Print("Popping stack# 1:")
	for i := 0; i < 2; i++ {
		value, _ := s.PopAt(1)
		fmt.Print(" ", value)
	}
	fmt.Println()

	fmt.Print("Popping stack# 1:")
	for i := 0; i < 5; i++ {
		value, _ := s.PopAt(1)
		fmt.Print(" ", value)
	}
	fmt.Println()

	s.PopAll()
}
