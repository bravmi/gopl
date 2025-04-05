// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.

// usage:
// go run chap3/images/main.go -w 14
package main

import (
	"bufio"
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	const (
		width, height         = 4096, 4096
		xmin, xmax    float64 = -2, +2
		ymin, ymax    float64 = -2, +2
	)

	fileName := "mandelbrot.png"
	maxWorkers := runtime.GOMAXPROCS(0)
	workers := flag.Int("w", maxWorkers, "number of workers")
	flag.Parse()
	if *workers < 1 {
		*workers = maxWorkers
	}

	startTime := time.Now()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	fi, err := os.Create(fileName)
	if err != nil {
		log.Fatal("failed to create a file")
	}
	defer fi.Close()

	imageRows := make(chan int, height)
	go func() {
		for py := 0; py < height; py++ {
			imageRows <- py
		}
		close(imageRows)
	}()

	var wg sync.WaitGroup
	for w := 0; w < *workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for py := range imageRows {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					// Image imagePoint (px, py) represents complex value z.
					z := complex(x, y)
					img.Set(px, py, mandelbrot(z))
				}
			}
		}()
	}
	wg.Wait()

	w := bufio.NewWriter(fi)
	err = png.Encode(w, img)
	if err != nil {
		log.Fatal("failed to encode an image")
	}
	w.Flush()
	log.Printf("rendered in %s with %d workers",
		time.Since(startTime).Truncate(time.Millisecond), *workers)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{G: 255 - contrast*n*2, B: 255 - contrast*n, A: 255}
		}
	}
	return color.Black
}
