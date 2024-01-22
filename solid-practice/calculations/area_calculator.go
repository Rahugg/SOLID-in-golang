package calculations

import solidpractice "solid/solid-practice"

type AreaCalculator struct {
	Shapes []solidpractice.ManageShapeInterface
}

func NewAreaCalculator(shapes ...solidpractice.ManageShapeInterface) *AreaCalculator {
	return &AreaCalculator{Shapes: shapes}
}

func (ac *AreaCalculator) Sum() float64 {
	var totalArea float64
	for _, shape := range ac.Shapes {
		totalArea += shape.Calculate()
	}
	return totalArea
}
