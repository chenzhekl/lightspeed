// +build float64

package math

import "math"

// Float is the IEEE-754 floating-point number.
// It' either 32-bit or 64-bit depending on the build flag.
type Float float64

// MaxFloat is the maximum number a Float can represent.
const MaxFloat = math.MaxFloat64
