package lightspeed

import (
	"fmt"

	"github.com/chenzhekl/lightspeed/math"
)

// Ray represents a ray in path tracing algorithms.
type Ray struct {
	Origin       math.Point3f
	direction    math.Vector3f
	directionInv math.Vector3f

	TMin math.Float
	TMax math.Float
}

// NewRay constructs a new Ray with given origin and direction.
func NewRay(origin math.Point3f, direction math.Vector3f) *Ray {
	ray := &Ray{
		Origin:       origin,
		TMin:         1e-6,
		TMax:         math.MaxFloat,
		direction:    direction,
		directionInv: direction.Inverse(),
	}

	return ray
}

// String provides the string representation for Ray.
func (r *Ray) String() string {
	return fmt.Sprintf("{Origin: %v, Direction: %v}", r.Origin, r.Direction())
}

// Direction returns the direction of ray
func (r *Ray) Direction() math.Vector3f {
	return r.direction
}

// Direction sets the direction of ray
func (r *Ray) SetDirection(direction math.Vector3f) {
	r.direction = direction
	r.directionInv = direction.Inverse()
}

// DirectionInv returns the inverse of direction of ray
func (r *Ray) DirectionInv() math.Vector3f {
	return r.directionInv
}
