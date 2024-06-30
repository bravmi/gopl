package main

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSelector(t *testing.T) {
	tests := []struct {
		sel      string
		expected Element
	}{
		{"div", Element{"div", "", nil}},
		{"div#id", Element{"div", "id", nil}},
		{"div.class", Element{"div", "", []string{"class"}}},
		{"div.class1.class2", Element{"div", "", []string{"class1", "class2"}}},
		{"div#id.class1.class2", Element{"div", "id", []string{"class1", "class2"}}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("ParseSelector(%s)", test.sel), func(t *testing.T) {
			assert.Equal(t, test.expected, ParseSelector(test.sel))
		})
	}
}

func TestFromToken(t *testing.T) {
	tests := []struct {
		token    xml.StartElement
		expected Element
	}{
		{
			xml.StartElement{Name: xml.Name{Local: "div"}},
			Element{"div", "", nil},
		},
		{
			xml.StartElement{Name: xml.Name{Local: "div"}, Attr: []xml.Attr{{Name: xml.Name{Local: "class"}, Value: "class1 class2"}}},
			Element{"div", "", []string{"class1", "class2"}},
		},
		{
			xml.StartElement{Name: xml.Name{Local: "div"}, Attr: []xml.Attr{{Name: xml.Name{Local: "id"}, Value: "id"}}},
			Element{"div", "id", nil},
		},
		{
			xml.StartElement{Name: xml.Name{Local: "div"}, Attr: []xml.Attr{{Name: xml.Name{Local: "id"}, Value: "id"}, {Name: xml.Name{Local: "class"}, Value: "class1 class2"}}},
			Element{"div", "id", []string{"class1", "class2"}},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("fromToken(%v)", test.token), func(t *testing.T) {
			assert.Equal(t, test.expected, fromToken(test.token))
		})
	}

}

func TestElementContains(t *testing.T) {
	tests := []struct {
		xSel     string
		ySel     string
		expected bool
	}{
		{"div", "div", true},
		{"div", "div#id", false},
		{"div", "div.class", false},
		{"div", "div.class1.class2", false},
		{"div", "div#id.class1.class2", false},
		{"div#id", "div", true},
		{"div#id", "div#id", true},
		{"div#id", "div.class", false},
		{"div#id", "div.class1.class2", false},
		{"div#id", "div#id.class1.class2", false},
		{"div.class", "div", true},
		{"div.class", "div#id", false},
		{"div.class", "div.class", true},
		{"div.class1", "div.class1.class2", false},
		{"div.class1", "div#id.class1.class2", false},
		{"div.class1.class2", "div", true},
		{"div.class1.class2", "div#id", false},
		{"div.class1.class2", "div.class1", true},
	}

	for _, test := range tests {
		x := ParseSelector(test.xSel)
		y := ParseSelector(test.ySel)
		t.Run(fmt.Sprintf("contains(%v, %v)", test.xSel, test.ySel), func(t *testing.T) {
			assert.Equal(t, test.expected, x.contains(y))
		})
	}
}
