package vector

// Checks whether this vector is the zero vector {0,0,0}
func (v Vec3) IsZero() bool {
	return v[0] == 0 && v[1] == 0 && v[2] == 0
}
