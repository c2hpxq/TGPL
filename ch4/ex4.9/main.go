// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCounts := make(map[string]int)    // counts of Unicode characters

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		wordCounts[in.Text()]++
	}
	fmt.Printf("rune\tcount\n")
	for w, c := range wordCounts {
		var s string
		if c > 1 {
			s = "s"
		}
		fmt.Printf("%s appears %d time%s\n", w, c, s)
	}
}

//!-
