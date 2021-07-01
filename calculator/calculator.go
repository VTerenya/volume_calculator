//Package has code, which describe
//volume calculator.
package calculator

import (
	"fmt"
	"github.com/shopspring/decimal"
	"os"
	"volume_calculator/shapes"
)

type Calculator struct {
	shape  string
	radius decimal.Decimal
	length decimal.Decimal
	width  decimal.Decimal
	hieght decimal.Decimal
}

//Volume is the implementation of counting the volume of all shapes.
func (calc Calculator) Volume() decimal.Decimal {
	switch calc.shape {
	case "sphere":
		sphere := shapes.Sphere{Radius: calc.radius}
		return sphere.Volume()
	case "pyramid":
		pyramid := shapes.Pyramid{Length: calc.length, Width: calc.width, Hieght: calc.hieght}
		return pyramid.Volume()
	case "cylinder":
		cylinder := shapes.Cylinder{Hieght: calc.hieght, Radius: calc.radius}
		return cylinder.Volume()
	default:
		fmt.Println("Error! Incorrect input!")
		os.Exit(1)
		return decimal.NewFromFloat(1)
	}
}

//NewCalc is constructor.
func NewCalc(shape string, radius, length, width, hieght decimal.Decimal) *Calculator {
	return &Calculator{shape: shape, radius: radius, length: length, width: width, hieght: hieght}
}
