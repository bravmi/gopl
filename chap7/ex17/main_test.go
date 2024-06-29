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
		x        Element
		y        Element
		expected bool
	}{
		{
			Element{"div", "", nil},
			Element{"div", "", nil},
			true,
		},
		{
			Element{"div", "", nil},
			Element{"div", "id", nil},
			false,
		},
		{
			Element{"div", "", nil},
			Element{"div", "", []string{"class"}},
			false,
		},
		{
			Element{"div", "", []string{"class"}},
			Element{"div", "", []string{"class"}},
			true,
		},
		{
			Element{"div", "", []string{"class1", "class2"}},
			Element{"div", "", []string{"class1", "class2"}},
			true,
		},
		{
			Element{"div", "", []string{"class1", "class2"}},
			Element{"div", "", []string{"class2", "class1"}},
			true,
		},
		{
			Element{"div", "", []string{"class1", "class2"}},
			Element{"div", "", []string{"class1"}},
			true,
		},
		{
			Element{"div", "", []string{"class1", "class2"}},
			Element{"div", "", []string{"class1", "class2", "class3"}},
			false,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("contains(%v, %v)", test.x, test.y), func(t *testing.T) {
			assert.Equal(t, test.expected, test.x.contains(test.y))
		})
	}
}
