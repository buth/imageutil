package imageutil

import (
	"image"
	"runtime"
	"sync"
)

// RP is a rectangle processor, any function that accepts a single
// image.Rectangle value as an argument.
type RP func(image.Rectangle)

// PP is a point processor, any function that accepts a single image.Point
// value as an argument.
type PP func(image.Point)

// noopRP is an RP that does nothing. It's defined as a variable to allow for
// compairability.
var (
	noopRP = func(image.Rectangle) {}
)

// ConcurrentRP wraps a given RP in a Go routine and the necessary WaitGroup
// operations to ensure that completion can be tracked.
func ConcurrentRP(w *sync.WaitGroup, rp RP) RP {
	return func(rect image.Rectangle) {
		w.Add(1)
		go func() {
			rp(rect)
			w.Done()
		}()
	}
}

// PointsRP returns a RP that runs a given PP at each point within an input
// rectangle starting at a given offset and seperated by the
// given horizontal and vertical stride.
func PointsRP(offset image.Point, strideH, strideV int, pp PP) RP {
	if strideH <= 0 || strideV <= 0 || offset.X < 0 || offset.Y < 0 {
		return noopRP
	}

	return func(rect image.Rectangle) {
		origin := rect.Min.Add(offset)
		for y := origin.Y; y < rect.Max.Y; y += strideH {
			for x := origin.X; x < rect.Max.X; x += strideV {
				pp(image.Pt(x, y))
			}
		}
	}
}

// AllPointsRP returns a RP that runs a given PP at every point within an
// input rectangle.
func AllPointsRP(pp PP) RP {
	return PointsRP(image.Pt(0, 0), 1, 1, pp)
}

// RowsRP returns a RP that devides an input rectangle into rows of a given
// hight and calls the provided RP on each. The last row proccessed will be
// any remainder and may not be of the given height.
func RowsRP(height int, rp RP) RP {
	if height <= 0 {
		return noopRP
	}

	return func(rect image.Rectangle) {

		// Process all but the last row.
		y := rect.Min.Y
		for ; y < rect.Max.Y-height; y += height {
			rp(image.Rect(rect.Min.X, y, rect.Max.X, y+height))
		}

		// Process the last row.
		rp(image.Rect(rect.Min.X, y, rect.Max.X, rect.Max.Y))
	}
}

// ColumnsRP returns a RP that devides an input rectangle into columns of a
// given width and calls the provided RP on each. The last column proccessed
// will be any remainder and may not be of the given width.
func ColumnsRP(width int, rp RP) RP {

	// If the specified column width is zero, there's nothing to do.
	if width <= 0 {
		return noopRP
	}

	return func(rect image.Rectangle) {

		// Process all but the last column.
		x := rect.Min.X
		for ; x < rect.Max.X-width; x += width {
			rp(image.Rect(x, rect.Min.Y, x+width, rect.Max.Y))
		}

		// Process the last column.
		rp(image.Rect(x, rect.Min.Y, rect.Max.X, rect.Max.Y))
	}
}

// NRowsRP calls the given RP on each of n horizontal rectangles that span the
// input rectangle.
func NRowsRP(n int, rp RP) RP {
	if n <= 0 {
		return noopRP
	}

	return func(rect image.Rectangle) {
		height := (rect.Dy() + n - 1) / n
		RowsRP(height, rp)(rect)
	}
}

// NColumnsRP calls the given RP on each of n vertical rectangles that span the
// input rectangle.
func NColumnsRP(n int, rp RP) RP {
	if n <= 0 {
		return noopRP
	}

	return func(rect image.Rectangle) {
		width := (rect.Dx() + n - 1) / n
		ColumnsRP(width, rp)(rect)
	}
}

// NRectanglesRP calls the given RP on each of n horizontal or vertical
// rectangles that span the input rectangle.
func NRectanglesRP(n int, rp RP) RP {
	if n <= 0 {
		return noopRP
	}

	return func(rect image.Rectangle) {
		if rect.Dx() > rect.Dy() {
			NColumnsRP(n, rp)(rect)
		} else {
			NRowsRP(n, rp)(rect)
		}
	}
}

// QuickRP
func QuickRP(rp RP) RP {
	gomaxprocs := runtime.GOMAXPROCS(-1)
	if gomaxprocs == 1 {
		return rp
	}

	return func(rect image.Rectangle) {

		// Create a new wait group and defer the wait.
		var w sync.WaitGroup
		defer w.Wait()

		// Wrap the processor with an asynchronous processor.
		rp = ConcurrentRP(&w, rp)

		// Wrap the processor with a set count columns processor.
		rp = NRectanglesRP(gomaxprocs, rp)

		// Call the processor on the entire bounds.
		rp(rect)
	}
}
