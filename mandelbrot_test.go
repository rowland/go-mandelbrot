package mandelbrot

import (
	"bytes"
	"image/color"
	"image/png"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterations(t *testing.T) {
	assert := assert.New(t)
	p1 := Complex{x: 0.9, y: 0.9}
	assert.Equal(iterations(p1, 10), 1)
	p2 := Complex{x: 0.4, y: 0.4}
	assert.Equal(iterations(p2, 10), 8)
	p3 := Complex{x: 0.1, y: 0.01}
	assert.Equal(iterations(p3, 10), 10)
}

func TestPixelColor(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(pixelColor(0, 10), color.RGBA{0, 0, 0, 255})
	assert.Equal(pixelColor(10, 10), color.RGBA{0, 0, 0, 255})
	assert.Equal(pixelColor(4, 10), color.RGBA{0, 7, 100, 255})
}

func TestIterAreaCount(t *testing.T) {
	assert := assert.New(t)
	w := 6
	h := 4
	c := Complex{0.0, 0.0}
	mag := 1.0
	count := 0
	iterArea(w, h, c, mag, func(i, x, y int, _point Complex) {
		count++
	})
	assert.Equal(count, w*h)
}

func TestIterAreaRange(t *testing.T) {
	assert := assert.New(t)
	w := 6
	h := 4
	c := Complex{0.0, 0.0}
	mag := 1.0
	minX := 0.0
	minY := 0.0
	maxX := 0.0
	maxY := 0.0
	var img [24]Complex
	iterArea(w, h, c, mag, func(i, _x, _y int, point Complex) {
		img[i] = point
		x, y := point.x, point.y
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	})
	assert.Equal(img[0], Complex{x: -3.0, y: 2.0})
	assert.Equal(img[(w*h-1)], Complex{x: 2.0, y: -1.0})
	assert.Equal(minX, -3.0)
	assert.Equal(maxX, 2.0)
	assert.Equal(minY, -1.0)
	assert.Equal(maxY, 2.0)
}

func TestDrawPix(t *testing.T) {
	assert := assert.New(t)
	w := 6
	h := 4
	c := Complex{x: 0.0, y: 0.0}
	mag := 1.0
	limit := 1000

	img := DrawPix(w, h, c.x, c.y, mag, limit)
	assert.Equal(img[3*4+0], uint8(25))
	assert.Equal(img[3*4+1], uint8(7))
	assert.Equal(img[3*4+2], uint8(26))
	assert.Equal(img[3*4+3], uint8(255))
}

func TestEncodePNG(t *testing.T) {
	w := 6
	h := 4
	c := Complex{x: 0.0, y: 0.0}
	mag := 1.0
	limit := 1000
	img := DrawNRGBA(w, h, c.x, c.y, mag, limit)
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		t.Error(err)
	}
}
