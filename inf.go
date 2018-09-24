package vector

import "math"

// Functions which returns the infinite vector
// mainly to be used in comparisons
func Inf(sign int) *Vec3 {
	return &Vec3{math.Inf(sign), math.Inf(sign), math.Inf(sign)}
}
