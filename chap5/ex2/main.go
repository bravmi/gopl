// Usage
// go run main.go https://golang.org
package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"

	"golang.org/x/net/html"
)

func visit(counts map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		counts = visit(counts, c)
	}
	return counts
}

func countNodes(url string) (map[string]int, error) {
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
	counts := map[string]int{}
	return visit(counts, doc), nil
}

func main() {
	for _, url := range os.Args[1:] {
		counts, err := countNodes(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "countNodes: %v\n", err)
			continue
		}
		keys := make([]string, 0, len(counts))
		for key := range counts {
			keys = append(keys, key)
		}
		sort.Slice(keys, func(i, j int) bool { return counts[keys[i]] < counts[keys[j]] })
		for _, key := range keys {
			fmt.Printf("%-10s => %d\n", key, counts[key])
		}
	}
}
