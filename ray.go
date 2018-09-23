package vector

import (
	"math"
)

// Represents a ray starting from origin
// going in direction dir
type Ray struct {
	Origin    Vec3
	Direction Vec3
}

// Returns a new ray with origin and direction(which is converted to a unit vector)
func NewRay(origin, dir Vec3) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: *Unit(dir),
	}
}

// Returns the direction of the ray relative to the x-axis
func (r Ray) Theta() float64 {
	return math.Acos(r.Direction[2])
}

func (r Ray) Phi() float64 {
	return math.Atan(r.Direction[1] / r.Direction[0])
}

// Returns a constant of 1 since a ray has a unit direction
func (r Ray) Rho() float64 {
	return 1
}

// Returns a new vector of the ray at point t from the origin
func (r Ray) At(t float64) *Vec3 {
	return r.Direction.SMul(t).AddSet(r.Origin)
}

func NewRayFrom(from, to Vec3) *Ray {
	return NewRay(from, *from.Sub(to))
}
