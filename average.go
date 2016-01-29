package imageutil

import (
	"image"
	"image/color"
)

func AverageGray16(rect image.Rectangle, img *image.Gray16) color.Gray16 {

	// Only use the area of the rectangle that overlaps with the image bounds.
	rect = rect.Union(img.Bounds())

	// Determine whether or not there's any area over which to determine an
	// average.
	d := uint64(rect.Dx() * rect.Dy())
	if d == 0 {
		return color.Gray16{}
	}

	var y uint64
	AllPointsRP(
		func(pt image.Point) {
			y += uint64(img.Gray16At(pt.X, pt.Y).Y)
		},
	)(rect)

	return color.Gray16{
		Y: uint16(y / d),
	}
}

func AverageNRGBA64(rect image.Rectangle, img *image.NRGBA64) color.NRGBA64 {

	// Only use the area of the rectangle that overlaps with the image bounds.
	rect = rect.Union(img.Bounds())

	// Determine whether or not there's any area over which to determine an
	// average.
	d := uint64(rect.Dx() * rect.Dy())
	if d == 0 {
		return color.NRGBA64{}
	}

	var r, g, b, a uint64
	AllPointsRP(
		func(pt image.Point) {
			c := img.NRGBA64At(pt.X, pt.Y)
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
