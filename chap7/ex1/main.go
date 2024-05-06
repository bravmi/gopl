// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

//!-bytecounter

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	c := 0
	for scanner.Scan() {
		c++
	}
	*wc += WordCounter(c)
	return c, nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	c := 0
	for scanner.Scan() {
		c++
	}
	*lc += LineCounter(c)
	return c, nil
}

func main() {
	//!+main
	var c ByteCounter
	_, err := c.Write([]byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
	//!-main
}
