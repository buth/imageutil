package imageutil

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand"
	"testing"
)

var (
	imageutilTestRect = image.Rect(0, 0, 1000, 1000)
)

func randomNRGBA64(rect image.Rectangle) image.Image {
	img := image.NewNRGBA64(rect)
	QuickRP(
		AllPointsRP(func(pt image.Point) {
			img.SetNRGBA64(pt.X, pt.Y, color.NRGBA64{
				R: uint16(rand.Intn(int(math.MaxUint16 + 1))),
				G: uint16(rand.Intn(int(math.MaxUint16 + 1))),
				B: uint16(rand.Intn(int(math.MaxUint16 + 1))),
				A: uint16(rand.Intn(int(math.MaxUint16 + 1))),
			})
		}),
	)(rect)
	return img
}

func convertToGray16StdDraw(src ImageReader) *image.Gray16 {
	dst := image.NewGray16(src.Bounds())
	draw.Draw(dst, src.Bounds(), src, image.Pt(0, 0), draw.Over)
	return dst
}

func convertToGray16QuickDraw(src ImageReader) *image.Gray16 {
	dst := image.NewGray16(src.Bounds())
	QuickRP(
		func(rect image.Rectangle) {
			draw.Draw(dst, rect, src, rect.Min, draw.Over)
		},
	)(src.Bounds())
	return dst
}

func convertToGray16QuickAllPoints(src ImageReader) *image.Gray16 {
	dst := image.NewGray16(src.Bounds())
	QuickRP(
		AllPointsRP(
			func(pt image.Point) {
				dst.Set(pt.X, pt.Y, src.At(pt.X, pt.Y))
			},
		),
	)(src.Bounds())
	return dst
}

func testGray16(dst *image.Gray16, src image.Image, t *testing.T) {
	bounds := src.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if c := dst.Gray16At(x, y); c != color.Gray16Model.Convert(src.At(x, y)).(color.Gray16) {
				t.Fatalf("Unexpected color %v", c)
			}
		}
	}
}

func TestConvertToGray16StdDraw(t *testing.T) {
	src := randomNRGBA64(imageutilTestRect)
	dst := convertToGray16StdDraw(src)
	testGray16(dst, src, t)
}

func TestConvertToGray16QuickDraw(t *testing.T) {
	src := randomNRGBA64(imageutilTestRect)
	dst := convertToGray16QuickDraw(src)
	testGray16(dst, src, t)
}

func TestConvertToGray16QuickAllPoints(t *testing.T) {
	src := randomNRGBA64(imageutilTestRect)
	dst := convertToGray16QuickAllPoints(src)
	testGray16(dst, src, t)
}

func TestConvertToGray16(t *testing.T) {
	src := randomNRGBA64(imageutilTestRect)
	dst := ConvertToGray16(src)
	testGray16(dst, src, t)
}

func BenchmarkConvertToGray16StdDraw(b *testing.B) {
	src := randomNRGBA64(imageutilTestRect)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		convertToGray16StdDraw(src)
	}
}

func BenchmarkConvertToGray16QuickDraw(b *testing.B) {
	src := randomNRGBA64(imageutilTestRect)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		convertToGray16QuickDraw(src)
	}
}

func BenchmarkConvertToGray16QuickAllPoints(b *testing.B) {
	src := randomNRGBA64(imageutilTestRect)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		convertToGray16QuickAllPoints(src)
	}
}

func BenchmarkConvertToGray16(b *testing.B) {
	src := randomNRGBA64(imageutilTestRect)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConvertToGray16(src)
	}
}
