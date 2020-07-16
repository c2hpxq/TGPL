// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", commaFloat(os.Args[i]))
	}
}

//!+
// comma inserts commas in a float number.
func commaFloat(s string) string {
	r, err := regexp.Compile(`([+-]?)\s*(\d*)(\.(\d+))?`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while compiling regexp, %v\n", err)
	}

	group := r.FindStringSubmatch(s)
	return group[1] + comma(group[2]) + group[3]
}

//!-

func comma(s string) string {
	var b bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	rem := n%3
	for _, c := range s {
		if rem == 0 {
			rem = 3
			b.WriteByte(',')
		}
		b.WriteByte(byte(c))
		rem--
	}

	return b.String()
}