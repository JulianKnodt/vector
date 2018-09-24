package vector

// Shifts and sets v by x, y, z
func (v *Vec3) ShiftSet(x, y, z float64) *Vec3 {
	v[0] += x
	v[1] += y
	v[2] += z
	return v
}
