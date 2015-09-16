package imageutil

import (
	"image"
	"runtime"
	"sync"
)

func processStripe(bounds image.Rectangle, processor func(x, y int)) {
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			processor(x, y)
		}
	}
}

// Stripe
func Stripe(bounds image.Rectangle, processor func(x, y int)) {

	// Set the original starting minimum.
	min := bounds.Min

	// We'll process the last stripe in this Go routine, so we only need a wait
	// group for the remainder.
	var w sync.WaitGroup

	// Get the current value of GOMAXPROCS. Theoretically this could be updated
	// while this code is running, but our priority is using a single,
	// consistent value.
	if gomaxprocs := runtime.GOMAXPROCS(-1); gomaxprocs > 1 {

		// Establish the total size.
		totalSize := bounds.Size()

		// Set the stripe delta and stripe size depending on which dimension is
		// greater.
		var delta, stripeSize image.Point
		if totalSize.X > totalSize.Y {
			delta.X = totalSize.X / gomaxprocs
			stripeSize.X = delta.X
			stripeSize.Y = totalSize.Y
		} else {
			stripeSize.X = totalSize.X
			delta.Y = totalSize.Y / gomaxprocs
			stripeSize.Y = delta.Y
		}

		// Check to make sure that stripe size is non-zero.
		if stripeSize.X != 0 && stripeSize.Y != 0 {
			w.Add(gomaxprocs - 1)

			// Process all of the rounded stripes
			for i := 0; i < gomaxprocs-1; i++ {
				max := min.Add(stripeSize)
				stripe := image.Rect(min.X, min.Y, max.X, max.Y)
				go func() {
					processStripe(stripe, processor)
					w.Done()
				}()

				min = min.Add(delta)
			}
		}
	}

	// Process the last stripe – or the entire rect if the stripe size was zero.
	stripe := image.Rect(min.X, min.Y, bounds.Max.X, bounds.Max.Y)
	processStripe(stripe, processor)

	w.Wait()
}
