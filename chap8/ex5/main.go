// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.

// usage:
// go run chap3/surface/main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sync"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type zfunc func(x, y float64) float64

type polygon struct {
	ax, ay float64
	bx, by float64
	cx, cy float64
	dx, dy float64
}

type position struct {
	i, j int
}

func main() {
	workers := flag.Int("w", 1, "number of workers")
	flag.Parse()

	color := "grey"
	fi, err := os.Create("surface.svg")
	if err != nil {
		log.Fatal("failed to create a file")
	}
	defer fi.Close()
	w := bufio.NewWriter(fi)
	svg(w, f, color, *workers)
	w.Flush()
}

func svg(w io.Writer, f zfunc, color string, workers int) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	input := make(chan position, 10)
	output := make(chan *polygon, 10)

	go func() {
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				input <- position{i: i, j: j}
			}
		}
		close(input)
	}()

	var wg sync.WaitGroup
	for k := 0; k < workers; k++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range input {
				ax, ay, aerr := corner(p.i+1, p.j, f)
				bx, by, berr := corner(p.i, p.j, f)
				cx, cy, cerr := corner(p.i, p.j+1, f)
				dx, dy, derr := corner(p.i+1, p.j+1, f)
				if aerr != nil || berr != nil || cerr != nil || derr != nil {
					fmt.Println("skipping invalid polygon")
					continue
				}
				output <- &polygon{
					ax: ax, ay: ay,
					bx: bx, by: by,
					cx: cx, cy: cy,
					dx: dx, dy: dy,
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	for p := range output {
		fmt.Fprintf(w, "<polygon style='stroke: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			color, p.ax, p.ay, p.bx, p.by, p.cx, p.cy, p.dx, p.dy)
	}

	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int, zf zfunc) (float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := zf(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, fmt.Errorf("invalid height")
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
