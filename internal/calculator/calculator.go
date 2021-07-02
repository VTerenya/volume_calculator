// Package calculator has code, which describe
// volume calculator.
package calculator

import (
	"github.com/shopspring/decimal"
	"volume_calculator/internal/builder"
)

type Calculator struct {
}

// Calculate is the implementation of counting the volume of all shapes.
func (calc Calculator) Calculate(shape string, radius, length, width, hieght decimal.Decimal) decimal.Decimal {
	var volume decimal.Decimal
	builder := builder.Builder{}
	switch shape {
	case "sphere":
		volume = builder.BuildSphere(radius).CalculateVolume()
	case "pyramid":
		volume = builder.BuildPyramid(length, width, hieght).CalculateVolume()
	case "cylinder":
		volume = builder.BuildCylinder(hieght, radius).CalculateVolume()
	}
	return volume
}
