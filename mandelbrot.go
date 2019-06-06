package mandelbrot

import "math"
import "image"

// Complex is our custom type to keep parity with other implementations.
type Complex struct{ x, y float64 }

func escapes(p Complex) bool {
	return p.x*p.x+p.y*p.y > 4.0
}

func next(p Complex, p0 Complex) Complex {
	return Complex{
		x: p.x*p.x - p.y*p.y + p0.x,
		y: 2.0*p.x*p.y + p0.y,
	}
}

func iterations(p Complex, limit int) int {
	p0 := p
	for i := 0; i <= limit; i++ {
		if escapes(p) {
			return i
		}
		p = next(p, p0)
	}
	return limit
}

func iterArea(
	w int, h int,
	center Complex,
	mag float64,
	iterCallback func(i, x, y int, p Complex)) {
	cx, cy := center.x, center.y
	p := 4.0 / math.Min(float64(w), float64(h)) / mag
	x0 := cx - float64(w)/2.0*p
	y0 := cy + float64(h)/2.0*p
	i := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			iterCallback(i, x, y, Complex{x: x0 + float64(x)*p, y: y0 - float64(y)*p})
			i++
		}
	}
}

// Color is an RGBA byte sequence.
type Color [4]uint8

// PALETTE is a set of colors stolen from another Mandelbrot set.
var PALETTE = [16]Color{
	{66, 30, 15, 255},
	{25, 7, 26, 255},
	{9, 1, 47, 255},
	{4, 4, 73, 255},
	{0, 7, 100, 255},
	{12, 44, 138, 255},
	{24, 82, 177, 255},
	{57, 125, 209, 255},
	{134, 181, 229, 255},
	{211, 236, 248, 255},
	{241, 233, 191, 255},
	{248, 201, 95, 255},
	{255, 170, 0, 255},
	{204, 128, 0, 255},
	{153, 87, 0, 255},
	{106, 52, 3, 255},
}

// BLACK is used to represent the "lake."
var BLACK = Color{0, 0, 0, 255}

func color(v int, limit int) Color {
	if v > 0 && v < limit {
		return PALETTE[v%16]
	}
	return BLACK
}

// DrawPix draws a Mandelbrot set in and returns an RGBA buffer.
func DrawPix(
	width int,
	height int,
	cx float64,
	cy float64,
	mag float64,
	limit int) []uint8 {
	bs := width * height * 4
	img := make([]uint8, bs)
	iterArea(
		width,
		height,
		Complex{x: cx, y: cy},
		mag,
		func(i, x, y int, p Complex) {
			v := iterations(p, limit)
			pixel := color(v, limit)
			offset := (i * 4)
			for b := 0; b < 4; b++ {
				img[offset+b] = pixel[b]
			}
		},
	)
	return img
}

// DrawNRGBA draws a Mandelbrot set in and returns an RGBA buffer.
func DrawNRGBA(
	width int, height int,
	cx float64, cy float64,
	mag float64, limit int) *image.NRGBA {
	pix := DrawPix(width, height, cx, cy, mag, limit)
	return &image.NRGBA{
		Pix:    pix,
		Stride: width * 4,
		Rect:   image.Rect(0, 0, width, height),
	}
}
