// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// usage: go run chap8/ex7/main.go -depth 1 -limit 10 -start http://gopl.dev
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

type Link struct {
	url   string
	depth int
}

var origHost string

func sameHost(rawurl string) bool {
	u, err := url.Parse(rawurl)
	if err != nil {
		return false
	}
	return u.Host == origHost
}

func save(rawurl string) error {
	u, err := url.Parse(rawurl)
	if err != nil {
		return err
	}
	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	path := filepath.Join(u.Host, u.Path)
	if filepath.Ext(u.Path) == "" {
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

// !+
func main() {
	var depth, limit int
	var start string
	flag.IntVar(&depth, "depth", 3, "max depth to crawl")
	flag.IntVar(&limit, "limit", 5, "limit of the number of links to be crawled")
	flag.StringVar(&start, "start", "http://go.dev", "starting url to crawl")
	flag.Parse()

	u, err := url.Parse(start)
	if err != nil {
		log.Fatalf("failed to parse start url: %v", err)
	}
	origHost = u.Host

	worklist := make(chan []Link)
	unseenLinks := make(chan Link)

	go func() {
		worklist <- []Link{{start, 0}}
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundURLs := crawl(link.url)
				var foundLinks []Link
				for _, u := range foundURLs {
					foundLinks = append(foundLinks, Link{u, link.depth + 1})
				}
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	// n counts pending worklist sends; when it hits 0, no more can arrive.
	seen := make(map[string]bool)
	admitted := 0
	for n := 1; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				if link.depth <= depth && sameHost(link.url) && admitted < limit {
					n++
					unseenLinks <- link
					admitted++
				}
			}
		}
	}
	close(unseenLinks)
}

//!-
