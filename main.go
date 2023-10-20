package main

import (
	"fmt"
	"github.com/lirongyangtao/mygo/base"
)

func main() {
	tree := base.NewBinaryTree[int]()

	tree.Add(5)
	tree.Add(6)
	tree.Add(7)
	tree.Add(8)
	tree.Add(9)
	tree.Add(10)
	tree.Print(func(node *base.BinaryTreeNode[int]) string {
		return fmt.Sprintf("%v", node.GetElement())
	})
}
