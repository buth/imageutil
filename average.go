package imageutil

import (
	"image"
	"image/color"
)

func AverageGray16(rect image.Rectangle, img ImageReader) color.Gray16 {
	d := uint64(rect.Dx() * rect.Dy())
	if d == 0 {
		return color.Gray16{}
	}

	var y uint64
	AllPointsRP(
		func(pt image.Point) {
			c := color.Gray16Model.Convert(img.At(pt.X, pt.Y)).(color.Gray16)
			y += uint64(c.Y)
		},
	)(rect)

	return color.Gray16{
		Y: uint16(y / d),
	}
}

func AverageNRGBA64(rect image.Rectangle, img ImageReader) color.NRGBA64 {
	d := uint64(rect.Dx() * rect.Dy())
	if d == 0 {
		return color.NRGBA64{}
	}

	var r, g, b, a uint64
	AllPointsRP(
		func(pt image.Point) {
			c := color.RGBA64Model.Convert(img.At(pt.X, pt.Y)).(color.NRGBA64)
			r += uint64(c.R)
			g += uint64(c.G)
			b += uint64(c.B)
			a += uint64(c.A)
		},
	)(rect)

	return color.NRGBA64{
		R: uint16(r / d),
		G: uint16(g / d),
		B: uint16(b / d),
		A: uint16(a / d),
	}
}
