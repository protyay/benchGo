
package main

import (
	"testing"
)

func BenchmarkSliceAppend(b *testing.B) {
	// Test different initial capacities
	initialSizes := []int{0, 1000, 10000}
	
	for _, size := range initialSizes {
		b.Run("InitialSize_"+string(rune(size)), func(b *testing.B) {
			slice := make([]int, 0, size)
			b.ResetTimer()
			
			for i := 0; i < b.N; i++ {
				slice = append(slice, i)
			}
		})
	}
}

func BenchmarkSliceWithPreallocation(b *testing.B) {
	b.Run("Preallocated", func(b *testing.B) {
		slice := make([]int, 0, b.N)
		b.ResetTimer()
		
		for i := 0; i < b.N; i++ {
			slice = append(slice, i)
		}
	})
}
