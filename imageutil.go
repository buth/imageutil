package imageutil

import (
	"image"
	"image/color"
)

// ImageReader is the standard image interface.
type ImageReader image.Image

// ImageWriter implements a method for setting colors at individual
// coordinates (which all standard image types implement).
type ImageWriter interface {
	Set(x, y int, c color.Color)
}

// ImageReadWriter implements both ImageReader and ImageWriter
type ImageReadWriter interface {
	ImageReader
	ImageWriter
}

// Copy copies values from a source image instance to a destination image
// writer instance.
func Copy(dst ImageReadWriter, src ImageReader) {
	if dst != src {
		commonBounds := dst.Bounds().Union(src.Bounds())
		Stripe(commonBounds, func(x, y int) {
			dst.Set(x, y, src.At(x, y))
		})
	}
}

func ConvertToRGBA(src ImageReader) *image.RGBA {
	dst := image.NewRGBA(src.Bounds())
	Copy(dst, src)
	return dst
}

func ConvertToRGBA64(src ImageReader) *image.RGBA64 {
	dst := image.NewRGBA64(src.Bounds())
	Copy(dst, src)
	return dst
}

func ConvertToNRGBA(src ImageReader) *image.NRGBA {
	dst := image.NewNRGBA(src.Bounds())
	Copy(dst, src)
	return dst
}

func ConvertToNRGBA64(src ImageReader) *image.NRGBA64 {
	dst := image.NewNRGBA64(src.Bounds())
	Copy(dst, src)
	return dst
}

func ConvertToAlpha(src ImageReader) *image.Alpha {
	dst := image.NewAlpha(src.Bounds())
	Copy(dst, src)
	return dst
}

func ConvertToAlpha16(src ImageReader) *image.Alpha16 {
	dst := image.NewAlpha16(src.Bounds())
	Copy(dst, src)
	return dst
}

func ConvertToGray(src ImageReader) *image.Gray {
	dst := image.NewGray(src.Bounds())
	Copy(dst, src)
	return dst
}

func ConvertToGray16(src ImageReader) *image.Gray16 {
	dst := image.NewGray16(src.Bounds())
	Copy(dst, src)
	return dst
}
