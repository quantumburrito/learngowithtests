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

	areaTests := []struct{
		shape Shape
		want float64
	} {
		{shape: Rectangle{Width: 10.0, Height: 10.0}, want: 100},
		{shape: Circle{Radius: 10.0}, want: 10.0*10.0*math.Pi},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		hasArea := shape.Area()
		if want != hasArea {
			t.Errorf("%#v Wanted: %g Got: %g", shape, want, hasArea)
		}

	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}
}