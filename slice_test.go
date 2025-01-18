package main

import (
	"testing"
)

func BenchmarkStackAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MakePerson("John", "Doe", 13)
	}
}

func BenchmarkHeapAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MakePersonPointer("John", "Doe", 14)
	}
}
