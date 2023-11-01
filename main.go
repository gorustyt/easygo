package main

import (
	"fmt"
	"github.com/lirongyangtao/mygo/apply"
	"time"
)

func main() {
	//tree := base.NewRbTree[int]()
	//cb := func(node *base.BinaryTreeNode[int]) string {
	//	if node.IsRed() {
	//		return fmt.Sprintf("%v(red)", node.GetElement())
	//	}
	//	return fmt.Sprintf("%v(black)", node.GetElement())
	//}
	//arr := []int{99, 100, 7, 5, 4, 3, 2, 1}
	//for _, v := range arr {
	//	tree.Add(v)
	//	base.PrintBinaryTree(tree.Root, cb)
	//}
	//
	//for i := len(arr) - 1; i >= 0; i-- {
	//	tree.Remove(arr[i])
	//	base.PrintBinaryTree(tree.Root, cb)
	//}
	//time.Sleep(1 * time.Second)

	wheel := apply.NewAsyncTimeWheel()
	id := wheel.Schedule(1*time.Second, 1*time.Second, func(ts time.Time) {
		fmt.Println("===", time.Now())
	})
	time.AfterFunc(5100*time.Millisecond, func() {
		fmt.Println("====stop=====", time.Now())
		wheel.Remove(id)
	})
	time.Sleep(1000 * time.Second)
}
