package main

import "testing"

// To run: go test -bench . -count 10

type MyStruct struct {
	A int
	B int
}

func BenchmarkSlicePointers(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var slice []*MyStruct
		for j := 0; j < 100; j++ {
			slice = append(slice, &MyStruct{A: j, B: j + 1})
		}
	}
}

func BenchmarkSliceNoPointers(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var slice []MyStruct
		for j := 0; j < 100; j++ {
			slice = append(slice, MyStruct{A: j, B: j + 1})
		}
	}
}
