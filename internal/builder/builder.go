// Package builder has function, which
// build necessary shape.
package builder

import (
	"errors"
	"github.com/VTerenya/volume_calculator/internal/shapes"
	"github.com/shopspring/decimal"
)

type Builder struct {
	shape                         string
	radius, length, width, hieght decimal.Decimal
}

func NewBuilder(shape string, radius, length, width, hieght decimal.Decimal) *Builder {
	return &Builder{shape: shape, radius: radius, length: length, width: width, hieght: hieght}
}

func (builder Builder) BuildShape() (shapes.Volumer, error){
	var sh shapes.Volumer
	switch builder.shape {
	case "sphere":
		sh = builder.buildSphere(builder.radius)
	case "pyramid":
		sh = builder.buildPyramid(builder.length, builder.width, builder.hieght)
	case "cylinder":
		sh = builder.buildCylinder(builder.hieght, builder.radius)
	default:
		return nil, errors.New("Builder doesn't know this shape!")
	}
	return sh,nil
}

func (b Builder) buildSphere(radius decimal.Decimal) *shapes.Sphere {
	return shapes.NewSphere(radius)
}

func (b Builder) buildPyramid(length, width, hieght decimal.Decimal) *shapes.Pyramid {
	return shapes.NewPyramid(length, width, hieght)
}

func (b Builder) buildCylinder(hieght, radius decimal.Decimal) *shapes.Cylinder {
	return shapes.NewCylinder(radius, hieght)
}
