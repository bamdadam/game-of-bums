package internal

import (
	"fmt"
	"math"
)

type Kernel struct {
	pixels []*bool
	knum   int
	ksize  int
	status bool
}

func NewKernel(num, size int, status bool) *Kernel {
	return &Kernel{
		pixels: make([]*bool, 0),
		knum:   num,
		ksize:  size,
		status: status,
	}
}
func (k *Kernel) IsAlive() bool {
	return k.status
}
func (k *Kernel) SetAlive(alive bool) {
	k.status = alive
}

func (k *Kernel) FillKernel() {
	for index := range k.pixels {
		*k.pixels[index] = k.status
	}
}

func (k *Kernel) InitKernel(mSize int, MapState []bool) {
	for i := 0; i < int(math.Pow(float64(k.ksize), 2)); i++ {
		index := calcPixelIndex(i, k.knum, k.ksize, mSize)
		k.pixels = append(k.pixels, &MapState[index])
	}
}

func (k *Kernel) String() string {
	return fmt.Sprintf("index: %v, status: %v", k.knum, k.status)
}

func calcPixelIndex(pNum, kNum, kSize, mSize int) int {
	row := int(math.Floor(float64(kNum/(mSize/kSize))))*kSize + int(math.Floor(float64(pNum/kSize)))
	column := (kNum%int(mSize/kSize))*kSize + (pNum % kSize)
	index := row*mSize + column
	return index
}
