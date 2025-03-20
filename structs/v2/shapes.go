package Struct

import "math"

type Rectangle struct {
	length  float64
	breadth float64
}

type Circle struct {
	radius float64
}

type Shapes interface {
	Area() float64
}

func (rect Rectangle) Area() float64 {
	return rect.breadth * rect.length
}

// Using methods - func() args() {}
func (circle Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}
