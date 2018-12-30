package main

import (
	"fmt"
	"log"
	"os"
)

type Interface interface {
	Swap(x, y int)
	Println()
}

func Permutations(set Interface, start, end int) {
	if start == end {
		set.Println()
	} else {
		for i := start; i <= end; i++ {
			set.Swap(start, i)
			Permutations(set, start+1, end)
			set.Swap(start, i)
		}
	}
}

type Charset struct {
	elems []byte
}

func (this Charset) Swap(x, y int) {
	this.elems[x], this.elems[y] = this.elems[y], this.elems[x]
}

func (this Charset) Println() {
	fmt.Printf("%c\n", this.elems)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "<string>")
	}

	orig := []byte(os.Args[1])
	Permutations(Charset{orig}, 0, len(orig)-1)
}
