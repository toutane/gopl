package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	ymin, xmin, ymax, xmax = -2, -2, 2, 2
	width, height          = 2024, 2024
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			img.Set(px, py, sampling(x, y))
		}
	}
	png.Encode(os.Stdout, img)
}

func sampling(x, y float64) color.Color {
	var r, g, b []uint8
	const offX = (xmax - xmin) / width
	const offY = (ymax - ymin) / height
	var subLenX = []float64{-offX, offX}
	var subLenY = []float64{-offY, offY}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			z := complex(x+subLenX[i], y+subLenY[j])
			red, green, blue := mandelbrot(z)
			r = append(r, red)
			g = append(g, green)
			b = append(b, blue)
		}
	}
	red := r[0]/4 + r[1]/4 + r[2]/4 + r[3]/4
	green := g[0]/4 + g[1]/4 + g[2]/4 + g[3]/4
	blue := b[0]/4 + b[1]/4 + b[2]/4 + b[3]/4
	return color.RGBA{red, green, blue, 255}

}

func mandelbrot(z complex128) (uint8, uint8, uint8) {
	const iterations = 200
	const contrast = 10

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return contrast * n, contrast * n, 0
		}
	}
	return 111, 0, 236
}
