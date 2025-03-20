package Struct

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shapes  Shapes
		hasArea float64
	}{
		{name: "Rectangle", shapes: Rectangle{length: 10.0, breadth: 10.0}, hasArea: 40.0},
		{name: "Circle", shapes: Circle{radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shapes: Triangle{height: 2.0, width: 4.0}, hasArea: 4.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shapes.Area()
			if got != tt.hasArea {
				t.Errorf("%g is got and %g is Area %#v shape", got, tt.hasArea, tt.shapes)
			}
		})
	}
}
