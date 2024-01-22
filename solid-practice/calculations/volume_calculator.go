package calculations

import solidpractice "solid/solid-practice"

type VolumeCalculator struct {
	Shapes []solidpractice.ManageShapeInterface
}

func NewVolumeCalculator(shapes ...solidpractice.ManageShapeInterface) *VolumeCalculator {
	return &VolumeCalculator{Shapes: shapes}
}

func (vc *VolumeCalculator) Sum() float64 {
	var totalVolume float64
	for _, shape := range vc.Shapes {
		totalVolume += shape.Calculate()
	}
	return totalVolume
}
