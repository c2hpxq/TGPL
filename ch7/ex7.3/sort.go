// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

//!+
type tree struct {
	value       int
	offset int
	left, right *tree
}

func (t *tree) String() string {
	return t.vString()
}

func (t *tree) hBuild(vbegin int) string {
	if t == nil {
		return ""
	}
	newBegin := vbegin + len(strconv.Itoa(t.value))
	var ss []string
	for _, s := range []string{t.left.hBuild(newBegin), fmt.Sprintf("%*s%d", vbegin, "", t.value), t.right.hBuild(newBegin)} {
		if s != "" {
			ss = append(ss, s)
		}
	}
	return strings.Join(ss, "\n")
}

func (t *tree) hString() string {
	return t.hBuild(0)
}

func (t *tree) vString() string {
	if t == nil {
		return ""
	}
	t.calcOffset(0)
	next := []*tree{t}
	var ss []string
	for len(next) > 0 {
		var tmp []*tree
		last := 0
		for _, node := range next {
			if node == nil {
				continue
			}
			s := fmt.Sprintf("%*s%d", node.offset - last, "", node.value)
			ss = append(ss, s)
			last += len(s)
			tmp = append(tmp, node.left, node.right)
		}
		ss = append(ss, "\n")
		next = tmp
	}
	return strings.Join(ss, "")
}

func (t *tree) calcOffset(vbegin int) (width int) {
	if t == nil {
		return 0
	}
	width += t.left.calcOffset(vbegin)
	t.offset = vbegin + width
	width += len(strconv.Itoa(t.value))
	width += t.right.calcOffset(vbegin + width)
	return
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	fmt.Println(root)
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

//!-
