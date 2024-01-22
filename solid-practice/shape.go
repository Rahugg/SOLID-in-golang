package solid_practice

import (
	"encoding/json"
	"fmt"
)

type SumCalculatorOutput struct {
	Calculator ManageShapeCalculators
}

func NewSumCalculatorOutput(calculator ManageShapeCalculators) *SumCalculatorOutput {
	return &SumCalculatorOutput{Calculator: calculator}
}

func (sco *SumCalculatorOutput) JSON() (string, error) {
	data := map[string]float64{
		"sum": sco.Calculator.Sum(),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func (sco *SumCalculatorOutput) HTML() string {
	return fmt.Sprintf("Sum of the areas of provided shapes: %.2f", sco.Calculator.Sum())
}
