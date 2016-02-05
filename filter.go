package imageutil

import (
	"image"
	"image/color"
	"math"
)

func Invert(img ImageReader) ImageReader {
	var (
		invertedImage ImageReadWriter
		pp            PP
	)

	bounds := img.Bounds()
	switch img.(type) {
	case *image.Alpha, *image.Alpha16:
		return img
	case *image.Gray:
		invertedImage = image.NewGray(bounds)
		pp = func(pt image.Point) {
			c := img.At(pt.X, pt.Y).(color.Gray)
			c.Y = math.MaxInt8 - c.Y
			invertedImage.Set(pt.X, pt.Y, c)
		}
	case *image.Gray16:
		invertedImage = image.NewGray16(bounds)
		pp = func(pt image.Point) {
			c := img.At(pt.X, pt.Y).(color.Gray16)
			c.Y = math.MaxInt8 - c.Y
			invertedImage.Set(pt.X, pt.Y, c)
		}
	case *image.NRGBA:
		invertedImage = image.NewNRGBA(bounds)
		pp = func(pt image.Point) {
			c := img.At(pt.X, pt.Y).(color.NRGBA)
			c.R = math.MaxUint8 - c.R
			c.G = math.MaxUint8 - c.G
			c.B = math.MaxUint8 - c.B
			invertedImage.Set(pt.X, pt.Y, c)
		}
	case *image.NRGBA64:
		invertedImage = image.NewNRGBA64(bounds)
		pp = func(pt image.Point) {
			c := img.At(pt.X, pt.Y).(color.NRGBA64)
			c.R = math.MaxUint16 - c.R
			c.G = math.MaxUint16 - c.G
			c.B = math.MaxUint16 - c.B
			invertedImage.Set(pt.X, pt.Y, c)
		}
	case *image.RGBA:
		invertedImage = image.NewNRGBA(bounds)
		pp = func(pt image.Point) {
			c := img.At(pt.X, pt.Y).(color.RGBA)
			c.R = c.A - c.R
			c.G = c.A - c.G
			c.B = c.A - c.B
			invertedImage.Set(pt.X, pt.Y, c)
		}
	case *image.RGBA64:
		invertedImage = image.NewRGBA64(bounds)
		pp = func(pt image.Point) {
			c := img.At(pt.X, pt.Y).(color.RGBA64)
			c.R = c.A - c.R
			c.G = c.A - c.G
			c.B = c.A - c.B
			invertedImage.Set(pt.X, pt.Y, c)
		}
	}

	QuickRP(AllPointsRP(pp))(bounds)
	return invertedImage
}

func EdgesGray16(radius int, img Channel) *image.Gray16 {
	bounds := img.Bounds()
	edgeImage := image.NewGray16(bounds)
	if radius < 1 {
		return edgeImage
	}

	// Compute the horizontal and vertical averages.
	hGA := RowAverageGray16(radius, img)
	vGA := ColumnAverageGray16(radius, img)

	QuickRP(
		AllPointsRP(
			func(pt image.Point) {
				e := float64(hGA.Gray16At(pt.X, pt.Y).Y)
				w := float64(hGA.Gray16At(pt.X-radius+1, pt.Y).Y)
				n := float64(vGA.Gray16At(pt.X, pt.Y).Y)
				s := float64(vGA.Gray16At(pt.X, pt.Y-radius+1).Y)
				edgeImage.Set(pt.X, pt.Y,
					color.Gray16{
						Y: uint16(math.Max(math.Abs(e-w), math.Abs(s-n))), //uint16((math.Abs(e-w) + math.Abs(s-n)) / 2.0),
					},
				)
			},
		),
	)(bounds)

	return edgeImage
}

func EdgesNRGBA64(radius int, img *image.NRGBA64) *image.NRGBA64 {
	r, g, b, a := NRGBA64ToChannels(img)
	r = EdgesGray16(radius, r)
	g = EdgesGray16(radius, g)
	b = EdgesGray16(radius, b)
	return ChannelsToNRGBA64(r, g, b, a)
}
