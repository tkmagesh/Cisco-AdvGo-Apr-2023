package models

type Rectangle struct {
	Length  float32
	Breadth float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Breadth
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Length + r.Breadth)
}
