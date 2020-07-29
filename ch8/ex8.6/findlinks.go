// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	type linkListWithDepth struct {
		linkList []string
		depth int
	}
	type linkWithDepth struct {
		link string
		depth int
	}
	worklist := make(chan linkListWithDepth, 20)  // lists of URLs, may have duplicates
	unseenLinks := make(chan linkWithDepth, 20) // de-duplicated URLs

	initList := linkListWithDepth{linkList: os.Args[1:], depth: 0}
	// Add command-line arguments to worklist.
	go func() { worklist <- initList }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for ld := range unseenLinks {
				if ld.depth >= 3 {
					continue
				}
				foundLinks := crawl(ld.link)
				go func() { worklist <- linkListWithDepth{linkList:foundLinks, depth: ld.depth+1} }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)

	for n := 1; n > 0; n-- {
		ld := <-worklist
		fmt.Printf("%d %d %d\n", len(worklist), n, ld.depth)
		for _, link := range ld.linkList {
			if !seen[link] {
				seen[link] = true
				if ld.depth < 3 {
					n++
				}
				unseenLinks <- linkWithDepth{link, ld.depth}
			}
		}
	}
}

//!-
