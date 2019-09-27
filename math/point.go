package math

// Point3 is the datatype for 3 dimendional point.
type Point3 struct {
	X, Y, Z Float
}

// NewPoint3 constructs a new Point3 instance with specified values.
func NewPoint3(x, y, z Float) Point3 {
	return Point3{
		X: x,
		Y: y,
		Z: z,
	}
}

// NewPoint3Zero constructs a new Point3 instance with all elements set to zero.
func NewPoint3Zero() Point3 {
	return Point3{}
}
