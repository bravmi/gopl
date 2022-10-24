// Usage
// go run main.go https://golang.org img
// go run main.go https://golang.org h1 h2 h3 h4
package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	var nodes []*html.Node

	pre := func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, name := range names {
				if n.Data == name {
					nodes = append(nodes, n)
				}
			}
		}
	}
	forEachNode(doc, pre, nil)

	return nodes
}

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

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s url tag1 tag2 ...", os.Args[0])
	}
	url := os.Args[1]
	names := os.Args[2:]

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	nodes := ElementsByTagName(doc, names...)
	for _, node := range nodes {
		log.Println(node)
	}
}
