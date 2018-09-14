package vector

import (
	"testing"
)

func TestAdd(t *testing.T) {
	v0 := NewVec(1)
	if v0.Add(*NewVec(3))[0] != 4 {
		t.Fail()
	}
}

func TestAddSet(t *testing.T) {
	v0 := NewVec(1)
	v0.AddSet(*NewVec(3))
	if v0[0] != 4 {
		t.Fail()
	}
}
