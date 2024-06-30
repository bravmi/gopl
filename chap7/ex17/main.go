// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 214.
//!+

// Xmlselect prints the text of selected elements of an XML document.

// usage:
// http -F http://www.w3.org/TR/2006/REC-xml11-20060816 | go run github.com/bravmi/gopl/chap7/ex17 div.body div.div1 h2
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"slices"
)

type Element struct {
	Name    string
	Id      string
	Classes []string
}

func (e Element) String() string {
	s := e.Name
	if e.Id != "" {
		s += "#" + e.Id
	}
	if len(e.Classes) > 0 {
		s += "." + strings.Join(e.Classes, ".")
	}
	return s
}

var nameRe = regexp.MustCompile(`^[\w]+`)
var idRe = regexp.MustCompile(`#([\w]+)`)
var classRe = regexp.MustCompile(`\.([\w]+)`)

func ParseSelector(sel string) Element {
	name := nameRe.FindString(sel)
	idMatch := idRe.FindStringSubmatch(sel)
	classMatch := classRe.FindAllStringSubmatch(sel, -1)
	id := ""
	if len(idMatch) > 1 {
		id = idMatch[1]
	}
	var classes []string
	if len(classMatch) > 0 {
		for _, classMatch := range classMatch {
			classes = append(classes, classMatch[1])
		}
	}
	return Element{Name: name, Id: id, Classes: classes}
}

func fromToken(e xml.StartElement) Element {
	name := e.Name.Local
	var classes []string
	id := ""
	for _, attr := range e.Attr {
		if attr.Name.Local == "class" {
			classes = strings.Fields(attr.Value)
			slices.Sort(classes)
		}
		if attr.Name.Local == "id" {
			id = attr.Value
		}
	}
	return Element{Name: name, Id: id, Classes: classes}
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []Element // stack of elements
	args := make([]Element, 0)
	for _, arg := range os.Args[1:] {
		args = append(args, ParseSelector(arg))
	}

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, fromToken(tok)) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			stackStr := make([]string, len(stack))
			for i, e := range stack {
				stackStr[i] = e.String()
			}
			if containsAll(stack, args) {
				fmt.Printf("%s: %s\n", strings.Join(stackStr, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []Element, y []Element) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0].contains(y[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func (x Element) contains(y Element) bool {
	if y.Id != "" && x.Id != y.Id {
		return false
	}
	for _, class := range y.Classes {
		if !slices.Contains(x.Classes, class) {
			return false
		}
	}
	return x.Name == y.Name
}

//!-
