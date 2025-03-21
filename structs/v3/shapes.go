package Struct

import "math"

// Further recfactoring

type Rectangle struct {
	length  float64
	breadth float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	height float64
	width  float64
}

type Shapes interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return 2 * (r.length + r.breadth)
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.height * t.width)
}
