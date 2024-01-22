package shapes

import "math"

type Sphere struct {
	Radius float64
}

func NewSphere(radius float64) *Sphere {
	return &Sphere{Radius: radius}
}

func (s *Sphere) Area() float64 {
	return 4 * math.Pi * math.Pow(s.Radius, 2)
}

func (s *Sphere) Volume() float64 {
	return (4 / 3) * math.Pi * math.Pow(s.Radius, 3)
}

func (s *Sphere) Calculate() float64 {
	return s.Volume()
}
