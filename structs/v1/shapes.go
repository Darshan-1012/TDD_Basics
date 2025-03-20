package Structs

type Rectangle struct {
	length  float64
	breadth float64
}

func Perimeter(rect Rectangle) float64 {
	var lengthRect = rect.length + rect.breadth
	return 2 * lengthRect
}

func Area(rect Rectangle) float64 {
	return rect.length * rect.breadth
}
