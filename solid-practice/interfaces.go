package solid_practice

type Shape interface {
	Area() float64
}

type ThreeDimensionalShapeInterface interface {
	Volume() float64
}

type ManageShapeInterface interface {
	Calculate() float64
}

type ManageShapeCalculators interface {
	Sum() float64
}
