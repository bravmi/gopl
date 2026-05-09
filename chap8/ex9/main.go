// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 250.

// usage: go run chap8/ex9/main.go -v ...
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type File struct {
	Dir  string
	Size int64
}

// !+
func main() {
	// ...determine roots...

	//!-
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//!+
	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan File)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	//!-

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	nfiles := make(map[string]int64)
	nbytes := make(map[string]int64)
loop:
	for {
		select {
		case file, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles[file.Dir]++
			nbytes[file.Dir] += file.Size
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes)
		}
	}

	printDiskUsage(roots, nfiles, nbytes) // final totals
	//!+
	// ...select loop...
}

//!-

func printDiskUsage(roots []string, nfiles, nbytes map[string]int64) {
	for _, root := range roots {
		fmt.Printf("%s: %d files  %.1f GB\n", root, nfiles[root], float64(nbytes[root])/1e9)
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
// !+walkDir
func walkDir(root string, dir string, n *sync.WaitGroup, fileSizes chan<- File) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(root, subdir, n, fileSizes)
		} else {
			info, _ := entry.Info()
			fileSizes <- File{Dir: root, Size: info.Size()}
		}
	}
}

//!-walkDir

// !+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.DirEntry {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
