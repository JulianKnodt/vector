package vector

func Dot(a, b Vec3) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func (a Vec3) Dot(b Vec3) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}
