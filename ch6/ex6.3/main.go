package main

import (
	"fmt"
	"gopl.io/ch6/intset"
)

func main() {
	var s, t intset.IntSet
	s.AddAll(1, 2, 3)
	t.AddAll(3, 4, 5)
	s.SymmetricDifference(&t)
	fmt.Println(&s)
	fmt.Println(s.Elems())
}

