// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch5/links"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst



//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	//!+crawl
	var order int
	crawl := func(url string) []string {
		if strings.HasPrefix(url, os.Args[1]) {
			fmt.Printf("fetching %s\n", url)
			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("length of content: %d\n", len(b))
			err = ioutil.WriteFile(strconv.Itoa(order) + ".html", b, 0644)
			if err != nil {
				log.Fatal(err)
			}
			order++
		}
		list, err := links.Extract(url)
		if err != nil {
		log.Print(err)
	}
		return list
	}

	//!-crawl
	breadthFirst(crawl, os.Args[1:])
}

//!-main
