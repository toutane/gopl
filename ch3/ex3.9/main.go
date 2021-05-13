package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	render(w, r)
}

// Render function specify query parameters to draw function.
func render(out http.ResponseWriter, r *http.Request) {
	var width, height float64 = 1024, 1024
	var ymin, xmin, ymax, xmax float64 = -2, -2, 2, 2

	params := [6]string{"width", "height", "ymin", "xmin", "ymax", "xmax"}
	query := r.URL.Query()

	if len(query) != 0 {
		for _, param := range params {
			if len(query[param]) != 0 {
				switch param {
				case "width":
					width, _ = strconv.ParseFloat(query["width"][0], 64)
				case "height":
					height, _ = strconv.ParseFloat(query["height"][0], 64)
				case "ymin":
					ymin, _ = strconv.ParseFloat(query["ymin"][0], 64)
				case "xmin":
					xmin, _ = strconv.ParseFloat(query["xmin"][0], 64)
				case "ymax":
					ymax, _ = strconv.ParseFloat(query["ymax"][0], 64)
				case "xmax":
					xmax, _ = strconv.ParseFloat(query["xmax"][0], 64)
				}
			}
		}
	}
	//fmt.Fprintf(out, "%v, %v, %v, %v", ymin, xmin, ymax, xmax)
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	draw(out, img, width, height, ymin, xmin, ymax, xmax)
}

// Draw function set each pixel of image.
func draw(out http.ResponseWriter, img *image.RGBA, width, height, ymin, xmin, ymax, xmax float64) {
	for py := float64(0); py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := float64(0); px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			img.Set(int(px), int(py), newton(z))
		}
	}

	png.Encode(out, img)
}

// Newton function does maths and return color of pixel.
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
