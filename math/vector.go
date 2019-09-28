package math

// Vector3f is a vector in 3D space.
type Vector3f [3]Float

// NewVector3f constructs a new Vector3 instance with specified values.
func NewVector3f(x, y, z Float) Vector3f {
	return Vector3f{x, y, z}
}

// Inverse returns the inverse of vector.
func (v Vector3f) Inverse() Vector3f {
	return NewVector3f(
		1.0/v.X(),
		1.0/v.Y(),
		1.0/v.Z(),
	)
}

func (v Vector3f) X() Float {
	return v[0]
}

func (v Vector3f) Y() Float {
	return v[1]
}

func (v Vector3f) Z() Float {
	return v[2]
}
