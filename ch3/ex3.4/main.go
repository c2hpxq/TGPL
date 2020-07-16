// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "web" {
		handler := func(rw http.ResponseWriter, r *http.Request) {
			h, w := height, width
			r.ParseForm()
			var err error
			v, ok := r.Form["height"]
			if ok {
				h, err = strconv.Atoi(v[0])
				if err != nil {
					fmt.Fprintln(rw, "invalid height")
					return
				}
			}

			v, ok = r.Form["width"]
			if ok {
				w, err = strconv.Atoi(v[0])
				if err != nil {
					fmt.Fprintln(rw, "invalid height")
					return
				}
			}

			rw.Header().Set("Content-Type", "image/svg+xml")

			Gensvg(rw, h, w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}

	Gensvg(os.Stdout, height, width)
}

func Gensvg(out io.Writer, height, width int) {
	xyscale := float64(width / 2 / xyrange) // pixels per x or y unit
	zscale := float64(height) * 0.4        // pixels per z unit
	fmt.Fprintf(out,"<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	zmax, zmin := -100.0, 100.0
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			t, _ := f(x, y)
			zmax = math.Max(t, zmax)
			zmin = math.Min(t, zmin)
		}
	}
	fmt.Println(zmax, zmin)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			zij, nanz := f(x, y)
			ax, ay, nana := corner(i+1, j, xyscale, zscale)
			bx, by,  nanb:= corner(i, j, xyscale, zscale)
			cx, cy, nanc := corner(i, j+1, xyscale, zscale)
			dx, dy, nand := corner(i+1, j+1, xyscale, zscale)
			if !(nanz||nana||nanb||nanc||nand) {
				fmt.Fprintf(out,"<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='rgb(%d, 0, %d)'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, int((zij-zmin)/(zmax-zmin)*255), int((zmax-zij)/(zmax-zmin)*255))
			}
		}
	}
	fmt.Fprintln(out,"</svg>")
}

func corner(i, j int, xyscale, zscale float64) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, nan := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nan
}

func f(x, y float64) (float64, bool) {
	//r := math.Hypot(x, y) // distance from (0,0)
	//if r < 0.001 {
	//	return 1.0
	//}
	//return math.Sin(r) / r
	res := 0.2 * (math.Cos(x) + math.Cos(y))
	return res, math.IsNaN(res) || math.IsInf(res, +1) || math.IsInf(res, -1)
}

//!-
