package vector

import (
	"math"
)

// Maxes components of v and o
// and returns a new vector
func (v Vec3) Max(o Vec3) *Vec3 {
	return &Vec3{
		math.Max(v[0], o[0]),
		math.Max(v[1], o[1]),
		math.Max(v[2], o[2]),
	}
}

// Sets v to be the maximum of v and o
func (v *Vec3) MaxSet(o Vec3) *Vec3 {
	v[0] = math.Max(v[0], o[0])
	v[1] = math.Max(v[1], o[1])
	v[2] = math.Max(v[2], o[2])
	return v
}

// Maxes components of v and o
// And returns a new vector
func Max(v, o Vec3) *Vec3 {
	return Max(o, v)
}

// Mins components of v and o
// And returns a new vector
func (v Vec3) Min(o Vec3) *Vec3 {
	return Min(v, o)
}

// Sets v to be the minimum of v and o
func (v *Vec3) MinSet(o Vec3) *Vec3 {
	v[0] = math.Min(v[0], o[0])
	v[1] = math.Min(v[1], o[1])
	v[2] = math.Min(v[2], o[2])
	return v
}

// Mins components of v and o
// And returns a new vector
func Min(v, o Vec3) *Vec3 {
	return &Vec3{
		math.Min(v[0], o[0]),
		math.Min(v[1], o[1]),
		math.Min(v[2], o[2]),
	}
}
