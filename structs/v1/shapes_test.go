package Structs

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0
	if got != want {
		t.Errorf("%.2f is got and %.2f is want", got, want)
	}
}

func TestArea(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Area(rect)
	want := 100.0
	if got != want {
		t.Errorf("%.2f is got and %.2f is want", got, want)
	}
}
