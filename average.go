package vector

func Average(vecs []Vec3) *Vec3 {
	out := &Vec3{0, 0, 0}
	num := float64(len(vecs))
	for _, v := range vecs {
		out.AddSet(*v.SMul(1 / num))
	}
	return out
}
