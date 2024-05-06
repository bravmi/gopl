package main

import "io"

type limitReader struct {
	r io.Reader
	n int
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	if lr.n <= 0 {
		return 0, io.EOF
	}
	if len(p) > lr.n {
		p = p[:lr.n]
	}
	n, err = lr.r.Read(p)
	lr.n -= n
	return
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &limitReader{r, n}
}
