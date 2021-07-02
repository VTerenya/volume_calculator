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
	radius decimal.Decimal
}

func (s *Sphere) New(radius decimal.Decimal) *Sphere {
	return &Sphere{
		radius,
	}
}

// Pyramid is struct, which has length, width, hieght
type Pyramid struct {
	length, width, hieght decimal.Decimal
}

func (p *Pyramid) New(length, width, hieght decimal.Decimal) *Pyramid {
	return &Pyramid{
		length, width, hieght,
	}
}

// Cylinder is struct, which has radius, hieght
type Cylinder struct {
	radius, hieght decimal.Decimal
}

func (c *Cylinder) New(radius, hieght decimal.Decimal) *Cylinder {
	return &Cylinder{
		radius, hieght,
	}
}

// CalculateVolume calculateVolume is realization calculateVolume() for Sphere
func (Sphere *Sphere) CalculateVolume() decimal.Decimal {
	return pi.Mul(decimal.NewFromFloat(4.0 / 3)).Mul(Sphere.radius.Pow(decimal.NewFromFloat(3.0)))
}

// CalculateVolume calculateVolume is realization calculateVolume() for Pyramid
func (Pyramid *Pyramid) CalculateVolume() decimal.Decimal {
	return decimal.NewFromFloat(1.0 / 3).Mul(Pyramid.hieght).Mul(Pyramid.length).Mul(Pyramid.width)
}

// CalculateVolume calculateVolume is realization calculateVolume() for Cylinder
func (Cylinder *Cylinder) CalculateVolume() decimal.Decimal {
	return Cylinder.hieght.Mul(pi.Mul(Cylinder.radius.Pow(decimal.NewFromFloat(2.0))))
}