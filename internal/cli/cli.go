// Package cli has cli, which convert
// input in necessary values.
package cli

import (
	"flag"
	"github.com/shopspring/decimal"
)

type Parameters struct {
	Shape                         string
	Radius, Length, Width, Hieght decimal.Decimal
}

func newParameters(shape string, radius, length, width, hieght decimal.Decimal) *Parameters {
	return &Parameters{Shape: shape, Radius: radius, Length: length, Width: width, Hieght: hieght}
}

func Load() *Parameters {
	var shape string
	flag.StringVar(&shape, "shape", "", "a string")
	var stringRadius, stringLength, stringWidth, stringHieght string
	flag.StringVar(&stringRadius, "radius", "", "a string")
	flag.StringVar(&stringLength, "length", "", "a string")
	flag.StringVar(&stringWidth, "width", "", "a string")
	flag.StringVar(&stringHieght, "hieght", "", "a string")
	flag.Parse()

	radius, _ := decimal.NewFromString(stringRadius)
	length, _ := decimal.NewFromString(stringLength)
	width, _ := decimal.NewFromString(stringWidth)
	hieght, _ := decimal.NewFromString(stringHieght)
	return newParameters(shape, radius, length, width, hieght)
}
