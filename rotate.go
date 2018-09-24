package vector

import (
	"math"
)

// https://stackoverflow.com/questions/14607640/rotating-a-vector-in-3d-space

// Rotate Vec Some portion around the z axis
func RotateZ(v Vec3, t float64) *Vec3 {
	cosT := math.Cos(t)
	sinT := math.Sin(t)
	return &Vec3{
		v[0]*cosT - v[1]*sinT,
		v[0]*sinT + v[1]*cosT,
		v[2],
	}
}

func (v Vec3) RotateZ(t float64) *Vec3 {
	return RotateZ(v, t)
}

// Rotate Vec Some portion around the y axis
func RotateY(v Vec3, t float64) *Vec3 {
	cosT := math.Cos(t)
	sinT := math.Sin(t)
	return &Vec3{
		v[0]*cosT + v[2]*sinT,
		v[1],
		v[2]*cosT - v[1]*sinT,
	}
}

func (v Vec3) RotateY(t float64) *Vec3 {
	return RotateY(v, t)
}

// Rotate Vec Some portion around the x axis
func RotateX(v Vec3, t float64) *Vec3 {
	cosT := math.Cos(t)
	sinT := math.Sin(t)
	return &Vec3{
		v[0],
		v[1]*cosT - v[2]*sinT,
		v[1]*sinT + v[2]*cosT,
	}
}

func (v Vec3) RotateX(t float64) *Vec3 {
	return RotateX(v, t)
}

func Rotate(vecs []Vec3, x, y, z float64) []Vec3 {
	out := make([]Vec3, len(vecs))
	for i, v := range vecs {
		out[i] = *v.RotateX(x).RotateY(y).RotateZ(z)
	}
	return out
}
