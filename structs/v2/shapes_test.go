package Struct

import "testing"

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shapes Shapes, want float64) {
		t.Helper()
		got := shapes.Area()
		if got != want {
			t.Errorf("%g is got and %g is want", got, want)
		}
	}
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
