package stack

type Integers []int

func (this Integers) At(index int) interface{} {
	return this[index]
}

func (this Integers) Equal(i, j int) bool {
	return this[i] == this[j]
}

func (this Integers) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Integers) LessValue(v1, v2 interface{}) bool {
	return v1.(int) < v2.(int)
}

func (this Integers) New(capacity int) Interface {
	var data Integers = make([]int, capacity)
	return data
}

func (this Integers) Set(index int, value interface{}) {
	this[index] = value.(int)
}
