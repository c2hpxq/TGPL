package myslice

import (
	"unicode"
	"unicode/utf8"
)

func Reverse(ptr *[10]int) {
	for i, j := 0, len(ptr)-1; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}

func Rotate(s []int, n int) {
	var last []int
	last = append(last, s[:n]...)
	copy(s, s[n:])
	copy(s[len(s)-n:], last)
}

func RemoveContDup(ss []string) []string{
	ans := ss[:0]
	for _, s := range ss {
		if len(ans)==0 || ans[len(ans)-1] != s {
			ans = append(ans, s)
		}
	}
	return ans
}

func DeSpaceDup(b []byte) []byte {
	var last rune
	j := 0
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		if !unicode.IsSpace(r) || j == 0 || !unicode.IsSpace(last) {
			copy(b[j:j+size], b[i:i+size])
			j += size
		}
		last = r
		i += size
	}
	return b[:j]
}

func rev(b []byte) []byte {
	for i, j := 0, len(b)-1; i<j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func RevUTF8 (b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		rev(b[i:i+size])
		i += size
	}
	rev(b)
	return b
}