package structures_methods_interfaces

import "math"

type Form interface {
	Area() float64
}

type Retangle struct {
	Length float64
	Width  float64
}

func (r Retangle) Area() float64 {
	return r.Length * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(retangle Retangle) float64 {
	return 2 * (retangle.Length + retangle.Width)
}
