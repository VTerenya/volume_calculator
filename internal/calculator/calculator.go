// Package calculator has code, which describe
// volume calculator.
package calculator

import (
	"github.com/VTerenya/volume_calculator/internal/builder"
	"github.com/shopspring/decimal"
)

// Calculate is the implementation of counting the volume of all shapes.
func Calculate(shape string, radius, length, width, hieght decimal.Decimal) (decimal.Decimal, error) {
	b := builder.NewBuilder(shape, radius, length, width, hieght)
	sh, err := b.BuildShape()
	if err != nil {
		return decimal.Decimal{}, err
	} else {
		return sh.CalculateVolume(), nil
	}
}
