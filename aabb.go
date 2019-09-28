package lightspeed

import (
	"fmt"

	"github.com/chenzhekl/lightspeed/math"
)

type IAABB interface {
	AABB() AABB
}

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

// AABB returns self.
func (b AABB) AABB() AABB {
	return b
}

// Union returns the union of two AABB.
func (b AABB) Union(other AABB) AABB {
	return NewAABB(b.Min.Min(other.Min), b.Max.Max(other.Max))
}

// Intersect checks if a ray hits this AABB.
func (b AABB) Intersect(ray *Ray) bool {
	intervalMin := ray.TMin
	intervalMax := ray.TMax

	var tMin, tMax math.Float
	tMin = (b.Min.X() - ray.Origin.X()) * ray.DirectionInv().X()
	tMax = (b.Max.X() - ray.Origin.X()) * ray.DirectionInv().X()
	if ray.DirectionInv().X() < 0.0 {
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

	tMin = (b.Min.Y() - ray.Origin.Y()) * ray.DirectionInv().Y()
	tMax = (b.Max.Y() - ray.Origin.Y()) * ray.DirectionInv().Y()
	if ray.DirectionInv().Y() < 0.0 {
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

	tMin = (b.Min.Z() - ray.Origin.Z()) * ray.DirectionInv().Z()
	tMax = (b.Max.Z() - ray.Origin.Z()) * ray.DirectionInv().Z()
	if ray.DirectionInv().Z() < 0.0 {
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
