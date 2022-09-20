package fastrand

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// fmt.Printf("%d\n", Generate())
		Generate(0)
	}
}

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		// fmt.Printf("%d\n", rand.Int())
		rand.Int()
	}
}
