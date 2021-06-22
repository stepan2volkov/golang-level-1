package main

import "testing"

func BenchmarkFibonacciWithouCache10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		f := New(false)
		f.Calculate(10)
	}
}
func BenchmarkFibonacciCache10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		f := New(true)
		f.Calculate(10)
	}
}
