package main

import (
	"fmt"
	"math/rand"
	"./stack"
)

type MyQueue struct {
	capacity, size int
	stack1, stack2 *stack.Stack
}

func (this MyQueue) New(capacity int, dataType stack.Interface) *MyQueue{
	this.capacity = capacity
	this.stack1 = stack.Stack{}.New(capacity, dataType)
	this.stack2 = stack.Stack{}.New(capacity, dataType)
	return &this
}

func (this *MyQueue) Empty() bool {
	if this == nil || this.size == 0 {
		return true
	}

	return false
}

func (this *MyQueue) Dequeue() (interface{}, bool) {
	if this.Empty() {
		return 0, false
	}

	// If the second stack is empty, move the elements from stack1 to it
	// to get the elements in the original order.
	if this.stack2.Empty() {
		for !this.stack1.Empty() {
			value, _ := this.stack1.Pop()
			this.stack2.Push(value)
		}
	}

	this.size--

	// Dequeue always haappens from the second stack.
	return this.stack2.Pop()
}

func (this *MyQueue) Enqueue(value interface{}) bool {
	if this == nil || this.size == this.capacity {
		return false
	}

	this.size++

	// Enqueue always haappens in the first stack.
	return this.stack1.Push(value)
}

func (this *MyQueue) DequeueAll() {
	fmt.Print("[")
	if !this.Empty() {
		value, _ := this.Dequeue()
		fmt.Print(value)
	}

	for !this.Empty() {
		value, _ := this.Dequeue()
		fmt.Print(", ", value)
	}
	fmt.Println("]")
}

func main() {
	capacity := 20
	q := MyQueue{}.New(capacity, stack.Integers{})

	fmt.Print("Enqueuing elements:")
	for i := 0; i < capacity/2; i++ {
		value := rand.Intn(100)
		q.Enqueue(value)
		fmt.Print(" ", value)
	}
	fmt.Println()

	fmt.Print("Dequeuing elements:")
	for i := 0; i < capacity/3; i++ {
		value, _ := q.Dequeue()
		fmt.Print(" ", value)
	}
	fmt.Println()

	fmt.Print("Enqueuing elements:")
	for {
		value := rand.Intn(100)
		if !q.Enqueue(value) {
			break
		}
		fmt.Print(" ", value)
	}
	fmt.Println()

	q.DequeueAll()
}
