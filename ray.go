package lightspeed

import (
	"github.com/chenzhekl/lightspeed/math"
)

// Ray represents a ray in path tracing algorithms.
type Ray struct {
	Origin    math.Point3
	Direction math.Vector3

	TMin math.Float
	TMax math.Float
}

// NewRay constructs a new Ray with given origin and direction.
func NewRay(origin math.Point3, direction math.Vector3) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: direction,
		TMin:      1e-6,
		TMax:      math.MaxFloat,
	}
}
