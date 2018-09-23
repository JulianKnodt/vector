package vector

func TriangleArea(a, b, c Vec3) float64 {
	return a.Sub(b).CrossSet(*a.Sub(c)).Magn() / 2
}
