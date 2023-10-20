package example

import (
	"fmt"
	"github.com/lirongyangtao/mygo/base"
)

func RbTreeAdd() {
	arr := []int{
		55, 87, 56, 74, 96, 22, 62, 20, 70, 68, 90, 50,
	}
	rbTree := base.NewRbTree(base.CmInt)
	for _, v := range arr {
		rbTree.Add(v)
	}
	rbTree.TreePrint()
	for _, v := range arr {
		rbTree.Remove(v)
		fmt.Println("======================================")
		rbTree.TreePrint()
	}
}
