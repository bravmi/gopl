// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
// usage:
// go run main.go https://golang.org
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		err := outline(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "outline: %v\n", err)
			continue
		}
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

// !+forEachNode
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

// !+startend
var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		attrs := ""
		for _, a := range n.Attr {
			attrs += fmt.Sprintf(" %s=%q", a.Key, a.Val)
		}
		format := "%*s<%s%s>\n"
		if n.FirstChild == nil {
			format = "%*s<%s%s/>\n"
		}
		fmt.Printf(format, depth*2, "", n.Data, attrs)
		depth++
	case html.TextNode:
		trimmed := strings.TrimSpace(n.Data)
		if trimmed != "" {
			// to correctly indent multiple lines
			lines := strings.Split(trimmed, "\n")
			for _, l := range lines {
				fmt.Printf("%*s%s\n", depth*2, "", l)
			}
		}
	case html.CommentNode:
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

//!-startend
