// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bravmi/gopl/chap5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

type Link struct {
	url   string
	depth int
}

// !+
func main() {
	var depth int
	flag.IntVar(&depth, "depth", 3, "max depth to crawl")
	flag.Parse()

	worklist := make(chan []Link)
	unseenLinks := make(chan Link)

	// Add command-line arguments to worklist.
	go func() {
		var seeds []Link
		for _, url := range flag.Args() {
			seeds = append(seeds, Link{url, 0})
		}
		worklist <- seeds
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				if link.depth > depth {
					continue
				}
				foundUrls := crawl(link.url)
				var foundLinks []Link
				for _, url := range foundUrls {
					foundLinks = append(foundLinks, Link{url, link.depth + 1})
				}
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}

//!-
