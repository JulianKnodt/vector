package vector

import (
	"math"
)

func Pow(v Vec3, k float64) *Vec3 {
	return &Vec3{
		math.Pow(v[0], k),
		math.Pow(v[1], k),
		math.Pow(v[2], k),
	}
}

func (v Vec3) Pow(k float64) *Vec3 {
	return Pow(v, k)
}

func PowSet(v *Vec3, k float64) {
	v[0] = math.Pow(v[0], k)
	v[1] = math.Pow(v[1], k)
	v[2] = math.Pow(v[2], k)
}

func (v *Vec3) PowSet(k float64) *Vec3 {
	PowSet(v, k)
	return v
}
