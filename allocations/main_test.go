package main

import "testing"

const creations = 20_000_000

func TestCopyIt(t *testing.T) {
	for i := 0; i < 20_000_000; i++ {
		_ = CreateCopy()
	}
}

func TestPointerIt(t *testing.T) {
	for i := 0; i < creations; i++ {
		_ = CreatePointer()
	}
}

func BenchmarkStackIt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main1()
	}
}

func BenchmarkStackIt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main2()
	}
}

func BenchmarkStackIt3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main3()
	}
}

func BenchmarkCopyIt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copyIt()
	}
}

func BenchmarkPointerIt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pointerIt()
	}
}
