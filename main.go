package main

import (
	"fmt"
	"github.com/lirongyangtao/mygo/base"
)

func main() {
	tree := base.NewBinaryTree[int]()
	tree.Add(99)
	tree.Add(100)
	tree.Add(7)
	tree.Add(6)
	tree.Add(5)
	tree.Print(nil)
	tree.Root = base.RotateRight(tree.Root)
	tree.Print(func(node *base.BinaryTreeNode[int]) string {
		return fmt.Sprintf("%v", node.GetElement())
	})
}
