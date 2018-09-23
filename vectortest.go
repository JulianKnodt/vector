package vector

// Some testing utils for vectors

import (
	"math/rand"
)

// Returns a random vector in ([0,1] * [0,1] * [0,1])
func RandomVector() *Vec3 {
	return &Vec3{
		rand.Float64(),
		rand.Float64(),
		rand.Float64(),
	}
}

// Returns 3 RandomVectors, for testing triangles
func RandomTriple() [3]Vec3 {
	return [3]Vec3{
		*RandomVector(),
		*RandomVector(),
		*RandomVector(),
	}
}
