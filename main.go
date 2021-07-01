//Main package.
package main

import (
	"flag"
	"fmt"
	"github.com/shopspring/decimal"
	"volume_calculator/calculator"
)

func calculate(){
	var shape string
	flag.StringVar(&shape, "shape", "string", "a string")
	radius := flag.Float64("radius", -1.0, "a float64")
	length := flag.Float64("length", -1.0, "a float64")
	width := flag.Float64("width", -1.0, "a float64")
	hieght := flag.Float64("hieght", -1.0, "a float64")
	flag.Parse()
	calc := calculator.NewCalc(shape, decimal.NewFromFloat(*radius), decimal.NewFromFloat(*length),
		decimal.NewFromFloat(*width), decimal.NewFromFloat(*hieght))
	fmt.Println(calc.Volume())
}

func main() {
	calculate()
}
