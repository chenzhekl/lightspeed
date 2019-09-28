package math

// Point3f is a point in 3D space.
type Point3f [3]Float

// NewPoint3f constructs a new Point3 instance with specified values.
func NewPoint3f(x, y, z Float) Point3f {
	return Point3f{x, y, z}
}

// OriginPoint3f is a point at origin.
var OriginPoint3f = Point3f{0.0, 0.0, 0.0}

func (p Point3f) X() Float {
	return p[0]
}

func (p Point3f) Y() Float {
	return p[1]
}

func (p Point3f) Z() Float {
	return p[2]
}

func (p Point3f) Min(other Point3f) Point3f {
	for i := 0; i < 3; i++ {
		if other[i] < p[i] {
			p[i] = other[i]
		}
	}

	return p
}

func (p Point3f) Max(other Point3f) Point3f {
	for i := 0; i < 3; i++ {
		if other[i] > p[i] {
			p[i] = other[i]
		}
	}

	return p
}

func (p Point3f) Add(other Point3f) Point3f {
	return NewPoint3f(p.X()+other.X(), p.Y()+other.Y(), p.Z()+other.Z())
}

func (p Point3f) Div(factor Float) Point3f {
	return NewPoint3f(p.X()/factor, p.Y()/factor, p.Z()/factor)
}
