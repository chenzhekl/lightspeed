package lightspeed

import "github.com/chenzhekl/lightspeed/math"

// BVH is short for bounding volume hierarchy.
// It's used for accelerating ray object hit in 3D space.
type BVH struct {
	data  IAABB
	left  *BVH
	right *BVH
	aabb  AABB
}

func NewBVH(objs []IAABB) *BVH {
	return buildBranch(objs, 0)
}

func (b *BVH) AABB() AABB {
	return b.aabb
}

func (b *BVH) Intersect(ray *Ray) []IAABB {
	hits := []IAABB{}
	b.intersect(ray, &hits)
	return hits
}

func (b *BVH) intersect(ray *Ray, hits *[]IAABB) {
	if !b.aabb.Intersect(ray) {
		return
	}
	if b.data != nil {
		*hits = append(*hits, b.data)
		return
	}

	if b.left != nil {
		b.left.intersect(ray, hits)
	}
	if b.right != nil {
		b.right.intersect(ray, hits)
	}
}

func buildBranch(objs []IAABB, axis int) *BVH {
	if len(objs) == 1 {
		return &BVH{
			data:  objs[0],
			left:  nil,
			right: nil,
			aabb:  objs[0].AABB(),
		}
	}
	if len(objs) == 2 {
		return &BVH{
			data: nil,
			left: &BVH{
				data:  objs[0],
				left:  nil,
				right: nil,
				aabb:  objs[0].AABB(),
			},
			right: &BVH{
				data:  objs[1],
				left:  nil,
				right: nil,
				aabb:  objs[1].AABB(),
			},
			aabb: objs[0].AABB().Union(objs[1].AABB()),
		}
	}

	bbox := objs[0].AABB()
	for i := 1; i < len(objs); i++ {
		bbox = bbox.Union(objs[i].AABB())
	}
	pivot := bbox.Min.Add(bbox.Max).Div(2.0)
	midPoint := qsplit(objs, pivot.X(), axis)

	return &BVH{
		data:  nil,
		left:  buildBranch(objs[:midPoint], (axis+1)%3),
		right: buildBranch(objs[midPoint:], (axis+1)%3),
		aabb:  bbox,
	}
}

func qsplit(objs []IAABB, pivot math.Float, axis int) int {
	ret := 0

	for i := 0; i < len(objs); i++ {
		bbox := objs[i].AABB()
		centroid := (bbox.Min[axis] + bbox.Max[axis]) * 0.5
		if centroid < pivot {
			objs[i], objs[ret] = objs[ret], objs[i]
			ret++
		}
	}

	if ret == 0 || ret == len(objs) {
		ret = len(objs) / 2
	}

	return ret
}
