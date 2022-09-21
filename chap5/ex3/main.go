// Usage
// go run main.go https://golang.org
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func visit(texts []string, n *html.Node) []string {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return texts
	}
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			texts = append(texts, text)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		texts = visit(texts, c)
	}
	return texts
}

func findTexts(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

func main() {
	for _, url := range os.Args[1:] {
		texts, err := findTexts(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findTexts: %v\n", err)
			continue
		}
		for _, text := range texts {
			fmt.Println(text)
		}
	}
}
