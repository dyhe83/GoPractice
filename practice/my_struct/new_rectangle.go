package my_struct

type NewRectangle struct {
	rectangle
}

func (r NewRectangle) area() float64 {
	return r.Width * r.Height * 10
}
