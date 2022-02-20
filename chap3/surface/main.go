// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.

// usage: http://localhost:8080/?color=blue&zf=eggbox
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
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

func main() {
	file := flag.Bool("f", false, "write to 'surface.svg'")
	flag.Parse()

	color := "grey"
	if *file {
		fi, err := os.Create("surface.svg")
		defer fi.Close()
		if err != nil {
			log.Fatal("failed to create a file")
		}
		w := bufio.NewWriter(fi)
		svg(w, f, color)
		w.Flush()
	} else {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			if s := r.FormValue("color"); s != "" {
				color = s
			}
			zf := f
			if s := r.FormValue("zf"); s != "" {
				switch s {
				case "saddle":
					zf = saddle
				case "eggbox":
					zf = eggbox
				default:
					zf = f
				}
			}
			svg(w, zf, color)
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}

func svg(w io.Writer, f zfunc, color string) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// for 3.3 you could return z values here as well
			// and check how close they are to global min / max
			ax, ay, err := corner(i+1, j, f)
			bx, by, err := corner(i, j, f)
			cx, cy, err := corner(i, j+1, f)
			dx, dy, err := corner(i+1, j+1, f)
			if err != nil {
				fmt.Printf("skipping invalid polygon: %v\n", err)
				continue
			}
			fmt.Fprintf(w, "<polygon style='stroke: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n", color,
				ax, ay, bx, by, cx, cy, dx, dy)
		}
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

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}

func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

//!-
