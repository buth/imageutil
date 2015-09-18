package imageutil

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"runtime"
	"testing"
)

type testRect struct {
	name string
	r    image.Rectangle
}

var (
	gomaxprocs  = runtime.GOMAXPROCS(-1)
	stripeRects = []testRect{
		{"zero", image.Rect(0, 0, 0, 0)},
		{"zeroHorizontal", image.Rect(0, 0, 1, 0)},
		{"zeroVertical", image.Rect(0, 0, 0, 1)},
		{"unit", image.Rect(0, 0, 1, 1)},
		{"unitEvenHorizontal", image.Rect(0, 0, 2*gomaxprocs, 1)},
		{"unitEvenVertical", image.Rect(0, 0, 1, 2*gomaxprocs)},
		{"unitOddHorizontal", image.Rect(0, 0, 2*gomaxprocs+1, 1)},
		{"unitOddVertical", image.Rect(0, 0, 1, 2*gomaxprocs+1)},
		{"zeroVertical", image.Rect(0, 0, 0, 1)},
	}
	bubbleRects = []testRect{
		{"unit", image.Rect(0, 0, 1, 1)},
		{"unitEvenHorizontal", image.Rect(0, 0, 2*gomaxprocs, 1)},
		{"unitEvenVertical", image.Rect(0, 0, 1, 2*gomaxprocs)},
		{"unitOddHorizontal", image.Rect(0, 0, 2*gomaxprocs+1, 1)},
		{"unitOddVertical", image.Rect(0, 0, 1, 2*gomaxprocs+1)},
	}
)

func isRed(testImage image.Image, t *testing.T) error {
	r := testImage.Bounds()
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			if testImage.At(x, y).(color.RGBA).R != 255 {
				return errors.New(fmt.Sprintf("wrong value at (%d, %d)", x, y))
			}
		}
	}
	return nil
}

func testStripe(r image.Rectangle, t *testing.T) error {

	testImage := image.NewRGBA(r)
	c := color.RGBA{R: 255}

	Stripe(testImage.Bounds(), func(x, y int) {
		testImage.Set(x, y, c)
	})

	return isRed(testImage, t)
}

func testBubble(r, b image.Rectangle, t *testing.T) error {

	testImage := image.NewRGBA(r)
	c := color.RGBA{R: 255}

	Bubble(testImage.Bounds(), b, func(x, y int) {
		testImage.Set(x, y, c)
	})

	return isRed(testImage, t)
}

func TestStripe(t *testing.T) {
	for _, r := range stripeRects {
		if err := testStripe(r.r, t); err != nil {
			t.Errorf("stripe %s: %s", r.name, err.Error())
		}
	}
}

func TestBubble(t *testing.T) {
	for _, a := range stripeRects {
		for _, b := range bubbleRects {
			if err := testBubble(a.r, b.r, t); err != nil {
				t.Errorf("bubble %s/%s: %s", a.name, b.name, err.Error())
			}
		}
	}
}
