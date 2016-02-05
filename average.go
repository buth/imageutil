package imageutil

import (
	"image"
	"image/color"
)

func RowAverageGray16(radius int, img Channel) *image.Gray16 {
	bounds := img.Bounds()
	resultBounds := image.Rect(bounds.Min.X-radius+1, bounds.Min.Y, bounds.Max.X, bounds.Max.Y)
	resultImg := image.NewGray16(resultBounds)

	QuickRowsRP(
		RowsRP(1, func(rect image.Rectangle) {
			y := rect.Min.Y
			n := 0
			d := 0

			// Heads.
			x := resultBounds.Min.X
			for ; x <= bounds.Min.X; x++ {
				n += int(img.Gray16At(x+radius-1, y).Y)
				d++

				resultImg.Set(x, y, color.Gray16{
					Y: uint16(n / d),
				})
			}

			// Middle.
			for ; x <= bounds.Max.X-radius; x++ {
				n += int(img.Gray16At(x+radius-1, y).Y)
				n -= int(img.Gray16At(x-1, y).Y)

				resultImg.Set(x, y, color.Gray16{
					Y: uint16(n / d),
				})
			}

			// Tails.
			for ; x < bounds.Max.X; x++ {
				n -= int(img.Gray16At(x-1, y).Y)
				d--

				resultImg.Set(x, y, color.Gray16{
					Y: uint16(n / d),
				})
			}
		}),
	)(bounds)

	return resultImg
}

func ColumnAverageGray16(radius int, img Channel) *image.Gray16 {
	bounds := img.Bounds()
	resultBounds := image.Rect(bounds.Min.X, bounds.Min.Y-radius+1, bounds.Max.X, bounds.Max.Y)
	resultImg := image.NewGray16(resultBounds)

	QuickColumnsRP(
		ColumnsRP(1, func(rect image.Rectangle) {
			x := rect.Min.X
			n := 0
			d := 0

			// Heads.
			y := resultBounds.Min.Y
			for ; y <= bounds.Min.Y; y++ {
				n += int(img.Gray16At(x, y+radius-1).Y)
				d++

				resultImg.Set(x, y, color.Gray16{
					Y: uint16(n / d),
				})
			}

			// Middle.
			for ; y <= bounds.Max.Y-radius; y++ {
				n += int(img.Gray16At(x, y+radius-1).Y)
				n -= int(img.Gray16At(x, y-1).Y)

				resultImg.Set(x, y, color.Gray16{
					Y: uint16(n / d),
				})
			}

			// Tails.
			for ; y < bounds.Max.Y; y++ {
				n -= int(img.Gray16At(x, y-1).Y)
				d--

				resultImg.Set(x, y, color.Gray16{
					Y: uint16(n / d),
				})
			}
		}),
	)(bounds)

	return resultImg
}

func AverageGray16(rect image.Rectangle, img Channel) color.Gray16 {

	// Only use the area of the rectangle that overlaps with the image bounds.
	rect = rect.Intersect(img.Bounds())

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
	rect = rect.Intersect(img.Bounds())

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
