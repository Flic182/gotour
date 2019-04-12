package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	Width, Height int
	col uint8
}

func (r Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}

func (r Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (r Image) At(x, y int) color.Color {
	return color.RGBA{r.col+uint8(x), r.col+uint8(y), 255, 255}
}

func main() {
	m := Image{200, 200, 5}
	pic.ShowImage(m)
}
