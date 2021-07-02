//Package has all shapes, which
//calculator can count calculateVolume.
package shapes

import (
	"github.com/shopspring/decimal"
	"math"
)

var pi = decimal.NewFromFloat(math.Pi)

type Volumer interface {
	CalculateVolume() decimal.Decimal
}

// Sphere is struct, which has only radius.
type Sphere struct {
	Radius decimal.Decimal
}

// Pyramid is struct, which has length, width, hieght
type Pyramid struct {
	Length, Width, Hieght decimal.Decimal
}

// Cylinder is struct, which has radius, hieght
type Cylinder struct {
	Radius, Hieght decimal.Decimal
}

// CalculateVolume calculateVolume is realization calculateVolume() for Sphere
func (Sphere *Sphere) CalculateVolume() decimal.Decimal {
	return pi.Mul(decimal.NewFromFloat(4.0 / 3)).Mul(Sphere.Radius.Pow(decimal.NewFromFloat(3.0)))
}

// CalculateVolume calculateVolume is realization calculateVolume() for Pyramid
func (Pyramid *Pyramid) CalculateVolume() decimal.Decimal {
	return decimal.NewFromFloat(1.0 / 3).Mul(Pyramid.Hieght).Mul(Pyramid.Length).Mul(Pyramid.Width)
}

// CalculateVolume calculateVolume is realization calculateVolume() for Cylinder
func (Cylinder *Cylinder) CalculateVolume() decimal.Decimal {
	return Cylinder.Hieght.Mul(pi.Mul(Cylinder.Radius.Pow(decimal.NewFromFloat(2.0))))
}
