package main

import (
	"testing"
)

func BenchmarkSimulateConcurrentProcessesWithNoSf(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		simulateConcurrentProcessesWithNoSf()
	}
}

func BenchmarkSimulateConcurrentProcessesWithSf(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		simulateConcurrentProcessesWithSf()
	}
}
