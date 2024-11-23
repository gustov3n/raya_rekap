package main

import "testing"

func BenchmarkMyFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Code to benchmark, e.g., calling a function from main.go
		main()
	}
}
