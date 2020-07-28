package main

import (
	"TGPL/ch4/myslice"
	"fmt"
)

func main() {
	a := [10]int{2: 10}
	myslice.Reverse(&a)
	myslice.Rotate(a[:], 2)
	ss := []string{"aa", "bb", "bb", "cc"}
	fmt.Println(myslice.RemoveContDup(ss))
	fmt.Println(ss)
	fmt.Println(a)

	s := "🏠哈  哈"
	fmt.Println([]byte(s))
	fmt.Println(string(myslice.RevUTF8(myslice.DeSpaceDup([]byte(s)))))
	fmt.Println([]byte(s))
	fmt.Println(s)
}
