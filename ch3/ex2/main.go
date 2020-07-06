// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("This command takes exactly one argument")
		os.Exit(1)
	}

	sf := os.Args[1]
	if sf != "eggbox" && sf != "ripple" && sf != "saddle" {
		fmt.Println("Valid options are: 'eggbox', 'ripple', 'saddle")
		os.Exit(1)
	}
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, sf)
			bx, by := corner(i, j, sf)
			cx, cy := corner(i, j+1, sf)
			dx, dy := corner(i+1, j+1, sf)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, sf string) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y, sf)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64, sf string) float64 {
	var output float64
	switch sf {
	case "eggbox": // don't know what I'm doing: https://mathcurve.com/surfaces.gb/boiteaoeufs/boiteaoeufs.shtml
		output = 0.1 * (math.Sin(x) + math.Sin(y))
	case "ripple":
		r := math.Hypot(x, y) // distance from (0,0)
		output = math.Sin(r) / r
	case "saddle": // don't know what I'm doing: https://mathcurve.com/surfaces.gb/translation/translation.shtml
		xsq := x * x
		ysq := y * y
		output = xsq*xsq/49000 - ysq/300
	}
	return output
}

//!-
