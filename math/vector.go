package math

// Vector3f is a vector in 3D space.
type Vector3f struct {
	X, Y, Z Float
}

// NewVector3f constructs a new Vector3 instance with specified values.
func NewVector3f(x, y, z Float) Vector3f {
	return Vector3f{
		X: x,
		Y: y,
		Z: z,
	}
}

// Inverse returns the inverse of vector.
func (v Vector3f) Inverse() Vector3f {
	return Vector3f{
		X: 1.0 / v.X,
		Y: 1.0 / v.Y,
		Z: 1.0 / v.Z,
	}
}
