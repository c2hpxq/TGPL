package main

import (
	"fmt"
	"os"
)

func sameChar(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	h := make(map[int32]int)
	for _, c := range s1 {
		h[c] += 1
	}
	for _, c := range s2 {
		h[c] -= 1
	}

	for _, v := range h {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	if len(os.Args) < 3 {
		return
	}

	fmt.Println(sameChar(os.Args[1], os.Args[2]))
}