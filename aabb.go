package lightspeed

import "github.com/chenzhekl/lightspeed/math"

// AABB is short for Axis-alighed bounding box.
type AABB struct {
	Min, Max math.Vector3
}

// NewAABB constructs a new AABB instnace with specified values.
func NewAABB(min, max math.Vector3) AABB {
	return AABB{
		Min: min,
		Max: max,
	}
}

// Hit checks if the ray hits this AABB.
func (b *AABB) Hit(ray *Ray) bool {
	// FIXME: implement it!
	return false
}
