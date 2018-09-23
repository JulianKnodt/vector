package vector

import (
	"testing"
)

func BenchmarkCross(b *testing.B) {
	a := *RandomVector()
	c := *RandomVector()
	for i := 0; i < b.N; i++ {
		Cross(a, c)
	}
}
