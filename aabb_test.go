package lightspeed

import (
	"testing"

	"github.com/chenzhekl/lightspeed/math"
)

func TestAABB_Intersect(t *testing.T) {
	aabb := NewAABB(math.OriginPoint3f, math.NewPoint3f(1.0, 1.0, 1.0))
	ray1 := NewRay(math.NewPoint3f(-1.0, -1.0, -1.0), math.NewVector3f(1.0, 1.0, 1.0))
	ray2 := NewRay(math.NewPoint3f(-1.0, -1.0, -1.0), math.NewVector3f(-1.0, 1.0, 1.0))

	if !aabb.Intersect(ray1) {
		t.Errorf("Ray %v should hit AABB %v", ray1, aabb)
	}

	if aabb.Intersect(ray2) {
		t.Errorf("Ray %v should not hit AABB %v", ray2, aabb)
	}
}
