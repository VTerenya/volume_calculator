//Package has all shapes, which
//calculator can count volume.
package shapes

import (
	"github.com/shopspring/decimal"
	"math"
)

var pi = decimal.NewFromFloat(math.Pi)

type Volumer interface {
	volume() decimal.Decimal
}

//Sphere is struct, which has only radius.
type Sphere struct {
	Radius decimal.Decimal
}

//Pyramid is struct, which has length, width, hieght
type Pyramid struct {
	Length, Width, Hieght decimal.Decimal
}

//Cylinder is struct, which has radius, hieght
type Cylinder struct {
	Radius, Hieght decimal.Decimal
}

//Volume is realization volume() for Sphere
func (Sphere *Sphere) Volume() decimal.Decimal {
	return pi.Mul(decimal.NewFromFloat(4.0 / 3)).Mul(Sphere.Radius.Pow(decimal.NewFromFloat(3.0)))
}

//Volume is realization volume() for Pyramid
func (Pyramid *Pyramid) Volume() decimal.Decimal {
	return decimal.NewFromFloat(1.0 / 3).Mul(Pyramid.Hieght).Mul(Pyramid.Length).Mul(Pyramid.Width)
}

//Volume is realization volume() for Cylinder
func (Cylinder *Cylinder) Volume() decimal.Decimal {
	return Cylinder.Hieght.Mul(pi.Mul(Cylinder.Radius.Pow(decimal.NewFromFloat(2.0))))
}
