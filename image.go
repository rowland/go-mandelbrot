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
	// Palette defaults to UltraFractalPalette.
	Palette color.Palette
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
	pc := p.colorForIterations(v, p.Limit)
	return pc
}

// UltraFractalPalette is a palette of colors allegedly from Ultra Fractal program.
var UltraFractalPalette = []color.Color{
	color.RGBA{0, 0, 0, 255},
	color.RGBA{66, 30, 15, 255},
	color.RGBA{25, 7, 26, 255},
	color.RGBA{9, 1, 47, 255},
	color.RGBA{4, 4, 73, 255},
	color.RGBA{0, 7, 100, 255},
	color.RGBA{12, 44, 138, 255},
	color.RGBA{24, 82, 177, 255},
	color.RGBA{57, 125, 209, 255},
	color.RGBA{134, 181, 229, 255},
	color.RGBA{211, 236, 248, 255},
	color.RGBA{241, 233, 191, 255},
	color.RGBA{248, 201, 95, 255},
	color.RGBA{255, 170, 0, 255},
	color.RGBA{204, 128, 0, 255},
	color.RGBA{153, 87, 0, 255},
	color.RGBA{106, 52, 3, 255},
}

func (p *Image) colorForIterations(iterations, limit int) color.Color {
	if p.Palette == nil {
		p.Palette = UltraFractalPalette
	}
	if iterations > 0 && iterations < limit {
		return p.Palette[iterations%(len(p.Palette)-1)+1]
	}
	return p.Palette[0]
}
