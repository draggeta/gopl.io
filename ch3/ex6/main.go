// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.

// Ex: implement supersampling to smooth colors

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
		pyS := py * 2
		y0 := float64(pyS)/(height*2)*(ymax-ymin) + ymin
		y1 := float64(pyS+1)/(height*2)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			pxS := px * 2
			x0 := float64(pxS)/(width*2)*(xmax-xmin) + xmin
			x1 := float64(pxS+1)/(width*2)*(xmax-xmin) + xmin
			z00 := complex(x0, y0)
			z10 := complex(x1, y0)
			z01 := complex(x0, y1)
			z11 := complex(x1, y1)

			z := (z00 + z10 + z01 + z11) / 4
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
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
