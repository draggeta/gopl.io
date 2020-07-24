// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	width, height = 1024, 1024
)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		var xmin, ymin, xmax, ymax float64 = -2, -2, +2, +2
		var zoom float64 = 100

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		for k, v := range r.Form {
			var err error
			switch k {
			case "xmin":
				if xmin, err = strconv.ParseFloat(v[0], 64); err != nil {
					fmt.Println(err)
				}
			case "xmax":
				if xmax, err = strconv.ParseFloat(v[0], 64); err != nil {
					fmt.Println(err)
				}
			case "ymin":
				if ymin, err = strconv.ParseFloat(v[0], 64); err != nil {
					fmt.Println(err)
				}
			case "ymax":
				if ymax, err = strconv.ParseFloat(v[0], 64); err != nil {
					fmt.Println(err)
				}
			case "zoom":
				if zoom, err = strconv.ParseFloat(v[0], 64); err != nil {
					fmt.Println(err)
				}
			}
		}

		ydif := (ymax - ymin) / 2
		ymax = (ymax - ydif) + ydif/zoom*100
		ymin = (ymin + ydif) - ydif/zoom*100
		xdif := (xmax - xmin) / 2
		xmax = (xmax - xdif) + xdif/zoom*100
		xmin = (xmin + xdif) - xdif/zoom*100

		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
			}
		}
		png.Encode(w, img) // NOTE: ignoring errors

	}
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
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
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
