// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset

const wlen = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := s.getIndex(x)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := s.getIndex(x)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wlen; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

// ex6.1
func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	return s
}

func (s *IntSet) Remove(x int) {
	idx, bit := x/wlen, x%wlen
	if idx < len(s.words) {
		s.words[idx] &^= 1 << bit
	}
}

func (s IntSet) Len() (count int) {
	for _, word := range s.words {
		for word != 0 {
			word &= word-1
			count++
		}
	}
	return
}

// ex6.2
func (s *IntSet) AddAll(keys ...int) {
	for _, key:= range keys {
		s.Add(key)
	}
}

// ex6.3
func (s *IntSet) getIndex(x int) (word int, bit uint) {
	return x/64, uint(x%64)
}

func (s *IntSet) getElem(word int, bit uint) int {
	return word*64 + int(bit)
}

// dual to Add, lol
func (s *IntSet) Delete(x int) {
	if !s.Has(x) {
		return
	}
	word, bit := s.getIndex(x)
	// bit operator &^ in Go represents set difference A-B
	s.words[word] &^= 1 << bit
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] &^= word
		} else {
			return
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] ^= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

// ex6.4
func (s IntSet) Elems() (res []int) {
	for i, word := range s.words {
		if word != 0 {
			for j := uint(0); j < wlen; j++ {
				if word & (1 << j) != 0 {
					res = append(res, s.getElem(i, j))
				}
			}
		}
	}
	return
}