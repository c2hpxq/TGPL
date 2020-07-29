package main

import "sort"

func IsPalinDrome(s sort.Interface) bool {
	for i, j := 0, s.Len(); i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

type byteArray []byte

func (b byteArray) Len() int { return len(b) }
func (b byteArray) Less(i, j int) bool { return b[i] < b[j]}
func (b byteArray) Swap(i, j int) { b[i], b[j] = b[j], b[i]}

func main() {
	IsPalinDrome(byteArray([]byte("aaa")))
}