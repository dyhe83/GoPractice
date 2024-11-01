package my_struct

import (
	"fmt"
)

func PrintShapeInfo(s shape) {
	fmt.Printf("Area: %f\n", s.area())
	fmt.Printf("Perimeter: %f\n", s.perimeter())
}

func Main() {
	r := rectangle{Width: 3, Height: 4}
	PrintShapeInfo(r)

	nr := NewRectangle{rectangle{Width: 3, Height: 4}}
	PrintShapeInfo(nr)

	// c := FakeCircle{Radius: 10}
	// PrintShapeInfo(c)
}
