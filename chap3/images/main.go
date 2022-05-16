// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"bufio"
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"os"
	"strconv"
)

func main() {
	const (
		width, height         = 720, 720
		xmin, xmax    float64 = -2, +2
		ymin, ymax    float64 = -2, +2
		epsX                  = (xmax - xmin) / width
		epsY                  = (ymax - ymin) / height
		avg                   = false
		zoom                  = 1.0
	)
	f := mandelbrot

	tofile := flag.Bool("f", false, "write to 'images.png'")
	flag.Parse()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	if *tofile {
		fi, err := os.Create("images.png")
		defer fi.Close()
		if err != nil {
			log.Fatal("failed to create a file")
		}

		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				// Image point (px, py) represents complex value z.
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			}
		}

		w := bufio.NewWriter(fi)
		png.Encode(w, img) // NOTE: ignoring errors
		w.Flush()
	} else {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/" {
				http.NotFound(w, r)
				return
			}
			w.Header().Set("Content-Type", "image/png")
			// avg
			avg := avg
			if s := r.FormValue("avg"); s != "" {
				b, err := strconv.ParseBool(s)
				if err != nil {
					log.Printf("invalid avg param: %v", err)
				} else {
					avg = b
				}
			}
			// f
			f := f
			if s := r.FormValue("f"); s != "" {
				switch s {
				case "mandelbrot":
					f = mandelbrot
				case "newton":
					f = newton
				case "acos":
					f = acos
				case "sqrt":
					f = sqrt
				default:
					f = mandelbrot
				}
			}
			// xmin, xmax, ymin, ymax, zoom
			xmin, xmax := xmin, xmax
			ymin, ymax := ymin, ymax
			zoom := zoom
			fparams := []string{"xmin", "xmax", "ymin", "ymax", "zoom"}
			for _, k := range fparams {
				if s := r.FormValue(k); s != "" {
					v, err := strconv.ParseFloat(s, 64)
					if err != nil {
						log.Printf("invalid %q param: %v", k, err)
					} else {
						switch k {
						case "xmin":
							xmin = v
						case "xmax":
							xmax = v
						case "ymin":
							ymin = v
						case "ymax":
							ymax = v
						case "zoom":
							zoom = v
						}
					}
				}
			}
			lenX := xmax - xmin
			lenY := ymax - ymin
			epsX := lenX / width
			epsY := lenY / height

			midX := xmin + lenX/2
			xmin = midX - lenX/2/zoom
			xmax = midX + lenX/2/zoom
			midY := ymin + lenY/2
			ymin = midY - lenY/2/zoom
			ymax = midY + lenY/2/zoom

			for py := 0; py < height; py++ {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					// Image point (px, py) represents complex value z.
					z := complex(x, y)
					if avg {
						colors := []color.Color{
							f(complex(x-epsX, y-epsY)),
							f(complex(x+epsX, y-epsY)),
							f(complex(x-epsX, y+epsY)),
							f(complex(x+epsX, y+epsY)),
						}
						img.Set(px, py, average(colors))
					} else {
						img.Set(px, py, f(z))
					}
				}
			}

			png.Encode(w, img)
		})
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
	}
}

func average(colors []color.Color) color.Color {
	var R, G, B, A uint8
	n := len(colors)
	for _, c := range colors {
		r, g, b, a := c.RGBA()
		R += uint8(r / uint32(n))
		G += uint8(g / uint32(n))
		B += uint8(b / uint32(n))
		A += uint8(a / uint32(n))
	}
	return color.RGBA{R, G, B, A}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{0, 255 - contrast*n*2, 255 - contrast*n, 255}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	const eps = 1e-6
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < eps {
			roots := []complex128{
				complex(1, 0), complex(-1, 0), complex(0, 1), complex(0, -1),
			}
			switch {
			case cmplx.Abs(z-roots[0]) < eps:
				return color.RGBA{255 - contrast*i, 0, 0, 255}
			case cmplx.Abs(z-roots[1]) < eps:
				return color.RGBA{0, 255 - contrast*i, 0, 255}
			case cmplx.Abs(z-roots[2]) < eps:
				return color.RGBA{0, 0, 255 - contrast*i, 255}
			case cmplx.Abs(z-roots[3]) < eps:
				return color.RGBA{255 - contrast*i, 255 - contrast*i, 0, 255}
			default:
				return color.Gray{255 - contrast*i}
			}
		}
	}
	return color.Black
}
