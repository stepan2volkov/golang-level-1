package main

import "testing"

func BenchmarkFibonacciWithoutCache10(b *testing.B) {
	// calculate fibonacci numbers without cache b.N times
	for n := 0; n < b.N; n++ {
		f := New(false)
		f.Calculate(10)
	}
}
func BenchmarkFibonacciCache10(b *testing.B) {
	// calculate fibonacci numbers with cache b.N times
	for n := 0; n < b.N; n++ {
		f := New(true)
		f.Calculate(10)
	}
}
