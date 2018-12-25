package stack

type Runes []rune

func (this Runes) At(index int) interface{} {
	return this[index]
}

func (this Runes) Equal(i, j int) bool {
	return this[i] == this[j]
}

func (this Runes) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Runes) LessValue(v1, v2 interface{}) bool {
	return v1.(rune) < v2.(rune)
}

func (this Runes) New(capacity int) Interface {
	var data Runes = make([]rune, capacity)
	return data
}

func (this Runes) Set(index int, value interface{}) {
	this[index] = value.(rune)
}
