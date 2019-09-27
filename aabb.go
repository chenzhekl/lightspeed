package lightspeed

import (
	"fmt"

	"github.com/chenzhekl/lightspeed/math"
)

// AABB is short for axis-alighed bounding box. It's the smallest box that contains the object.
type AABB struct {
	Min, Max math.Point3f
}

// NewAABB constructs a new AABB instnace with specified values.
func NewAABB(min, max math.Point3f) AABB {
	return AABB{
		Min: min,
		Max: max,
	}
}

// String provides the string representation for AABB.
func (b AABB) String() string {
	return fmt.Sprintf("{Min: %v, Max: %v}", b.Min, b.Max)
}

// Hit checks if the ray hits this AABB.
func (b AABB) Hit(ray *Ray) bool {
	intervalMin := ray.TMin
	intervalMax := ray.TMax

	var tMin, tMax math.Float
	tMin = (b.Min.X - ray.Origin.X) * ray.directionInv.X
	tMax = (b.Max.X - ray.Origin.X) * ray.directionInv.X
	if ray.directionInv.X < 0.0 {
		tMin, tMax = tMax, tMin
	}
	if tMin > intervalMin {
		intervalMin = tMin
	}
	if tMax < intervalMax {
		intervalMax = tMax
	}
	if intervalMin > intervalMax {
		return false
	}

	tMin = (b.Min.Y - ray.Origin.Y) * ray.directionInv.Y
	tMax = (b.Max.Y - ray.Origin.Y) * ray.directionInv.Y
	if ray.directionInv.Y < 0.0 {
		tMin, tMax = tMax, tMin
	}
	if tMin > intervalMin {
		intervalMin = tMin
	}
	if tMax < intervalMax {
		intervalMax = tMax
	}
	if intervalMin > intervalMax {
		return false
	}

	tMin = (b.Min.Z - ray.Origin.Z) * ray.directionInv.Z
	tMax = (b.Max.Z - ray.Origin.Z) * ray.directionInv.Z
	if ray.directionInv.Z < 0.0 {
		tMin, tMax = tMax, tMin
	}
	if tMin > intervalMin {
		intervalMin = tMin
	}
	if tMax < intervalMax {
		intervalMax = tMax
	}
	return intervalMin <= intervalMax
}
