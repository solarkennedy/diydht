package main

import "testing"

func BenchmarkStockInserts(b *testing.B) {
	a := make(map[int]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a[i] = 42
	}
}

func BenchmarkStockGets(b *testing.B) {
	a := make(map[int]int)
	for i := 0; i < b.N; i++ {
		a[i] = 42
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = a[i]
	}
}

func BenchmarkStockDeletes(b *testing.B) {
	a := make(map[int]int)
	for i := 0; i < b.N; i++ {
		a[i] = 42
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		delete(a, i)
	}
}
