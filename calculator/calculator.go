// Package calculator has code, which describe
// volume calculator.
package calculator

import (
	"github.com/shopspring/decimal"
	"volume_calculator/shapes"
)

type Calculator struct {
	shape  string
	radius decimal.Decimal
	length decimal.Decimal
	width  decimal.Decimal
	hieght decimal.Decimal
}

func (calc Calculator) buildSphere(radius decimal.Decimal) *shapes.Sphere {
	return &shapes.Sphere{Radius: radius}
}

func (calc Calculator) buildPyramid(length, width, hieght decimal.Decimal) *shapes.Pyramid {
	return &shapes.Pyramid{Length: length, Width: width, Hieght: hieght}
}

func (calc Calculator) buildCylinder(hieght, radius decimal.Decimal) *shapes.Cylinder {
	return &shapes.Cylinder{Hieght: hieght, Radius: radius}
}

// CalcVolume is the implementation of counting the volume of all shapes.
func (calc Calculator) CalcVolume() decimal.Decimal{
	var result decimal.Decimal
	switch calc.shape {
	case "sphere":
		result = calc.buildSphere(calc.radius).CalculateVolume()
	case "pyramid":
		result = calc.buildPyramid(calc.length, calc.width, calc.hieght).CalculateVolume()
	case "cylinder":
		result = calc.buildCylinder(calc.hieght, calc.radius).CalculateVolume()
	}
	return result
}

// NewCalc is constructor.
func NewCalc(shape string, radius, length, width, hieght decimal.Decimal) *Calculator {
	return &Calculator{shape: shape, radius: radius, length: length, width: width, hieght: hieght}
}
