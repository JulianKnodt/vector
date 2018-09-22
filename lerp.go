package vector

// Linearly interpolates "from" to "to" w/ proportion [0,1]
func Lerp(from, to Vec3, proportion float64) *Vec3 {
	return AddSet(Sub(to, from).SMulSet(proportion), from)
}

// Linearly interpolates "from" to "to" w/ proportion [0,1]
func (from Vec3) Lerp(to Vec3, proportion float64) *Vec3 {
	return Lerp(from, to, proportion)
}

func LerpInv(v Vec3, p float64) *Vec3 {
	prop := (1 - 2*p)
	return v.SMul(prop)
}

func (v Vec3) LerpInv(p float64) *Vec3 {
	return LerpInv(v, p)
}

// Linearly Interpolates v with the inverse of itself with propoprtion [0,1]
func LerpInvSet(v *Vec3, p float64) *Vec3 {
	return v.SMulSet(1 - 2*p)
}

func (v *Vec3) LerpInvSet(p float64) *Vec3 {
	return LerpInvSet(v, p)
}
