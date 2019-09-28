package lightspeed

import (
	"github.com/chenzhekl/lightspeed/math"
	"testing"
)

func TestQsplit(t *testing.T) {
	objs := []IAABB{
		NewAABB(math.OriginPoint3f, math.NewPoint3f(1.0, 1.0, 1.0)),
		NewAABB(math.NewPoint3f(1.5, 1.0, 1.0), math.NewPoint3f(2.0, 1.0, 1.0)),
		NewAABB(math.NewPoint3f(0.5, 1.0, 1.0), math.NewPoint3f(1.0, 1.0, 1.0)),
		NewAABB(math.NewPoint3f(1.0, 1.0, 1.0), math.NewPoint3f(1.2, 1.0, 1.0)),
	}

	splitPoint := qsplit(objs, 0.0, 0)
	if splitPoint != 2 {
		t.Errorf("splitPoint should be 2, but get %d", splitPoint)
	}

	splitPoint = qsplit(objs, 1.0, 0)
	if splitPoint != 2 {
		t.Errorf("splitPoint should be 2, but get %d", splitPoint)
	}
}

func TestNewBVH(t *testing.T) {
	objs := []IAABB{
		NewAABB(math.OriginPoint3f, math.NewPoint3f(1.0, 1.0, 1.0)),
		NewAABB(math.NewPoint3f(1.0, 1.0, 1.0), math.NewPoint3f(2.0, 2.0, 2.0)),
		NewAABB(math.NewPoint3f(2.0, 2.0, 2.0), math.NewPoint3f(3.0, 3.0, 3.0)),
		NewAABB(math.NewPoint3f(3.0, 3.0, 3.0), math.NewPoint3f(4.0, 4.0, 4.0)),
	}

	bvh := NewBVH(objs)

	if bvh.left.left.data != objs[0] || bvh.left.right.data != objs[1] || bvh.right.left.data != objs[2] || bvh.right.right.data != objs[3] {
		t.Error("BVH is not correctly built")
	}

	if bvh.AABB() != NewAABB(math.OriginPoint3f, math.NewPoint3f(4.0, 4.0, 4.0)) {
		t.Error("BVH is not correctly built")
	}
}

func TestBVH_Intersect(t *testing.T) {
	objs := []IAABB{
		NewAABB(math.OriginPoint3f, math.NewPoint3f(1.0, 1.0, 1.0)),
		NewAABB(math.NewPoint3f(1.0, 1.0, 1.0), math.NewPoint3f(2.0, 2.0, 2.0)),
		NewAABB(math.NewPoint3f(2.0, 2.0, 2.0), math.NewPoint3f(3.0, 3.0, 3.0)),
		NewAABB(math.NewPoint3f(3.0, 3.0, 3.0), math.NewPoint3f(4.0, 4.0, 4.0)),
	}
	bvh := NewBVH(objs)
	ray1 := NewRay(math.NewPoint3f(0.5, 0.5, 0.5), math.NewVector3f(1.0, 1.0, 1.0))
	ray2 := NewRay(math.NewPoint3f(0.5, 0.5, 0.5), math.NewVector3f(-1.0, -1.0, -1.0))

	hits := bvh.Intersect(ray1)
	if len(hits) != 4 {
		t.Errorf("Expect 4 his, but got %d", len(hits))
	}

	hits = bvh.Intersect(ray2)
	if len(hits) != 1 {
		t.Errorf("Expect 1 his, but got %d", len(hits))
	}
}
