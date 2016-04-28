package imageutil

import (
	"image"
	"image/color"
)

// Channel is any object that implements ImageReader as well as providing a
// method for getting color.Gray16 values at given coordinates. The standard
// library's image.Gray16 implements Channel.
type Channel interface {
	ImageReader
	Gray16At(x, y int) color.Gray16
}

// channel is an implementation of Channel.
type channel struct {
	bounds   func() image.Rectangle
	gray16At func(x, y int) color.Gray16
}

// Channels decomposes a given NRGBA64 into red, green, blue, and alpha Channels.
func NRGBA64ToChannels(img *image.NRGBA64) (r, g, b, a Channel) {
	r = channel{
		bounds: img.Bounds,
		gray16At: func(x, y int) color.Gray16 {
			return color.Gray16{
				Y: img.NRGBA64At(x, y).R,
			}
		},
	}

	g = channel{
		bounds: img.Bounds,
		gray16At: func(x, y int) color.Gray16 {
			return color.Gray16{
				Y: img.NRGBA64At(x, y).G,
			}
		},
	}

	b = channel{
		bounds: img.Bounds,
		gray16At: func(x, y int) color.Gray16 {
			return color.Gray16{
				Y: img.NRGBA64At(x, y).B,
			}
		},
	}

	a = channel{
		bounds: img.Bounds,
		gray16At: func(x, y int) color.Gray16 {
			return color.Gray16{
				Y: img.NRGBA64At(x, y).A,
			}
		},
	}

	return
}

func (c channel) Bounds() image.Rectangle {
	return c.bounds()
}

func (c channel) At(x, y int) color.Color {
	return c.gray16At(x, y)
}

func (c channel) Gray16At(x, y int) color.Gray16 {
	return c.gray16At(x, y)
}

func (c channel) ColorModel() color.Model {
	return color.Gray16Model
}

func ChannelsToNRGBA64(r, g, b, a Channel) *image.NRGBA64 {
	bounds := r.Bounds().Union(g.Bounds()).Union(b.Bounds()).Union(a.Bounds())
	img := image.NewNRGBA64(bounds)
	QuickRP(
		AllPointsRP(
			func(pt image.Point) {
				img.Set(pt.X, pt.Y,
					color.NRGBA64{
						R: r.Gray16At(pt.X, pt.Y).Y,
						G: g.Gray16At(pt.X, pt.Y).Y,
						B: b.Gray16At(pt.X, pt.Y).Y,
						A: a.Gray16At(pt.X, pt.Y).Y,
					},
				)
			},
		),
	)(bounds)

	return img
}
