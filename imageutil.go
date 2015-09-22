package imageutil

import (
	"image"
	"image/color"
)

// ImageReader implements the same methods as the standard image interface.
type ImageReader image.Image

// ImageWriter implements a method for setting colors at individual
// coordinates (which all standard image types implement).
type ImageWriter interface {
	Set(x, y int, c color.Color)
}

// ImageReadWriter implements both ImageReader and ImageWriter.
type ImageReadWriter interface {
	ImageReader
	ImageWriter
}

// Copy copies Color values from a source ImageReader instance to a
// destination ImageReadWriter instance concurrently (potentially in parallel)
// using Stripe.
func Copy(dst ImageReadWriter, src ImageReader) {
	if dst != src {
		commonBounds := dst.Bounds().Union(src.Bounds())
		Stripe(commonBounds, func(x, y int) {
			dst.Set(x, y, src.At(x, y))
		})
	}
}

// ConvertToRGBA returns a new copy of an ImageReader instance as a new RGBA
// instance using Copy to concurrently (potentially in parallel) set Color
// values.
func ConvertToRGBA(src ImageReader) *image.RGBA {
	dst := image.NewRGBA(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToRGBA64 returns a new copy of an ImageReader instance as a new
// RGBA64 instance using Copy to concurrently (potentially in parallel) set
// Color values.
func ConvertToRGBA64(src ImageReader) *image.RGBA64 {
	dst := image.NewRGBA64(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToNRGBA returns a new copy of an ImageReader instance as a new NRGBA
// instance using Copy to concurrently (potentially in parallel) set Color
// values.
func ConvertToNRGBA(src ImageReader) *image.NRGBA {
	dst := image.NewNRGBA(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToNRGBA64 returns a new copy of an ImageReader instance as a new
// NRGBA64 instance using Copy to concurrently (potentially in parallel) set
// Color values.
func ConvertToNRGBA64(src ImageReader) *image.NRGBA64 {
	dst := image.NewNRGBA64(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToAlpha returns a new copy of an ImageReader instance as a new Alpha
// instance using Copy to concurrently (potentially in parallel) set Color
// values.
func ConvertToAlpha(src ImageReader) *image.Alpha {
	dst := image.NewAlpha(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToAlpha16 returns a new copy of an ImageReader instance as a new
// Alpha16 instance using Copy to concurrently (potentially in parallel) set
// Color values.
func ConvertToAlpha16(src ImageReader) *image.Alpha16 {
	dst := image.NewAlpha16(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToGray returns a new copy of an ImageReader instance as a new Gray
// instance using Copy to concurrently (potentially in parallel) set Color
// values.
func ConvertToGray(src ImageReader) *image.Gray {
	dst := image.NewGray(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToGray16 returns a new copy of an ImageReader instance as a new
// Gray16 instance using Copy to concurrently (potentially in parallel) set
// Color values.
func ConvertToGray16(src ImageReader) *image.Gray16 {
	dst := image.NewGray16(src.Bounds())
	Copy(dst, src)
	return dst
}
