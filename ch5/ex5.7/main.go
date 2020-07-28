// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

const templ = `{{range .Attr}} {{.Key}}: "{{.Val}}" {{end}}`
var report *template.Template
func main() {
	var err error
	report, err = template.New("report").Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

//!+startend
var depth int
func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		if err := report.Execute(os.Stdout, n); err != nil {
			log.Fatal(err)
		}
		if n.FirstChild ==nil {
			fmt.Println("/>")
		} else {
			fmt.Println(">")
			depth++
		}

	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if (len(text)>0) {
			for _, line := range strings.Split(text, "\n") {
				fmt.Printf("%*s%s\n", depth*2, "", strings.TrimSpace(line))
			}
		}
	case html.CommentNode:
		fmt.Printf("<!--%s-->\n", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			return
		}
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

//!-startend
