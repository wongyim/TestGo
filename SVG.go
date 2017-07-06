package main

import (
	"fmt"
	"log"
	"math"
)

const (
	width, height = 1200, 800           // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var max, min float64

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", 1600, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_,_= corner(i+1, j)
			_,_= corner(i, j)
			_,_= corner(i, j+1)
			_,_= corner(i+1, j+1)
		}
	}
	color:=
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay:= corner(i+1, j)
			bx, by:= corner(i, j)
			cx, cy:= corner(i, j+1)
			dx, dy:= corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke=\"%s\"/>\n",
				ax, ay, bx, by, cx, cy, dx, dy,color)
		}
	}
	fmt.Println("</svg>")
	log.Println(max, min)
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if z > max {
		max = z
	}
	if z < min {
		min = z
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return float64(sx), sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	ret := math.Sin(r) / r
	return ret
}
