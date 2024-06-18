// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
// usage:
// go run main.go -limit=100 -start=https://golang.org
// go run main.go
package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/bravmi/gopl/chap5/links"
)

// !+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string, limit int) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				more := f(item)
				if len(seen) >= limit {
					log.Println("breadthFirst: limit reached")
					return
				}
				worklist = append(worklist, more...)
			}
		}
	}
}

//!-breadthFirst

var origHost string

func save(rawurl string) error {
	url, err := url.Parse(rawurl)
	if err != nil {
		return err
	}
	if origHost == "" {
		origHost = url.Host
	}
	if url.Host != origHost {
		return nil
	}

	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	path := filepath.Join(url.Host, url.Path)
	if filepath.Ext(url.Path) == "" {
		path = filepath.Join(path, "index.html")
	}
	dir := filepath.Dir(path)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// !+crawl
func crawl(url string) []string {
	log.Println("crawl:", url)
	err := save(url)
	if err != nil {
		log.Println("crawl:", err)
	}
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

// !+main
func main() {
	limit := flag.Int("limit", 5, "limit of the number of links to be crawled")
	start := flag.String("start", "https://go.dev/", "start url")
	flag.Parse()
	// Crawl the web breadth-first,
	// starting from the start url.
	worklist := []string{*start}
	breadthFirst(crawl, worklist, *limit)
}

//!-main
