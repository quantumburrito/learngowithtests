package shapes

import (
	"math"
	"testing"
)


func TestPerimeter(t * testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if want != got {
		t.Errorf("Got: %.2f Want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if want != got {
			t.Errorf("Wanted: %g Got: %g", want, got)
		}

	}
	t.Run("Test function 'Area' on struct Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})
	t.Run("Test function 'Area' on struct Circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 10.0 * 10.0 * math.Pi)
	})
}