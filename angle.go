package vector

import (
	"math"
)

// uses identity
// cos theta = a•b/(||a||*||b||)
func AngleBetween(a, b Vec3) float64 {
	return math.Acos(a.Dot(b) / (a.Magn() * b.Magn()))
}

// Scales up such that the angle
// between dir + result and dir is equal to angle
func AngleVector(dir, up Vec3, angle float64) *Vec3 {
	return up.SMul(math.Tan(angle) * dir.Magn())
}
