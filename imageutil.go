package imageutil

import (
	"image"
	"image/color"
)

// ImageReader implements the same methods as the standard image interface.
type ImageReader interface {
	At(x, y int) color.Color
	Bounds() image.Rectangle
	ColorModel() color.Model
}

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

// Copy concurrently copies Color values from a source ImageReader to a
// destination ImageReadWriter.
func Copy(dst ImageReadWriter, src ImageReader) {
	if dst != src {
		QuickRP(
			AllPointsRP(
				func(pt image.Point) {
					dst.Set(pt.X, pt.Y, src.At(pt.X, pt.Y))
				},
			),
		)(dst.Bounds().Union(src.Bounds()))
	}
}

// ConvertToRGBA returns an *image.RGBA instance by asserting the given
// ImageReader has that type or, if it does not, using Copy to concurrently
// set the color.Color values of a new *image.RGBA instance with the same
// bounds.
func ConvertToRGBA(src ImageReader) *image.RGBA {
	if dst, ok := src.(*image.RGBA); ok {
		return dst
	}
	dst := image.NewRGBA(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToRGBA64 returns an *image.RGBA64 instance by asserting the given
// ImageReader has that type or, if it does not, using Copy to concurrently
// set the color.Color values of a new *image.RGBA64 instance with the same
// bounds.
func ConvertToRGBA64(src ImageReader) *image.RGBA64 {
	if dst, ok := src.(*image.RGBA64); ok {
		return dst
	}
	dst := image.NewRGBA64(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToNRGBA returns an *image.NRGBA instance by asserting the given
// ImageReader has that type or, if it does not, using Copy to concurrently
// set the color.Color values of a new *image.NRGBA instance with the same
// bounds.
func ConvertToNRGBA(src ImageReader) *image.NRGBA {
	if dst, ok := src.(*image.NRGBA); ok {
		return dst
	}
	dst := image.NewNRGBA(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToNRGBA64 returns an *image.NRGBA64 instance by asserting the given
// ImageReader has that type or, if it does not, using Copy to concurrently
// set the color.Color values of a new *image.NRGBA64 instance with the same
// bounds.
func ConvertToNRGBA64(src ImageReader) *image.NRGBA64 {
	if dst, ok := src.(*image.NRGBA64); ok {
		return dst
	}
	dst := image.NewNRGBA64(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToAlpha returns an *image.Alpha instance by asserting the given
// ImageReader has that type or, if it does not, using Copy to concurrently
// set the color.Color values of a new *image.Alpha instance with the same
// bounds.
func ConvertToAlpha(src ImageReader) *image.Alpha {
	if dst, ok := src.(*image.Alpha); ok {
		return dst
	}
	dst := image.NewAlpha(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToAlpha16 returns an *image.Alpha16 instance by asserting the given
// ImageReader has that type or, if it does not, using Copy to concurrently
// set the color.Color values of a new *image.Alpha16 instance with the same
// bounds.
func ConvertToAlpha16(src ImageReader) *image.Alpha16 {
	if dst, ok := src.(*image.Alpha16); ok {
		return dst
	}
	dst := image.NewAlpha16(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToGray returns an *image.Gray instance by asserting the given
// ImageReader has that type or, if it does not, using Copy to concurrently
// set the color.Color values of a new *image.Gray instance with the same
// bounds.
func ConvertToGray(src ImageReader) *image.Gray {
	if dst, ok := src.(*image.Gray); ok {
		return dst
	}
	dst := image.NewGray(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToGray16 returns an *image.Gray16 instance by asserting the given
// ImageReader has that type or, if it does not, using Copy to concurrently
// set the color.Color values of a new *image.Gray16 instance with the same
// bounds.
func ConvertToGray16(src ImageReader) *image.Gray16 {
	if dst, ok := src.(*image.Gray16); ok {
		return dst
	}
	dst := image.NewGray16(src.Bounds())
	Copy(dst, src)
	return dst
}
