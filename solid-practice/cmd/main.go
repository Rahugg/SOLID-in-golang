package main

import (
	"fmt"
	"solid/solid-practice"
	"solid/solid-practice/calculations"
	shapesPkg "solid/solid-practice/shapes"
)

func main() {
	shapes := []solid_practice.ManageShapeInterface{
		shapesPkg.NewCircle(2),
		shapesPkg.NewSquare(5),
		shapesPkg.NewSquare(6),
	}

	solidShapes := []solid_practice.ManageShapeInterface{
		shapesPkg.NewCuboid(3, 4, 5),
		shapesPkg.NewSphere(4),
	}

	areas := calculations.NewAreaCalculator(shapes...)
	volumes := calculations.NewVolumeCalculator(solidShapes...)

	output := solid_practice.NewSumCalculatorOutput(areas)
	output2 := solid_practice.NewSumCalculatorOutput(volumes)

	fmt.Println(output.HTML())
	fmt.Println(output2.HTML())
}
