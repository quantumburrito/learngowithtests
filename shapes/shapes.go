package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.height + rectangle.width)
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64{
	return math.Pow(c.radius, 2) * math.Pi
}