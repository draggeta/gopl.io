// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.

// Ex: Color each polygon based on its height. Peaks are red, valleys blue.

package main

import (
	"fmt"
	"math"
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

func heightDelta() (float64, float64) {
	// Max height of Z is 1 at x,y coordinates (0,0), min height is somewhere between pi and 2pi,
	// as the f function divides the sine by increasingly higher numbers. Looking beyond 2pi is
	// useless.
	// minimum amount of change for the x, y and z axis is 1.0/cells * xyrange
	// sqrt((1.0/cells * xyrange)^2 + (0.0/cells * xyrange)^2) = 1.0/cells * xyrange
	// find minimum value between pi and 2pi with steps of delta z over 2 (increase granularity)
	var m float64

	var dz = 1.0 / cells * xyrange

	for i := math.Pi; i < 2*math.Pi; i += (dz / 2) {
		h := f(0.0, i)
		if h < m {
			m = h
		}
	}

	// range and minimum
	m = m * zscale
	r := zscale - m
	return r, m
}

func main() {
	dh, m := heightDelta()

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var c string
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			avg := (az + bz + cz + dz) / 4
			b := int(255.0 / dh * (zscale - avg))
			r := int(255.0 / dh * (avg - m))
			c = fmt.Sprintf("rgb(%d,0,%d)", r, b)

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, c)
		}
	}
	fmt.Println("</svg>")
	// fmt.Println(heightDiff())
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	sz := z * zscale
	return sx, sy, sz
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
