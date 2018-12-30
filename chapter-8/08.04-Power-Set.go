package main

import (
	"fmt"
	"math"
)

type Set []byte

// Create subsets of a given set.
func Subsets(set *Set) (subsets []Set) {
	subsets = make([]Set, 0, int(math.Pow(2, float64(set.Len()))))
	subsets = append(subsets, Set{})

	// For each element of the set, create new subsets by appending the
	// element to the existing subsets and merge the new subsets to the
	// existig one.
	for i := 0; i < set.Len(); i++ {
		for j, subsetLength := 0, len(subsets); j < subsetLength; j++ {
			newSubset := NewSet(&subsets[j], set.ElementAt(i))
			subsets = append(subsets, *newSubset)
		}
	}

	return subsets
}

func NewSet(source *Set, elem interface{}) *Set {
	set := source.Clone()
	return set.Append(elem)
}

func (this *Set) Append(elem interface{}) *Set {
	if this == nil {
		return nil
	}

	*this = append(*this, elem.(byte))
	return this
}

func (this *Set) Clone() *Set {
	if this == nil {
		return nil
	}

	var set Set = make([]byte, this.Len())
	copy(set, *this)

	return &set
}

func (this *Set) ElementAt(index int) interface{} {
	if this == nil || index > this.Len() {
		return 0
	}

	return (*this)[index]
}

func (this *Set) Len() int {
	if this == nil {
		return 0
	}

	return len(*this)
}

func main() {
	subsets := Subsets(&Set{'a', 'b', 'c', 'd', 'e', 'f'})
	fmt.Printf("Total subsets = %d\n%c\n", len(subsets), subsets)
}
