// Package calculator has code, which describe
// volume calculator.
package calculator

import (
	"github.com/shopspring/decimal"
	"volume_calculator/internal/builder"
)

// Calculate is the implementation of counting the volume of all shapes.
func Calculate(shape string, radius, length, width, hieght decimal.Decimal) decimal.Decimal {
	b := builder.NewBuilder(shape, radius, length, width, hieght)
	sh := b.BuildShape()
	return sh.CalculateVolume()
}
