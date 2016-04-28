package imageutil

import (
	"image"
	"image/color"
	"testing"
)

func TestRowAverageGray16(t *testing.T) {
	testImage := image.NewGray16(image.Rect(0, 0, 1024, 1))
	for i := 0; i < 1024; i++ {
		testImage.Set(i, 0, color.Gray16{Y: uint16(i)})
	}

	resultImage := RowAverageGray16(10, testImage)

	for i := 0; i < 1024; i++ {
		n := 0
		d := 0
		for j := 0; (j < 10) && (i+j < 1024); j++ {
			n += i + j
			d++
		}

		expected := uint16(n / d)
		if y := resultImage.Gray16At(i, 0).Y; expected != y {
			t.Errorf("Expected %d, found %d at x=%d", expected, y, i)
		}
	}
}

func TestColumnAverageGray16(t *testing.T) {
	testImage := image.NewGray16(image.Rect(0, 0, 1, 1024))
	for i := 0; i < 1024; i++ {
		testImage.Set(0, i, color.Gray16{Y: uint16(i)})
	}

	resultImage := ColumnAverageGray16(10, testImage)

	for i := 0; i < 1024; i++ {
		n := 0
		d := 0
		for j := 0; (j < 10) && (i+j < 1024); j++ {
			n += i + j
			d++
		}

		expected := uint16(n / d)
		if y := resultImage.Gray16At(0, i).Y; expected != y {
			t.Errorf("Expected %d, found %d at x=%d", expected, y, i)
		}
	}
}
