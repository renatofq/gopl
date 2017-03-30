// This file is a derivative work of "mandelbrot"
// Original work Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Original work can be found at https://github.com/adonovan/gopl.io
// Derivative work Copyright © 2017 Renato Fernandes de Queioz.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See exercise 8.5 of The Go Programming Language (http://www.gopl.io/)
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"runtime"
	"strconv"
)

const maxConcurrency = 1024 // no special reason for this number

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

type operand struct {
	px, py int
	z      complex128
}

type result struct {
	px, py int
	color  color.Color
}

func main() {

	nConcurrent := runtime.NumCPU()
	if len(os.Args) == 2 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"invalid number of procesees %v\n",
				err)
		} else if i > maxConcurrency {
			fmt.Fprintf(os.Stderr,
				"number of processes (%d) is higher than max(%d)\n",
				i, maxConcurrency)
		} else if i <= 0 {
			fmt.Fprintf(os.Stderr,
				"number of processes (%d) cannot be lesser than or equal to 0\n",
				i)
		}
		} else {
			nConcurrent = i
		}
	}

	inch := make(chan operand, nConcurrent)
	outch := make(chan result, nConcurrent)
	defer close(outch)

	for i := 0; i < nConcurrent; i++ {
		go func() {
			for op := range inch {
				c := newton(op.z)
				outch <- result{
					px:    op.px,
					py:    op.py,
					color: c,
				}
			}
		}()
	}

	go func() {
		defer close(inch)

		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				inch <- operand{px, py, z}
			}
		}
	}()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < width*height; i++ {
		res := <-outch
		img.Set(res.px, res.py, res.color)
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func newton(z complex128) color.Color {
	const iterations = 100
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
