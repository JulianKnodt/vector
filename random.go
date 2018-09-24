package vector

import (
	"math/rand"
)

func Random(rs *rand.Rand) *Vec3 {
	if rs == nil {
		return &Vec3{
			rand.Float64(),
			rand.Float64(),
			rand.Float64(),
		}
	}
	return &Vec3{rs.Float64(), rs.Float64(), rs.Float64()}
}
