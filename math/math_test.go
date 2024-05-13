package main

import (
	"math"
	"testing"
	"testing/quick"
)

func BenchmarkSqrt(b *testing.B) {
	f := func(x float64) bool {
		y := math.Sqrt(x)
		return y*y <= x && x < (y+1)*(y+1)
	}
	if err := quick.Check(f, nil); err != nil {
		b.Errorf("Sqrt is not valid: %v", err)
	}
}
