package vector

import (
	"math"
	"math/rand"
)

type VecN []float64

func (a *VecN) Add(b *VecN) *VecN {
	out := make(VecN, len(*a))
	for i, v := range *a {
		out[i] = v + (*b)[i]
	}
	return &out
}

func (a *VecN) Sub(b *VecN) *VecN {
	out := make(VecN, len(*a))
	for i, v := range *a {
		out[i] = v - (*b)[i]
	}
	return &out
}

func (a *VecN) SMul(k float64) *VecN {
	out := make(VecN, len(*a))
	for i, v := range *a {
		out[i] = v * k
	}
	return &out
}

func (a *VecN) SqrDist(b *VecN) float64 {
	sum := 0.0
	for i, v := range *a {
		sum += sqr(v - (*b)[i])
	}
	return sum
}

func (a *VecN) Round() *VecN {
	out := make(VecN, len(*a))
	for i, v := range *a {
		out[i] = math.Round(v)
	}
	return &out
}

func (a *VecN) ToIntSlice() []int {
	out := make([]int, len(*a))
	for i, v := range *a {
		out[i] = int(v)
	}
	return out
}

func (a *VecN) RandomWalk() {
	for i := range *a {
		if rand.Float64() >= 0.5 {
			(*a)[i] += 1
		} else {
			(*a)[i] -= 1
		}
	}
}

func (a *VecN) SqrMagn() float64 {
	sum := 0.0
	for _, v := range *a {
		sum += sqr(v)
	}
	return sum
}

func (a *VecN) Unit() *VecN {
	out := make(VecN, len(*a))
	magn := math.Sqrt(a.SqrMagn())
	for i, v := range *a {
		out[i] = v / magn
	}
	return &out
}

func ZeroVecOf(dim int) *VecN {
	out := make(VecN, dim)
	return &out
}

func VecOfAll(dim int, val float64) *VecN {
	out := make(VecN, dim)
	for i := range out {
		out[i] = val
	}
	return &out
}
