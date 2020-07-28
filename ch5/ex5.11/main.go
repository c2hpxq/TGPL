// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"log"
	"sort"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"linear algebra": {"calculus"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	order, containLoop := topoSort(prereqs)
	if containLoop {
		log.Fatal("contain loop!")
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, bool) {
	var order []string
	seen := make(map[string]bool)
	inStack := make(map[string]bool)
	var visitAll func(items []string) bool

	visitAll = func(items []string) bool {
		for _, item := range items {
			if inStack[item] {
				return true
			}
			if !seen[item] {
				inStack[item] = true
				seen[item] = true
				if visitAll(m[item]) {
					return true
				}
				order = append(order, item)
				inStack[item] = false
			}
		}
		return false
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	if visitAll(keys) {
		return nil, true
	}
	return order, false
}

//!-main
