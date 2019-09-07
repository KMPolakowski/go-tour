package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 128, 87)
}

func (img Image) At(x, y int) color.Color {
	code := uint8(x ^ y)
	return color.RGBA{code, code, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
