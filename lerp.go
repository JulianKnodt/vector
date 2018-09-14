package vector

func Lerp(from, to Vec3, proportion float64) *Vec3 {
	return AddSet(Sub(to, from).SMulSet(proportion), from)
}

func (from Vec3) Lerp(to Vec3, proportion float64) *Vec3 {
	return Lerp(from, to, proportion)
}
