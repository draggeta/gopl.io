// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.

// Ex: create a Newton's method color fractal

package main

import (
	"encoding/binary"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
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
			c := 0 + contrast*i
			m := 0 + contrast*i
			y := 0 + contrast*i
			k := 0 + contrast*i
			if math.Round(real(z)) == 1 {
				c = 255
			} else if math.Round(real(z)) == -1 {
				m = 255
			} else if imag(z) < 0 {
				y = 255
			} else if imag(z) > 0 {
				k = 0
			}
			// fmt.Println(real(z), imag(z))
			return color.CMYK{c, m, y, k}
		}
	}
	return color.Black
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint32(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			by := make([]byte, 4)
			binary.LittleEndian.PutUint32(by, math.MaxUint32-(contrast*n))
			r, g, b, a := by[0], by[1], by[2], by[3]
			return color.RGBA{r, g, b, a}
		}
	}
	return color.Black
}

//!-
