package lightspeed

import (
	"fmt"

	"github.com/chenzhekl/lightspeed/math"
)

// Ray represents a ray in path tracing algorithms.
type Ray struct {
	Origin    math.Point3f
	Direction math.Vector3f

	TMin math.Float
	TMax math.Float

	directionInv math.Vector3f
}

// NewRay constructs a new Ray with given origin and direction.
func NewRay(origin math.Point3f, direction math.Vector3f) *Ray {
	ray := &Ray{
		Origin:       origin,
		Direction:    direction,
		TMin:         1e-6,
		TMax:         math.MaxFloat,
		directionInv: direction.Inverse(),
	}

	return ray
}

// String provides the string representation for Ray.
func (r *Ray) String() string {
	return fmt.Sprintf("{Origin: %v, Direction: %v}", r.Origin, r.Direction)
}
