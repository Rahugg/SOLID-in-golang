package shapes

import "math"

type Circle struct {
	Radius float64
}

func (c *Circle) Calculate() float64 {
	return c.Area()
}

func NewCircle(radius float64) *Circle {
	return &Circle{Radius: radius}
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
