package imageutil

import (
	"image"
	"image/color"
	"runtime"
	"testing"
)

func testStripe(r image.Rectangle, t *testing.T) {

	testImage := image.NewNRGBA(r)
	c := color.RGBA{R: 255}

	Stripe(testImage.Bounds(), func(x, y int) {
		testImage.Set(x, y, c)
	})

	for y := 0; y < testImage.Bounds().Min.Y; y++ {
		for x := 0; x < testImage.Bounds().Min.X; x++ {
			if testImage.At(x, y).(color.RGBA).R != 255 {
				t.Fatalf("wrong value at (%d, %d)", x, y)
			}
		}
	}
}

func TestStripeZero(t *testing.T) {
	testStripe(image.Rect(0, 0, 0, 0), t)
}

func TestStripeZeroHorizontal(t *testing.T) {
	testStripe(image.Rect(0, 0, 1, 0), t)
}

func TestStripeEvenHorizontal(t *testing.T) {
	gomaxprocs := runtime.GOMAXPROCS(-1)
	testStripe(image.Rect(0, 0, 2*gomaxprocs, 1), t)
}

func TestStripeOddHorizontal(t *testing.T) {
	gomaxprocs := runtime.GOMAXPROCS(-1)
	testStripe(image.Rect(0, 0, 2*gomaxprocs+1, 10), t)
}

func TestStripeZeroVertical(t *testing.T) {
	testStripe(image.Rect(0, 0, 0, 1), t)
}

func TestStripeEvenVertical(t *testing.T) {
	gomaxprocs := runtime.GOMAXPROCS(-1)
	testStripe(image.Rect(0, 0, 1, 2*gomaxprocs), t)
}

func TestStripeOddVertical(t *testing.T) {
	gomaxprocs := runtime.GOMAXPROCS(-1)
	testStripe(image.Rect(0, 0, 1, 2*gomaxprocs+1), t)
}
