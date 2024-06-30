// usage:
// http -F http://www.w3.org/TR/2006/REC-xml11-20060816 | go run github.com/bravmi/gopl/chap7/ex18
// echo '<A><B><C>hello</C><D>abc</D></B><C>world</C></A>' | go run github.com/bravmi/gopl/chap7/ex18
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface {
	String() string
	Pretty(indent int) string
} // CharData or *Element

type CharData string

func (c CharData) String() string {
	return string(c)
}

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	node, err := ParseXML(os.Stdin)
	if err != nil {
		fmt.Printf("Error parsing XML: %v\n", err)
		return
	}
	fmt.Print(node)
}

func ParseXML(r io.Reader) (Node, error) {
	dec := xml.NewDecoder(r)
	var stack []*Element // stack of element nodes

	var root *Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading token: %v\n", err)
			return nil, err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			elem := &Element{Type: tok.Name, Attr: tok.Attr}
			if len(stack) == 0 {
				root = elem
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, elem)
			}
			stack = append(stack, elem)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				s := strings.TrimSpace(string(tok))
				if len(s) > 0 {
					parent.Children = append(parent.Children, CharData(s))
				}
			}
		}
	}

	if root == nil {
		fmt.Println("No root element found")
		return nil, fmt.Errorf("no root element found")
	}
	return root, nil
}

func (elem *Element) String() string {
	return elem.Pretty(0)
}

func (elem *Element) Pretty(indent int) string {
	var sb strings.Builder
	padding := strings.Repeat(" ", indent*2)
	sb.WriteString(fmt.Sprintf("%s<%s", padding, elem.Type.Local))
	for _, attr := range elem.Attr {
		sb.WriteString(fmt.Sprintf(" %s=%q", attr.Name.Local, attr.Value))
	}
	sb.WriteString(">\n")
	for _, child := range elem.Children {
		sb.WriteString(child.Pretty(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", padding, elem.Type.Local))
	return sb.String()
}

func (c CharData) Pretty(indent int) string {
	padding := strings.Repeat(" ", indent*2)
	return fmt.Sprintf("%s%s\n", padding, c)
}
