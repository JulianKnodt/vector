package vector

import (
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	b.Run("Add functional", func(b *testing.B) {
		out := Vec3{0, 0, 0}
		result := Vec3{1, 1, 1}
		for i := 0; i < b.N; i++ {
			result = *Add(out, result)
		}
	})

	b.Run("Add pointer", func(b *testing.B) {
		out := &Vec3{0, 0, 0}
		result := Vec3{1, 1, 1}
		for i := 0; i < b.N; i++ {
			out = out.AddSet(result)
		}
	})
}

func BenchmarkSub(b *testing.B) {
	b.Run("Sub functional", func(b *testing.B) {
		out := Vec3{0, 0, 0}
		result := Vec3{1, 1, 1}
		for i := 0; i < b.N; i++ {
			result = *Sub(out, result)
		}
	})

	b.Run("Sub pointer", func(b *testing.B) {
		out := &Vec3{0, 0, 0}
		result := Vec3{1, 1, 1}
		for i := 0; i < b.N; i++ {
			out = out.SubSet(result)
		}
	})
}

func BenchmarkLerp(b *testing.B) {
	b.Run("Lerp functional", func(b *testing.B) {
		out := Vec3{0, 0, 0}
		result := Vec3{1, 1, 1}
		for i := 0; i < b.N; i++ {
			result = *Lerp(out, result, 0.3)
		}
	})

	b.Run("Lerp pointer", func(b *testing.B) {
		out := Vec3{0, 0, 0}
		result := Vec3{1, 1, 1}
		for i := 0; i < b.N; i++ {
			out = *out.Lerp(result, 0.3)
		}
	})
}

func BenchmarkDot(b *testing.B) {
	b.Run("Dot functional", func(b *testing.B) {
		v1 := RandomVector()
		v2 := RandomVector()
		for i := 0; i < b.N; i++ {
			Dot(v1, v2)
		}
	})

	b.Run("Dot receiver", func(b *testing.B) {
		v1 := RandomVector()
		v2 := RandomVector()
		for i := 0; i < b.N; i++ {
			v1.Dot(v2)
		}
	})
}
