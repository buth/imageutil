package imageutil

import (
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

func TestOpenImage(t *testing.T) {

	testImage := image.NewNRGBA(image.Rect(0, 0, 1, 1))

	testColor := color.NRGBA{1, 2, 3, 4}
	testImage.Set(0, 0, testColor)

	testImageFile, err := ioutil.TempFile("", "testOpenImage.png")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := os.Remove(testImageFile.Name()); err != nil {
			t.Error(err)
		}
	}()

	if err := png.Encode(testImageFile, testImage); err != nil {
		t.Fatal(err)
	}

	if err := testImageFile.Close(); err != nil {
		t.Fatal(err)
	}

	readImage, format, err := OpenImage(testImageFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if format != "png" {
		t.Error("image had the wrong format:", format)
	}

	if bounds := readImage.Bounds(); !bounds.Eq(testImage.Bounds()) {
		t.Error("image had the wrong bounds:", readImage.Bounds())
	}

	if readColor := color.NRGBAModel.Convert(readImage.At(0, 0)); readColor != testColor {
		t.Error("image had the wrong color at (0, 0):", readColor)
	}
}
