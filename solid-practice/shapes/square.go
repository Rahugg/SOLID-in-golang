package shapes

import "math"

type Square struct {
	Length float64
}

func NewSquare(length float64) *Square {
	return &Square{Length: length}
}

func (s *Square) Area() float64 {
	return math.Pow(s.Length, 2)
}

func (s *Square) Calculate() float64 {
	return s.Area()
}
