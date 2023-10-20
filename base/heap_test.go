package base

import "testing"

func NewHeap(t int) Heap {
	switch t {
	case 4:
		return NewQuadHeap(CmInt)
	}
	return NewBinaryHeap(CmInt)
}
func testHeap(int int, num int) {
	heap := NewHeap(int)
	for i := 0; i < num; i++ {
		heap.Add(i)
	}
	for i := 0; i < num; i++ {
		heap.Replace(i)
	}
	for i := 0; i < num; i++ {
		heap.Remove(i)
	}
}
func BenchmarkHeapBinary(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testHeap(2, 100000)
	}
}

func BenchmarkHeapFour(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testHeap(4, 100000)
	}
}

//go test -bench="^BenchmarkHeap" .
//BenchmarkHeapBinary-16                26          39099192 ns/op
//BenchmarkHeapFour-16                  16          66876188 ns/op
//PASS
//ok      awesomeProject5/base    2.523s
//ok      awesomeProject5/base    2.578s
