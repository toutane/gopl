// Exercise 3.7 render a png image of the Newton's fractal.
// https://fr.wikipedia.org/wiki/Fractale_de_Newton
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		ymin, xmin, ymax, xmax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

// Newton's calculation method:
// z(n+1) = z(n) - (f(z(n))/f'(z(n)))      with f(z(n)) = z^4 - 1
// z(n+1) -= (z(n)^4 -1)/(4z(n)^3)
// z(n+1) -= (z(n) -1/z(n)^3)/4

func newton(z complex128) color.Color {
	const iterations = 40
	const contrast = 7

	for n := uint8(0); n < iterations; n++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
