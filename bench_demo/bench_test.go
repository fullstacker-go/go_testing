package bench_demo

import "testing"

func BenchmarkChangeName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChangeName(3)
	}
}
