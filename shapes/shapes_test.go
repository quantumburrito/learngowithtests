package shapes

import "testing"


func TestPerimeter(t * testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if want != got {
		t.Errorf("Got: %.2f Want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10,10}
	got := Area(rectangle)
	want := 100.0
	if want != got {
		t.Errorf("Got : %.2f Want %.2f", got, want)
	}
}