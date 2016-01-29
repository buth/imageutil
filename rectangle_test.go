package imageutil

import (
	"image"
	"sync"
	"testing"
	"time"
)

func TestConcurrentRP(t *testing.T) {
	var w sync.WaitGroup
	rp := ConcurrentRP(&w,
		func(rp image.Rectangle) {
			time.Sleep(time.Millisecond)
		},
	)

	// Start a timer, and process 100 one-millisecond tasks.
	now := time.Now()
	for i := 0; i < 100; i++ {
		rp(image.Rect(0, 0, 0, 0))
	}
	w.Wait()

	// Get the result of the timer, and check it against the fastest non-
	// concurrent execution time.
	if delta := time.Now().Sub(now); delta >= time.Millisecond*100 {
		t.Error("100 one-millisecond tasks took", delta)
	}
}

func TestPointsRP(t *testing.T) {
	oneByOne := image.Rect(0, 0, 1, 1)

	n := 0
	PointsRP(image.Pt(0, 0), 1, 1, func(image.Point) {
		n++
	})(oneByOne)

	if n != 1 {
		t.Error("point proccessor did not run exactly once for a 1x1 rectangle")
	}

	PointsRP(image.Pt(0, 0), 0, 0, func(image.Point) {
		t.Error("specifying zero horizontal and vertical stride did not result in a noop")
	})(oneByOne)

	PointsRP(image.Pt(0, 0), 0, 1, func(image.Point) {
		t.Error("specifying zero horizontal stride did not result in a noop")
	})(oneByOne)

	PointsRP(image.Pt(0, 0), 1, 0, func(image.Point) {
		t.Error("specifying zero vertical stride did not result in a noop")
	})(oneByOne)

	for _, pt := range []image.Point{image.Pt(1, 0), image.Pt(-1, 0), image.Pt(0, 1), image.Pt(0, -1)} {
		PointsRP(pt, 1, 1, func(pt image.Point) {
			t.Error("specifying an origin not in the rectangle did not result in a noop")
		})(oneByOne)
	}

}
