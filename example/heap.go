package example

import (
	"fmt"
	"github.com/lirongyangtao/mygo/base"
)

func ExampleHeap() {
	heap := base.NewQuadHeap(base.CmInt)
	heap.Add(199)
	heap.Add(99)
	heap.Add(3000)
	heap.Add(100)
	heap.Add(11)
	heap.Add(1000)
	for heap.Len() != 0 {
		fmt.Println(heap.Pop(), heap.Len(), heap.PeekTop())
	}
}
