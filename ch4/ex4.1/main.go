package main

import (
	"crypto/sha256"
	"fmt"
)

var popcount [256]int8

func init() {
	for i := range popcount {
		popcount[i] = popcount[i/2] + int8(i&1)
	}
}

func main() {
	dig1 := sha256.Sum256([]byte("rune"))
	dig2 := sha256.Sum256([]byte("ansi"))
	ans := 0
	for i, _ := range dig1 {
		ans += int(popcount[dig1[i]^dig2[i]])
	}
	fmt.Printf("Popcount of xor = %d, Type = %[1]T\n", ans)
}
