package math

// Vector3 is the datatype for 3 dimendional vector.
type Vector3 struct {
	X, Y, Z Float
}

// NewVector3 constructs a new Vector3 instance with specified values.
func NewVector3(x, y, z Float) Vector3 {
	return Vector3{
		X: x,
		Y: y,
		Z: z,
	}
}

// NewVector3Zero constructs a new Vector3 instance with all elements set to zero.
func NewVector3Zero() Vector3 {
	return Vector3{}
}
