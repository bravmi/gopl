// usage:
// http -F golang.org | go run main.go quote_slide0
package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: ID")
		os.Exit(1)
	}

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	node := ElementByID(doc, os.Args[1])
	fmt.Printf("node: %v\n", node)
}

func ElementByID(doc *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					return true
				}
			}
		}
		return false
	}
	return forEachNode(doc, pre, nil)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if pre(n) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, pre, post)
		if node != nil {
			return node
		}
	}

	if post != nil {
		if post(n) {
			return n
		}
	}

	return nil
}
