package shapes

import (
	"github.com/shopspring/decimal"
	"math"
)

var pi = decimal.NewFromFloat(math.Pi)

type Sphere struct {
	Radius decimal.Decimal
}

type Pyramid struct {
	Length, Width, Hieght decimal.Decimal
}

type Cylinder struct {
	Radius, Hieght decimal.Decimal
}

func (Sphere *Sphere) Volume() decimal.Decimal {
	return pi.Mul(decimal.NewFromFloat(4.0 / 3)).Mul(Sphere.Radius.Pow(decimal.NewFromFloat(3.0)))
}

func (Pyramid *Pyramid) Volume() decimal.Decimal {
	return decimal.NewFromFloat(1.0 / 3).Mul(Pyramid.Hieght).Mul(Pyramid.Length).Mul(Pyramid.Width)
}

func (Cylinder *Cylinder) Volume() decimal.Decimal {
	return Cylinder.Hieght.Mul(pi.Mul(Cylinder.Radius.Pow(decimal.NewFromFloat(2.0))))
}
