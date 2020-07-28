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
		FindID(url, "js-clientConfig")
	}
}

func FindID(url string, id string ) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	startElement := func (n *html.Node) bool {
		return idMatch(n, id)
	}
	//!+call
	fmt.Println(forEachNode(doc, startElement, nil))
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if pre(n) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		m := forEachNode(c, pre, post)
		if m != nil {
			return m
		}
	}

	if post != nil {
		if post(n) {
			return n
		}
	}

	return nil
}

//!-forEachNode

//!+startend
var depth int
func idMatch(n *html.Node, id string) bool {
	for _, a := range n.Attr {
		if strings.ToLower(a.Key) == "id" && a.Val == id{
			return true
		}
	}
	return false
}

//!-startend
