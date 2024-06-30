package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseXML(t *testing.T) {
	xml := `<A><B><C>hello</C><D>abc</D></B><C>world</C></A>`
	expected := strings.TrimSpace(`
<A>
  <B>
    <C>
      hello
    </C>
    <D>
      abc
    </D>
  </B>
  <C>
    world
  </C>
</A>`)

	r := strings.NewReader(xml)
	node, err := ParseXML(r)
	assert.NoError(t, err)
	nodeStr := strings.TrimSpace(node.String())
	assert.Equal(t, expected, nodeStr)
}
