package imageutil

import (
	"image"
	"image/color"
	"runtime"
	"sync"
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
		var w sync.WaitGroup
		NRectanglesRP(runtime.GOMAXPROCS(-1),
			ConcurrentRP(&w,
				AllPointsRP(
					func(pt image.Point) {
						dst.Set(pt.X, pt.Y, src.At(pt.X, pt.Y))
					},
				),
			),
		)(dst.Bounds().Union(src.Bounds()))
		w.Wait()
	}
}

// ConvertToRGBA returns a new copy of an ImageReader as a new RGBA using Copy
// to concurrently set Color values.
func ConvertToRGBA(src ImageReader) *image.RGBA {
	dst := image.NewRGBA(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToRGBA64 returns a new copy of an ImageReader as a new RGBA64 using
// Copy to concurrently set Color values.
func ConvertToRGBA64(src ImageReader) *image.RGBA64 {
	dst := image.NewRGBA64(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToNRGBA returns a new copy of an ImageReader as a new NRGBA using
// Copy to concurrently set Color values.
func ConvertToNRGBA(src ImageReader) *image.NRGBA {
	dst := image.NewNRGBA(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToNRGBA64 returns a new copy of an ImageReader as a new NRGBA64
// using Copy to concurrently set Color values.
func ConvertToNRGBA64(src ImageReader) *image.NRGBA64 {
	dst := image.NewNRGBA64(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToAlpha returns a new copy of an ImageReader as a new Alpha using
// Copy to concurrently set Color values.
func ConvertToAlpha(src ImageReader) *image.Alpha {
	dst := image.NewAlpha(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToAlpha16 returns a new copy of an ImageReader as a new Alpha16
// using Copy to concurrently set Color values.
func ConvertToAlpha16(src ImageReader) *image.Alpha16 {
	dst := image.NewAlpha16(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToGray returns a new copy of an ImageReader as a new Gray using Copy
// to concurrently set Color values.
func ConvertToGray(src ImageReader) *image.Gray {
	dst := image.NewGray(src.Bounds())
	Copy(dst, src)
	return dst
}

// ConvertToGray16 returns a new copy of an ImageReader as a new Gray16 using
// Copy to concurrently set Color values.
func ConvertToGray16(src ImageReader) *image.Gray16 {
	dst := image.NewGray16(src.Bounds())
	Copy(dst, src)
	return dst
}
