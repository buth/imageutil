package imageutil

import (
	"image"
	"image/color"
)

// ImageReader is the standard image interface.
type ImageReader image.Image

// ImageWriter is the standard image interface with the addition of a method
// for setting values (which all standard image types implement).
type ImageWriter interface {
	Set(x, y int, c color.Color)
}

type ImageReadWriter interface {
	ImageReader
	ImageWriter
}

// Copy copies values from a source image instance to a destination image
// writer instance.
func Copy(dst ImageReadWriter, src image.Image) {
	if dst != src {
		commonBounds := dst.Bounds().Union(src.Bounds())
		Stripe(commonBounds, func(x, y int) {
			dst.Set(x, y, src.At(x, y))
		})
	}
}

func ConvertToRGBA(src image.Image) *image.RGBA {
	dst := image.NewRGBA(src.Bounds())
	Copy(dst, src)
	return dst
}

func ConvertToRGBA64(src image.Image) *image.RGBA64 {
	dst := image.NewRGBA64(src.Bounds())
	Copy(dst, src)
	return dst
}

func ConvertToNRGBA(src image.Image) *image.NRGBA {
	dst := image.NewNRGBA(src.Bounds())
	Copy(dst, src)
	return dst
}

func ConvertToNRGBA64(src image.Image) *image.NRGBA64 {
	dst := image.NewNRGBA64(src.Bounds())
	Copy(dst, src)
	return dst
}

func ConvertToAlpha(src image.Image) *image.Alpha {
	dst := image.NewAlpha(src.Bounds())
	Copy(dst, src)
	return dst
}

func ConvertToAlpha16(src image.Image) *image.Alpha16 {
	dst := image.NewAlpha16(src.Bounds())
	Copy(dst, src)
	return dst
}

func ConvertToGray(src image.Image) *image.Gray {
	dst := image.NewGray(src.Bounds())
	Copy(dst, src)
	return dst
}

func ConvertToGray16(src image.Image) *image.Gray16 {
	dst := image.NewGray16(src.Bounds())
	Copy(dst, src)
	return dst
}
