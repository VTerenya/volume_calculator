// Package builder has function, which
// build necessary shape.
package builder

import (
	"github.com/shopspring/decimal"
	"volume_calculator/internal/shapes"
)

type Builder struct {

}

func (b Builder) BuildSphere(radius decimal.Decimal) *shapes.Sphere {
	sphere := shapes.Sphere{}
	return sphere.New(radius)
}

func (b Builder) BuildPyramid(length, width, hieght decimal.Decimal) *shapes.Pyramid {
	pyramid := shapes.Pyramid{}
	return pyramid.New(length, width, hieght)
}

func (b Builder) BuildCylinder(hieght, radius decimal.Decimal) *shapes.Cylinder {
	cylinder := shapes.Cylinder{}
	return cylinder.New(radius, hieght)
}