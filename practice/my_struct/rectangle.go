package my_struct

type rectangle struct {
	shape
	Width, Height float64
}

func (r rectangle) area() float64 {
	return r.Width * r.Height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
