package vector

// Returns a new copy of a vector
func (v Vec3) Copy() *Vec3 {
	return &Vec3{v[0], v[1], v[2]}
}
