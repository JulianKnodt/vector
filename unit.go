package vector

import (
	"math"
)

func Unit(dir Vec3) *Vec3 {
	magn := math.Sqrt(Dot(dir, dir))
	return &Vec3{
		dir[0] / magn,
		dir[1] / magn,
		dir[2] / magn,
	}
}

func UnitSet(dir *Vec3) *Vec3 {
	magn := math.Sqrt(Dot(*dir, *dir))
	if magn == 0 {
		return dir
	}
	dir[0] /= magn
	dir[1] /= magn
	dir[2] /= magn
	return dir
}

func (v Vec3) Unit() *Vec3 {
	return Unit(v)
}

func (v *Vec3) UnitSet() *Vec3 {
	return UnitSet(v)
}
