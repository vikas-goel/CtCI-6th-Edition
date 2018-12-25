package stack

import  "fmt"

type Interface interface {
	At(index int) interface{}
	Equal(i, j int) bool
	Less(i, j int) bool
	LessValue(v1, v2 interface{}) bool
	New(cpacity int) Interface
	Set(index int, value interface{})
}

type Stack struct {
	value Interface
	min []int
	capacity, size int
}

func (this Stack) New(capacity int, dataType Interface) *Stack {
	if capacity < 2 {
		return nil
	}

	this.capacity = capacity
	this.min = make([]int, this.capacity)
	this.value = dataType.New(this.capacity)

	return &this
}

func (this *Stack) Empty() bool {
	if this == nil || this.size == 0 {
		return true
	}

	return false
}

func (this *Stack) Full() bool {
	if this == nil || this.size < this.capacity {
		return false
	}

	return true
}

func (this *Stack) Min() (interface{}, bool) {
	if this.Empty() {
		return 0, false
	}

	return this.value.At(this.min[this.size-1]), true
}

func (this *Stack) Peek() (interface{}, bool) {
	if this.Empty() {
		return 0, false
	}

	// Return the value of the min index.
	return this.value.At(this.size-1), true
}

func (this *Stack) Pop() (interface{}, bool) {
	if this.Empty() {
		return 0, false
	}

	this.size--
	return this.value.At(this.size), true
}

func (this *Stack) Push(value interface{}) bool {
	if this == nil || this.size == this.capacity {
		return false
	}

	// Add the value to the top.
	this.value.Set(this.size, value)

	// Set the current min to itself.
	this.min[this.size] = this.size

	// If the current value is greater than the previous min, then the
	// current min is same as the previous min.
	if this.size > 0 && this.value.Less(this.min[this.size-1], this.size) {
		this.min[this.size] = this.min[this.size-1]
	}

	this.size++
	return true
}

func (this *Stack) Sort() {
	if this == nil || this.size <= 1 {
		return
	}

	temp := Stack{}.New(this.capacity, this.value)
	value, _ := this.Pop()
	temp.Push(value)

	for !this.Empty() {
		value1, ok := this.Peek()
		value2, _ := temp.Peek()
		for ok && !this.value.LessValue(value1, value2) {
			value1, _ = this.Pop()
			temp.Push(value1)

			value1, ok = this.Peek()
			value2, _ = temp.Peek()
		}

		value1, ok = this.Pop()
		if !ok {
			break
		}

		value2, ok = temp.Peek()
		for ok && this.value.LessValue(value1, value2) {
			value2, _ = temp.Pop()
			this.Push(value2)

			value2, ok = temp.Peek()
		}

		temp.Push(value1)
	}

	for !temp.Empty() {
		value2, _ := temp.Pop()
		this.Push(value2)
	}
}

func (this *Stack) PopAll() {
	fmt.Print("[")
	if !this.Empty() {
		value, _ := this.Peek()
		fmt.Print(value)
		this.Pop()
	}

	for ; !this.Empty(); this.Pop() {
		value, _ := this.Peek()
		fmt.Print(", ", value)
	}
	fmt.Println("]")
}

func (this *Stack) PopMinAll() {
	fmt.Print("[")
	if !this.Empty() {
		value, _ := this.Peek()
		min, _ := this.Min()
		fmt.Print("{", value, ",", min, "}")
		this.Pop()
	}

	for ; !this.Empty(); this.Pop() {
		value, _ := this.Peek()
		min, _ := this.Min()
		fmt.Print(", {", value, ",", min, "}")
	}
	fmt.Println("]")
}
