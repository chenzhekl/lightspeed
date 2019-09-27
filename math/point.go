package math

// Point3f is a point in 3D space.
type Point3f struct {
	X, Y, Z Float
}

// NewPoint3f constructs a new Point3 instance with specified values.
func NewPoint3f(x, y, z Float) Point3f {
	return Point3f{
		X: x,
		Y: y,
		Z: z,
	}
}

// Origin3f is a point at origin.
var Origin3f = Point3f{}
