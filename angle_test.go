package vector

import (
	"math"
	"testing"
)

func TestAngleVector(t *testing.T) {
	up := Vec3{0, 0, 1}
	dir := Vec3{0, 1, 0}

	angle := math.Pi / 6.0 // 60 degrees

	av := AngleVector(dir, up, math.Pi/6.0)
	if ab := AngleBetween(*av.Add(dir), dir); math.Abs(ab-angle) > 0.0000000001 {
		t.Error("Expected angles to be equal", ab, angle, math.Abs(ab-angle))
	}
}
