package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i *Image) At(x, y int) color.Color {
	v := x * x * y * y
	return color.RGBA{uint8(v), uint8(v), 144, 233}
}

func main() {
	m := &Image{150, 300}
	pic.ShowImage(m)
}
