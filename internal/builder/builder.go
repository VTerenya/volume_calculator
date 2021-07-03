// Package builder has function, which
// build necessary shape.
package builder

import (
	"github.com/shopspring/decimal"
	"volume_calculator/internal/shapes"
)

type Builder struct {
	sphere   shapes.Sphere
	pyramid  shapes.Pyramid
	cylinder shapes.Cylinder

	shape                         string
	radius, length, width, hieght decimal.Decimal
}

func NewBuilder(shape string, radius, length, width, hieght decimal.Decimal) *Builder {
	return &Builder{shape: shape, radius: radius, length: length, width: width, hieght: hieght}
}

func (builder Builder) BuildShape() shapes.Volumer{
	var sh shapes.Volumer
	switch builder.shape {
	case "sphere":
		sh = builder.buildSphere(builder.radius)
	case "pyramid":
		sh = builder.buildPyramid(builder.length, builder.width, builder.hieght)
	case "cylinder":
		sh = builder.buildCylinder(builder.hieght, builder.radius)
	}
	return sh
}

func (builder Builder) getSphere() shapes.Sphere {
	return builder.sphere
}

func (builder Builder) getPyramid() shapes.Pyramid {
	return builder.pyramid
}

func (builder Builder) getCylinder() shapes.Cylinder {
	return builder.cylinder
}

func (b Builder) buildSphere(radius decimal.Decimal) *shapes.Sphere {
	sphere := shapes.Sphere{}
	return sphere.NewSphere(radius)
}

func (b Builder) buildPyramid(length, width, hieght decimal.Decimal) *shapes.Pyramid {
	pyramid := shapes.Pyramid{}
	return pyramid.NewPyramid(length, width, hieght)
}

func (b Builder) buildCylinder(hieght, radius decimal.Decimal) *shapes.Cylinder {
	cylinder := shapes.Cylinder{}
	return cylinder.NewCylinder(radius, hieght)
}
