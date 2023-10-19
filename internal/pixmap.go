package internal

import (
	"fmt"
	"math"
)

// Pixel map is a 2D matrix grouping of pixels
// Each map has an index based on its position in the game window
type Pixmap struct {
	pixels []*bool
	idx    int
	size   int
	status bool
}

func NewPixmap(num, size int, status bool) *Pixmap {
	return &Pixmap{
		pixels: make([]*bool, 0),
		idx:    num,
		size:   size,
		status: status,
	}
}
func (p *Pixmap) IsAlive() bool {
	return p.status
}
func (p *Pixmap) SetAlive(alive bool) {
	p.status = alive
}

func (p *Pixmap) FillPixmap() {
	for index := range p.pixels {
		*p.pixels[index] = p.status
	}
}

// Set the pointer of pixels in a pixel map to its calculated pixel in the game window
func (p *Pixmap) InitPixmap(mSize int, MapState []bool) {
	for i := 0; i < int(math.Pow(float64(p.size), 2)); i++ {
		index := calcPixelIndex(i, p.idx, p.size, mSize)
		p.pixels = append(p.pixels, &MapState[index])
	}
}

func (p Pixmap) String() string {
	return fmt.Sprintf("index: %v, status: %v", p.idx, p.status)
}

// Calculate a pixel in the pixel-map's index in the 1D game window
func calcPixelIndex(pixelNum, pixelMapNum, pixelMapSize, mainWindowSize int) int {
	row := int(math.Floor(float64(pixelMapNum/(mainWindowSize/pixelMapSize))))*pixelMapSize + int(math.Floor(float64(pixelNum/pixelMapSize)))
	column := (pixelMapNum%int(mainWindowSize/pixelMapSize))*pixelMapSize + (pixelNum % pixelMapSize)
	index := row*mainWindowSize + column
	return index
}
