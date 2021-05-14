package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	const (
		ymin, xmin, ymax, xmax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func newton(z complex64) color.Color {
	const iterations = 40
	const contrast = 7

	for n := uint8(0); n < iterations; n++ {
		z -= (z - 1/(z*z*z)) / 4
		if abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func abs(z complex64) float32 {
	return float32(real(z)*real(z) + imag(z)*imag(z))
}
