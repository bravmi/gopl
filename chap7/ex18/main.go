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
	fmt.Println(node)
}

func ParseXML(r io.Reader) (*Element, error) {
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
				lines := strings.Split(string(tok), "\n")
				for _, line := range lines {
					line = strings.TrimSpace(line)
					if len(line) > 0 {
						parent.Children = append(parent.Children, CharData(line))
					}
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
	sb.WriteString(fmt.Sprintf("%*s", indent*2, ""))
	sb.WriteString(fmt.Sprintf("<%s", elem.Type.Local))
	for _, attr := range elem.Attr {
		sb.WriteString(fmt.Sprintf(" %s=%q", attr.Name.Local, attr.Value))
	}
	sb.WriteString(">\n")
	for _, child := range elem.Children {
		switch child := child.(type) {
		case *Element:
			sb.WriteString(child.Pretty(indent + 1))
		case CharData:
			sb.WriteString(fmt.Sprintf("%*s", (indent+1)*2, ""))
			sb.WriteString(fmt.Sprintf("%s\n", child))
		}
	}
	sb.WriteString(fmt.Sprintf("%*s", indent*2, ""))
	sb.WriteString(fmt.Sprintf("</%s>\n", elem.Type.Local))
	return sb.String()
}
