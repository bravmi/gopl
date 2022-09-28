package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestCanBeParsed(t *testing.T) {
	text := `
		<html>
		<body>
			<p class="something" id="short">
				<span class="special">hi</span>
			</p>
			<br/>
		</body>
		</html>
	`
	// ...
	doc, err := html.Parse(strings.NewReader(text))
	assert.NoError(t, err)

	stdout := os.Stdout
	r, w, _ := os.Pipe()
	defer func() { os.Stdout = stdout }()
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	out := <-outC
	forEachNode(doc, startElement, endElement)
	_, err = html.Parse(strings.NewReader(out))
	assert.NoError(t, err)
}
