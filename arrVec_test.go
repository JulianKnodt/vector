package vector

import (
	"testing"
)

func TestShift(t *testing.T) {
	a := []Vec3{Vec3{0, 0, 0}}
	a = Shift(a, 1, 1, 1)
	if a[0][0] != 1 {
		t.Error("Did not shift correctly")
	}
}

func BenchmarkNormalNoCheck(b *testing.B) {
	x, y, z := RandomVector(), RandomVector(), RandomVector()
	for i := 0; i < b.N; i++ {
		normalNoCheck(x, y, z)
	}
}

func BenchmarkCoplanar(b *testing.B) {
	items := []Vec3{RandomVector(), RandomVector(), RandomVector(), RandomVector()}

	for i := 0; i < b.N; i++ {
		Coplanar(items)
	}
}

func BenchmarkIntersectsTriangle(b *testing.B) {
	t0 := Vec3{5, 0, 0}
	t1 := Vec3{5, 10, 0}
	t2 := Vec3{5, 5, 10}

	origin := Vec3{0, 3, 0}
	dir := Vec3{5, 2, 5}

	for i := 0; i < b.N; i++ {
		IntersectsTriangle(t0, t1, t2, origin, dir)
	}
}

func BenchmarkIntersectsTriangle2(b *testing.B) {
	t0 := Vec3{5, 0, 0}
	t1 := Vec3{5, 10, 0}
	t2 := Vec3{5, 5, 10}

	origin := Vec3{0, 3, 0}
	dir := Vec3{5, 2, 5}

	for i := 0; i < b.N; i++ {
		IntersectsTriangle2(t0, t1, t2, origin, dir)
	}
}
