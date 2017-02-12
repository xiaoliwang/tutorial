package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width int
	height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	color_func := func (x, y int) uint8 {
		return uint8(x ^ y)
	}
	v := color_func(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}