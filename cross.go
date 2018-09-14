package vector

func Cross(a, b Vec3) *Vec3 {
	return &Vec3{a[1]*b[2] - a[2]*b[1], a[2]*b[0] - a[0]*b[2], a[0]*b[1] - a[1]*b[0]}
}

func (a Vec3) Cross(b Vec3) *Vec3 {
	return Cross(a, b)
}

func CrossSet(dst *Vec3, o Vec3) *Vec3 {
	dst[0] = dst[1]*o[2] - dst[2]*o[1]
	dst[1] = dst[2]*o[0] - dst[0]*o[2]
	dst[2] = dst[0]*o[1] - dst[1]*o[0]
	return dst
}

func (dst *Vec3) CrossSet(o Vec3) *Vec3 {
	CrossSet(dst, o)
	return dst
}
