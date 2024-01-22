package main

import "math"

// this is not right
type Rectangle struct {
	Width  float64
	Height float64
}

func Area(rectangle *Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

// this is right
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
