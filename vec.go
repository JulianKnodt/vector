package vector

import (
	"image/color"
	"math"
)

// Represents a vector in 3-space
type Vec3 [3]float64

// The origin of all things vector.
// Provided for convenience
var Origin Vec3 = Vec3{0, 0, 0}

// Returns a pointer to a new vector at the origin
func NewOrigin() *Vec3 {
	return &Vec3{0, 0, 0}
}

// Returns a vector with all three components equal to a
func NewVec(a float64) *Vec3 {
	return &Vec3{a, a, a}
}

// Returns X component of v, Can also use [0]
func (v Vec3) X() float64 {
	return v[0]
}

// Returns Y component of v, Can also use [1]
func (v Vec3) Y() float64 {
	return v[1]
}

// Returns Z component of v, Can also use [2]
func (v Vec3) Z() float64 {
	return v[2]
}

// Checks if all three components of a and b are equal
func Equal(a, b Vec3) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2]
}

// Returns whether a and b are in the same direction
func Proportional(a, b Vec3) bool {
	firstRel := a[0] / b[0]
	return firstRel == a[1]/b[1] && firstRel == a[2]/b[2]
}

// Returns the negative of each component of a and returns a
func InvSet(a *Vec3) *Vec3 {
	a[0] = -a[0]
	a[1] = -a[1]
	a[2] = -a[2]
	return a
}

// Returns the squared magnitude of a
func SqrMagn(a Vec3) float64 {
	return a[0]*a[0] + a[1]*a[1] + a[2]*a[2]
}

func (v Vec3) SqrMagn() float64 {
	return SqrMagn(v)
}

// Returns the magnitude of a
func Magn(a Vec3) float64 {
	return math.Sqrt(SqrMagn(a))
}

// Returns the magnitude of a
func (a Vec3) Magn() float64 {
	return Magn(a)
}

// Returns a new vector with inverted components of a
func Inv(a Vec3) *Vec3 {
	return &Vec3{-a[0], -a[1], -a[2]}
}

func (v Vec3) Inv() *Vec3 {
	return Inv(v)
}

func sqr(a float64) float64 {
	return a * a
}

// Returns radians between a and b
func Theta(a, b Vec3) float64 {
	return math.Acos(Dot(a, b) / (Magn(a) * Magn(b)))
}

// Returns a new vector of the hadamard(piecewise multiplication)
func Hadamard(a, b Vec3) *Vec3 {
	return &Vec3{
		a[0] * b[0],
		a[1] * b[1],
		a[2] * b[2],
	}
}

// Returns the squared distance between a and b
func DistSquared(a, b Vec3) float64 {
	return sqr(a[0]-b[0]) + sqr(a[1]-b[1]) + sqr(a[2]-b[2])
}

func ToRGBA(a Vec3) color.RGBA {
	return color.RGBA{uint8(a[0]), uint8(a[1]), uint8(a[2]), 255}
}

// runs op on every component of in and returns a new vector
func Op(in Vec3, op func(float64) float64) *Vec3 {
	return &Vec3{
		op(in[0]),
		op(in[1]),
		op(in[2]),
	}
}
