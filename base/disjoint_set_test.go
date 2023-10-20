package base

import (
	"math/rand"
	"testing"
)

func disjointSet(count int, s string) {
	set := NewDisjointIntSet(count)

	switch s {
	case "findUnion":
		for i := 0; i < count; i++ {
			set.QuickFind_Union(int(rand.Float64()*float64(count)), int(rand.Float64()*float64(count)))
		}
	case "Union":
		for i := 0; i < count; i++ {
			set.QuickUnion_Union(int(rand.Float64()*float64(count)), int(rand.Float64()*float64(count)))
		}
	case "UnionOnlySize":
		for i := 0; i < count; i++ {
			set.QuickUnion_UnionPathSize(int(rand.Float64()*float64(count)), int(rand.Float64()*float64(count)))
		}
	case "UnionOnlyRank":
		for i := 0; i < count; i++ {
			set.QuickUnion_UnionPathRank(int(rand.Float64()*float64(count)), int(rand.Float64()*float64(count)))
		}
	case "PathCompress":
		for i := 0; i < count; i++ {
			set.QuickUnion_UnionPathCompress(int(rand.Float64()*float64(count)), int(rand.Float64()*float64(count)))
		}
	case "PathSplit":
		for i := 0; i < count; i++ {
			set.QuickUnion_UnionPathSplit(int(rand.Float64()*float64(count)), int(rand.Float64()*float64(count)))
		}
	case "PathHalving":
		for i := 0; i < count; i++ {
			set.QuickUnion_UnionPathHalving(int(rand.Float64()*float64(count)), int(rand.Float64()*float64(count)))
		}
	}

	for i := 0; i < count; i++ {
		set.IsSame(int(rand.Float64()*float64(count)), int(rand.Float64()*float64(count)))
	}
}

func BenchmarkDisjointSet_FindUnion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		disjointSet(10000, "findUnion")
	}
}
func BenchmarkDisjointSet_Union(b *testing.B) {
	for n := 0; n < b.N; n++ {
		disjointSet(10000, "Union")
	}
}
func BenchmarkDisjointSet_UnionOnlyRank(b *testing.B) {
	for n := 0; n < b.N; n++ {
		disjointSet(10000, "UnionOnlyRank")
	}
}
func BenchmarkDisjointSet_UnionOnlySize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		disjointSet(10000, "UnionOnlySize")
	}
}
func BenchmarkDisjointSet_PathCompress(b *testing.B) {
	for n := 0; n < b.N; n++ {
		disjointSet(10000, "PathCompress")
	}
}
func BenchmarkDisjointSet_PathSplit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		disjointSet(10000, "PathSplit")
	}
}

func BenchmarkDisjointSet_PathHalving(b *testing.B) {
	for n := 0; n < b.N; n++ {
		disjointSet(10000, "PathHalving")
	}
}

//go test -bench="^BenchmarkDisjointSet" .
//BenchmarkDisjointSet_FindUnion-16                     28          39115268 ns/op
//BenchmarkDisjointSet_Union-16                         68          19775535 ns/op
//BenchmarkDisjointSet_UnionOnlyRank-16                 40          26639280 ns/op
//BenchmarkDisjointSet_UnionOnlySize-16                 21          51213933 ns/op
//BenchmarkDisjointSet_PathCompress-16                 492           2451634 ns/op
//BenchmarkDisjointSet_PathSplit-16                    496           2384593 ns/op
//BenchmarkDisjointSet_PathHalving-16                  505           2354068 ns/op
