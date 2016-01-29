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

	switch img.(type) {
	case *image.Alpha, *image.Alpha16:
		return img
	case *image.Gray:
		invertedImage = image.NewGray(img.Bounds())
		pp = func(pt image.Point) {
			c := img.At(pt.X, pt.Y).(color.Gray)
			invertedImage.Set(pt.X, pt.Y,
				color.Gray{
					Y: math.MaxInt8 - c.Y,
				},
			)
		}
	case *image.Gray16:
		invertedImage = image.NewGray16(img.Bounds())
		pp = func(pt image.Point) {
			c := img.At(pt.X, pt.Y).(color.Gray16)
			invertedImage.Set(pt.X, pt.Y,
				color.Gray16{
					Y: math.MaxInt16 - c.Y,
				},
			)
		}
	case *image.NRGBA:
		invertedImage = image.NewNRGBA(img.Bounds())
		pp = func(pt image.Point) {
			c := img.At(pt.X, pt.Y).(color.NRGBA)
			invertedImage.Set(pt.X, pt.Y,
				color.NRGBA{
					R: math.MaxUint8 - c.R,
					G: math.MaxUint8 - c.G,
					B: math.MaxUint8 - c.B,
					A: c.A,
				},
			)
		}
	case *image.NRGBA64:
		invertedImage = image.NewNRGBA64(img.Bounds())
		pp = func(pt image.Point) {
			c := img.At(pt.X, pt.Y).(color.NRGBA64)
			invertedImage.Set(pt.X, pt.Y,
				color.NRGBA64{
					R: math.MaxUint16 - c.R,
					G: math.MaxUint16 - c.G,
					B: math.MaxUint16 - c.B,
					A: c.A,
				},
			)
		}
	case *image.RGBA:
		invertedImage = image.NewNRGBA(img.Bounds())
		pp = func(pt image.Point) {
			c := img.At(pt.X, pt.Y).(color.RGBA)
			invertedImage.Set(pt.X, pt.Y,
				color.RGBA{
					R: c.A - c.R,
					G: c.A - c.G,
					B: c.A - c.B,
					A: c.A,
				},
			)
		}
	case *image.RGBA64:
		invertedImage = image.NewRGBA64(img.Bounds())
		pp = func(pt image.Point) {
			c := img.At(pt.X, pt.Y).(color.RGBA64)
			invertedImage.Set(pt.X, pt.Y,
				color.RGBA64{
					R: c.A - c.R,
					G: c.A - c.G,
					B: c.A - c.B,
					A: c.A,
				},
			)
		}
	}

	QuickRP(AllPointsRP(pp))(img.Bounds())
	return invertedImage
}

func EdgesGray16(radius, padding int, img ImageReader) *image.Gray16 {
	bounds := img.Bounds()
	edgeImage := image.NewGray16(bounds)
	if radius < 1 || padding < 0 {
		return edgeImage
	}

	// Compute the horizontal and vertical averages.
	hGA := image.NewGray16(bounds)
	vGA := image.NewGray16(bounds)
	QuickRP(
		AllPointsRP(
			func(pt image.Point) {
				hGA.Set(pt.X, pt.Y,
					AverageGray16(image.Rect(pt.X, pt.Y-padding, pt.X+radius, pt.Y+padding+1), img),
				)
				vGA.Set(pt.X, pt.Y,
					AverageGray16(image.Rect(pt.X-padding, pt.Y, pt.X+padding+1, pt.Y+radius), img),
				)
			},
		),
	)(bounds)

	QuickRP(
		AllPointsRP(
			func(pt image.Point) {
				e := float64(hGA.Gray16At(pt.X+1, pt.Y).Y)
				w := float64(hGA.Gray16At(pt.X-radius, pt.Y).Y)
				n := float64(vGA.Gray16At(pt.X, pt.Y+1).Y)
				s := float64(vGA.Gray16At(pt.X, pt.Y-radius).Y)
				edgeImage.Set(pt.X, pt.Y,
					color.Gray16{
						Y: uint16((math.Abs(e-w) + math.Abs(s-n)) / 2.0),
					},
				)
			},
		),
	)(bounds)

	return edgeImage
}

func EdgesNRGBA64(radius, padding int, img *image.NRGBA64) *image.NRGBA64 {
	r, g, b, a := NRGBA64ToChannels(img)
	r = EdgesGray16(radius, padding, r)
	g = EdgesGray16(radius, padding, g)
	b = EdgesGray16(radius, padding, b)
	return ChannelsToNRGBA64(r, g, b, a)
}
