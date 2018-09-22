package vector

import (
	"testing"
)

func TestLerpInvSet(t *testing.T) {
	a := Vec3{0.6, 1.3, 3.6}
	if *LerpInv(a, 0) != a {
		t.Error("Lerp Inv failed for 0")
	}
	if *LerpInv(a, 1) != *a.Inv() {
		t.Error("Lerp Inv failed for 1")
	}
	if *LerpInv(a, 0.5) != Origin {
		t.Error("Lerp Inv failed for 0.5")
	}
}
