package main

import "io"

type byteCounter struct {
	w io.Writer
	c int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &byteCounter{w: w}
	return cw, &cw.c
}

func (cw *byteCounter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	cw.c += int64(n)
	return n, err
}
