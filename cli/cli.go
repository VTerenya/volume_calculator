// Package cli has cli, which convert
// input in necessary values.
package cli

import (
	"flag"
	"github.com/shopspring/decimal"
	"volume_calculator/calculator"
)

type Parameters struct{
	shape string
	radius, length, width, hieght decimal.Decimal
}

func newParameters(shape string, radius, length, width, hieght decimal.Decimal) *Parameters {
	return &Parameters{shape: shape, radius: radius,length: length,width: width,hieght: hieght}
}

func load() *Parameters {
	var shape string
	flag.StringVar(&shape, "shape", "string", "a string")
	radius, _ := decimal.NewFromString(*flag.String("radius", "", "a float64"))
	length, _ := decimal.NewFromString(*flag.String("length", "", "a float64"))
	width, _ := decimal.NewFromString(*flag.String("width", "", "a float64"))
	hieght, _ := decimal.NewFromString(*flag.String("hieght", "", "a float64"))
	flag.Parse()
	return newParameters(shape,radius,length,width,hieght)
}


func Handle() *calculator.Calculator{
	param := load()
	calc := calculator.NewCalc(param.shape, param.radius, param.length, param.width, param.hieght)
	return calc
}
