package vector

func SMul(k float64, v Vec3) *Vec3 {
	return &Vec3{
		k * v[0],
		k * v[1],
		k * v[2],
	}
}

func SMulSet(k float64, dst *Vec3) {
	dst[0] = k * dst[0]
	dst[1] = k * dst[1]
	dst[2] = k * dst[2]
}

func (v Vec3) SMul(k float64) *Vec3 {
	return SMul(k, v)
}

func (v *Vec3) SMulSet(k float64) *Vec3 {
	SMulSet(k, v)
	return v
}
