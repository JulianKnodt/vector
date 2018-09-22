package vector

// Reflects v over "across" which is expected to be normalized
func (v *Vec3) Reflection(across Vec3) *Vec3 {
	return v.Sub(*across.SMul(2.0 * v.Dot(across)))
}
