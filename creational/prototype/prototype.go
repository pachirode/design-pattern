package prototype

type Shape interface {
	Clone() Shape
	GetType() string
}

type Circle struct {
	Type string
}

func (c *Circle) Clone() Shape {
	return &Circle{Type: c.Type}
}

func (c *Circle) GetType() string {
	return c.Type
}
