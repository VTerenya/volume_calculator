//Pachage handle has handle, which convert
//input in necessary values.
package handle

import (
	"flag"
	"github.com/shopspring/decimal"
	"volume_calculator/calculator"
)

func Handle() *calculator.Calculator{
	var shape string
	flag.StringVar(&shape, "shape", "string", "a string")
	radius, _ := decimal.NewFromString(*flag.String("radius", "", "a float64"))
	length, _ := decimal.NewFromString(*flag.String("length", "", "a float64"))
	width, _ := decimal.NewFromString(*flag.String("width", "", "a float64"))
	hieght, _ := decimal.NewFromString(*flag.String("hieght", "", "a float64"))
	flag.Parse()
	calc := calculator.NewCalc(shape, radius, length, width, hieght)
	return calc
}
