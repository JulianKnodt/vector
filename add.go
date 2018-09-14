package vector

// Adds a and b and returns a new vector
func Add(a, b Vec3) *Vec3 {
	return &Vec3{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
	}
}

// Adds dst and src, and sets dst as the result
func AddSet(dst *Vec3, src Vec3) *Vec3 {
	dst[0] = dst[0] + src[0]
	dst[1] = dst[1] + src[1]
	dst[2] = dst[2] + src[2]
	return dst
}

// Adds v and o and returns a new vector
func (v Vec3) Add(o Vec3) *Vec3 {
	return Add(v, o)
}

// Adds v and o, and sets v as the result
func (v *Vec3) AddSet(o Vec3) *Vec3 {
	AddSet(v, o)
	return v
}
