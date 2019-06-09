package mandelbrot

import (
	"image"
	"image/color"
	"math"
)

// Image is a dynamic representation of a Mandelbrot set
// with each pixel computed on demand.
type Image struct {
	Width   int
	Height  int
	CenterX float64
	CenterY float64
	Mag     float64
	Limit   int
}

// ColorModel returns the Image's color model.
func (p *Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (p *Image) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: p.Width, Y: p.Height},
	}
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (p *Image) At(x, y int) color.Color {
	px := 4.0 / math.Min(float64(p.Width), float64(p.Height)) / p.Mag
	x0 := p.CenterX - float64(p.Width)/2.0*px
	y0 := p.CenterY + float64(p.Height)/2.0*px
	c := Complex{
		x: x0 + float64(x)*px,
		y: y0 - float64(y)*px,
	}
	v := iterations(c, p.Limit)
	pc := pixelColor(v, p.Limit)
	return color.RGBA{pc[0], pc[1], pc[2], pc[3]}
}
