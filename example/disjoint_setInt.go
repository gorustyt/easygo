package example

import (
	"fmt"
	"github.com/lirongyangtao/mygo/base"
)

func DisjointSetIntExample() {
	set := base.NewDisjointIntSet(100)
	set.QuickUnion_UnionPathHalving(2, 3)
	set.QuickUnion_UnionPathHalving(3, 4)
	set.QuickUnion_UnionPathHalving(4, 5)
	set.QuickUnion_UnionPathHalving(5, 6)
	set.QuickUnion_UnionPathHalving(6, 7)

	set.QuickUnion_UnionPathHalving(10, 11)
	set.QuickUnion_UnionPathHalving(12, 13)
	fmt.Println(set.IsSame(12, 13))
	fmt.Println(set.IsSame(2, 3))
	fmt.Println(set.IsSame(2, 7))
	fmt.Println(set.IsSame(2, 10))
}
func DisjointSetExample() {
	set := base.NewDisjointSet()
	set.Union(2, 3)
	set.Union(3, 4)
	set.Union(4, 5)
	set.Union(5, 6)
	set.Union(6, 7)

	set.Union(10, 11)
	set.Union(12, 13)
	fmt.Println(set.IsSame(12, 13))
	fmt.Println(set.IsSame(2, 3))
	fmt.Println(set.IsSame(2, 7))
	fmt.Println(set.IsSame(2, 10))
	fmt.Println(set.IsSame(2, 100))
}
