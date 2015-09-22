package imageutil

import (
	"image"
	"image/color"
	"testing"
)

type testImage struct {
	image ImageReader
	color color.Color
}

func testImageReaders() []ImageReader {
	testPairs := []struct {
		image ImageReadWriter
		color color.Color
	}{
		{image.NewRGBA(image.Rect(0, 0, 1, 1)), color.RGBA{R: 10, G: 20, B: 30, A: 40}},
		{image.NewRGBA64(image.Rect(0, 0, 1, 1)), color.RGBA64{R: 10, G: 20, B: 30, A: 40}},
		{image.NewNRGBA(image.Rect(0, 0, 1, 1)), color.NRGBA{R: 10, G: 20, B: 30, A: 40}},
		{image.NewNRGBA64(image.Rect(0, 0, 1, 1)), color.NRGBA64{R: 10, G: 20, B: 30, A: 40}},
		{image.NewAlpha(image.Rect(0, 0, 1, 1)), color.Alpha{A: 40}},
		{image.NewAlpha16(image.Rect(0, 0, 1, 1)), color.Alpha16{A: 40}},
		{image.NewGray(image.Rect(0, 0, 1, 1)), color.Gray{Y: 40}},
		{image.NewGray16(image.Rect(0, 0, 1, 1)), color.Gray16{Y: 40}},
	}

	imageReaders := make([]ImageReader, 0, len(testPairs))
	for _, testPair := range testPairs {
		testPair.image.Set(0, 0, testPair.color)
		imageReaders = append(imageReaders, testPair.image)
	}

	return imageReaders
}

func TestConvertToRGBA(t *testing.T) {
	for _, imageReader := range testImageReaders() {
		c := imageReader.At(0, 0)
		expectedC := color.RGBAModel.Convert(c)
		outputImage := ConvertToRGBA(imageReader)
		if outputImage.At(0, 0) != expectedC {
			t.Fatalf("unexpected color")
		}
	}
}

func TestConvertToRGBA64(t *testing.T) {
	for _, imageReader := range testImageReaders() {
		c := imageReader.At(0, 0)
		expectedC := color.RGBA64Model.Convert(c)
		outputImage := ConvertToRGBA64(imageReader)
		if outputImage.At(0, 0) != expectedC {
			t.Fatalf("unexpected color")
		}
	}
}

func TestConvertToNRGBA(t *testing.T) {
	for _, imageReader := range testImageReaders() {
		c := imageReader.At(0, 0)
		expectedC := color.NRGBAModel.Convert(c)
		outputImage := ConvertToNRGBA(imageReader)
		if outputImage.At(0, 0) != expectedC {
			t.Fatalf("unexpected color")
		}
	}
}

func TestConvertToNRGBA64(t *testing.T) {
	for _, imageReader := range testImageReaders() {
		c := imageReader.At(0, 0)
		expectedC := color.NRGBA64Model.Convert(c)
		outputImage := ConvertToNRGBA64(imageReader)
		if outputImage.At(0, 0) != expectedC {
			t.Fatalf("unexpected color")
		}
	}
}

func TestConvertToAlpha(t *testing.T) {
	for _, imageReader := range testImageReaders() {
		c := imageReader.At(0, 0)
		expectedC := color.AlphaModel.Convert(c)
		outputImage := ConvertToAlpha(imageReader)
		if outputImage.At(0, 0) != expectedC {
			t.Fatalf("unexpected color")
		}
	}
}

func TestConvertToAlpha16(t *testing.T) {
	for _, imageReader := range testImageReaders() {
		c := imageReader.At(0, 0)
		expectedC := color.Alpha16Model.Convert(c)
		outputImage := ConvertToAlpha16(imageReader)
		if outputImage.At(0, 0) != expectedC {
			t.Fatalf("unexpected color")
		}
	}
}

func TestConvertToGray(t *testing.T) {
	for _, imageReader := range testImageReaders() {
		c := imageReader.At(0, 0)
		expectedC := color.GrayModel.Convert(c)
		outputImage := ConvertToGray(imageReader)
		if outputImage.At(0, 0) != expectedC {
			t.Fatalf("unexpected color")
		}
	}
}

func TestConvertToGray16(t *testing.T) {
	for _, imageReader := range testImageReaders() {
		c := imageReader.At(0, 0)
		expectedC := color.Gray16Model.Convert(c)
		outputImage := ConvertToGray16(imageReader)
		if outputImage.At(0, 0) != expectedC {
			t.Fatalf("unexpected color")
		}
	}
}
