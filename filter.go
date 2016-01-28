package imageutil

import (
	"image"
	"image/color"
	"math"
)

func EdgesGray16(radius, padding int, img Channel) *image.Gray16 {
	bounds := img.Bounds()
	edgeImage := image.NewGray16(bounds)
	if radius < 1 || padding < 0 {
		return edgeImage
	}

	// Compute the horizontal and vertical averages.
	averagesBounds := image.Rect(bounds.Min.X-radius, bounds.Min.Y-radius, bounds.Max.X, bounds.Max.Y)
	hGA := image.NewGray16(averagesBounds)
	vGA := image.NewGray16(averagesBounds)
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
	)(averagesBounds)

	QuickRP(
		AllPointsRP(
			func(pt image.Point) {
				e := float64(hGA.Gray16At(pt.X+1, pt.Y).Y)
				w := float64(hGA.Gray16At(pt.X-radius, pt.Y).Y)
				n := float64(hGA.Gray16At(pt.X, pt.Y+1).Y)
				s := float64(hGA.Gray16At(pt.X, pt.Y-radius).Y)
				edgeImage.Set(pt.X, pt.Y,
					color.Gray16{
						Y: math.MaxUint16 - uint16((math.Abs(e-w)+math.Abs(s-n))/2.0),
					},
				)
			},
		),
	)(bounds)

	return edgeImage
}

func EdgesRGBA64(radius, padding int, img ImageReader) *image.RGBA64 {
	r, g, b, a := Channels(img)
	r = EdgesGray16(radius, padding, r)
	g = EdgesGray16(radius, padding, g)
	b = EdgesGray16(radius, padding, b)
	a = EdgesGray16(radius, padding, a)

	return ChannelsToRGBA64(r, g, b, a)
}
