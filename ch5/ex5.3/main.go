// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	m := make(map[string]int)
	visit(m, doc)
	for k, v := range m {
		fmt.Printf("Data named %s appears %d times\n", k, v)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(dataCount map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		dataCount[n.Data]++
	}
	/*
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	 */
	if n.NextSibling != nil {
		visit(dataCount, n.NextSibling)
	}
	if n.FirstChild != nil {
		visit(dataCount, n.FirstChild)
	}
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
