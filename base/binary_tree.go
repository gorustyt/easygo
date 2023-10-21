package base

import (
	"cmp"
	"fmt"
	"math"
	"strings"
	"unicode"
)

type CmpT interface {
	any
	cmp.Ordered
}

// 二叉树
type BinaryTree[T CmpT] struct {
	size int
	Root *BinaryTreeNode[T]
}

type BinaryTreeNode[T CmpT] struct {
	parent  *BinaryTreeNode[T] //父节点
	left    *BinaryTreeNode[T] //左子节点
	right   *BinaryTreeNode[T] //右子节点
	element T                  //元素
	Index   int                //在二叉树中索引,打印辅助
}

func (n *BinaryTreeNode[T]) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

// 是否有2个孩子
func (n *BinaryTreeNode[T]) HasTwoChildren() bool {
	if n == nil {
		return false
	}
	return n.left != nil && n.right != nil
}
func (n *BinaryTreeNode[T]) GetElement() interface{} {
	if n == nil {
		return nil
	}
	return n.element
}
func NewBinaryTreeNode[T CmpT](element T, parent *BinaryTreeNode[T]) *BinaryTreeNode[T] {
	return &BinaryTreeNode[T]{
		parent:  parent,
		element: element,
	}
}

func NewBinaryTree[T CmpT]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

func (t *BinaryTree[T]) Len() int {
	return t.size
}

// 树的高度
func (t *BinaryTree[T]) Height() int {
	return t.heightByIter(t.Root)
}

// 递归实现
func (t *BinaryTree[T]) height(node *BinaryTreeNode[T]) int {
	if node == nil {
		return 0
	}
	return 1 + int(math.Max(float64(t.height(node.left)), float64(t.height(node.right))))
}

// 迭代实现
func (t *BinaryTree[T]) heightByIter(node *BinaryTreeNode[T]) (height int) {
	if node == nil {
		return 0
	}
	que := NewSimpleQueue()
	que.Offer(node)
	size := que.Len()
	for que.Len() != 0 {
		node := que.Poll().(*BinaryTreeNode[T])
		size--
		if node.left != nil {
			que.Offer(node.left)
		}
		if node.right != nil {
			que.Offer(node.right)
		}
		if size == 0 {
			height++
			size = que.Len()
		}
	}
	return
}

// 获取二叉树的每一层节点索引
func (t *BinaryTree[T]) getPrintList(cb func(node *BinaryTreeNode[T]) string) (res [][]string, maxLen int) {
	node := t.Root
	if node == nil {
		return
	}
	if cb == nil {
		cb = func(node *BinaryTreeNode[T]) string {
			return fmt.Sprintf("%v", node.element)
		}
	}
	que := NewSimpleQueue()
	que.Offer(node)
	size := 1
	height := 0
	arr := make([]string, int(math.Pow(2, float64(height))))
	//获取元素最大字符长度
	for que.Len() != 0 {
		node = que.Poll().(*BinaryTreeNode[T])
		index := node.Index - (int(math.Pow(2, float64(height))) - 1)
		arr[index] = cb(node)
		length := t.getPrintCount(arr[index])
		if length > maxLen {
			maxLen = length
		}
		size--
		if node.left != nil {
			node.left.Index = 2*node.Index + 1
			que.Offer(node.left)
		}
		if node.right != nil {
			node.right.Index = 2*node.Index + 2
			que.Offer(node.right)
		}
		if size == 0 {
			res = append(res, arr)
			height++
			arr = make([]string, int(math.Pow(2, float64(height))))
			size = que.Len()
		}
	}
	return
}

// //树状字符
// //"└" "─" "┌─────────┴─────────┐" "┬"
// //"├"
// //"─┴─"
func (t *BinaryTree[T]) Print(cb func(node *BinaryTreeNode[T]) string, compress ...bool) {
	list, maxLen := t.getPrintList(cb)
	if maxLen%2 == 0 {
		maxLen++
	}
	//填充长度
	for i, arr := range list {
		for j, v := range arr {
			length := t.getPrintCount(v)
			if len(v) == maxLen {
				continue
			}
			delta := maxLen - length
			leftCount := delta / 2
			rightCount := delta - leftCount
			list[i][j] = fmt.Sprintf("%v%v%v",
				strings.Repeat(space, leftCount),
				v,
				strings.Repeat(space, rightCount),
			)
		}
	}
	t.print(list, maxLen)
}

var (
	space      = " "
	_preSpace  = "┌"
	_preSpace1 = "┴"
	_preSpace2 = "┘"
	_preSpace3 = "└"
	_inSpace   = "─"
	_postSpace = "┐"

	preSpace  = "/"
	preSpace1 = "^"
	preSpace2 = "$"
	preSpace3 = "@"
	inSpace   = "~"
	postSpace = "\\"
)

func (t *BinaryTree[T]) print(list [][]string, maxLen int, compress ...bool) {
	p := 0
	values := make([]string, len(list))
	lines := make([]string, len(list))
	for i := len(list) - 1; i >= 0; i-- {
		preValue := strings.Repeat(space, maxLen*int(math.Pow(2, float64(p))))
		preLine := strings.Repeat(space, maxLen*int(math.Pow(2, float64(p)))+maxLen/2)
		valueInterStr := strings.Repeat(space, maxLen*int(math.Pow(2, float64(p+1))-1))
		lineInterStr := strings.Repeat(space, maxLen*int(math.Pow(2, float64(p+1))-1)+maxLen/2*2)
		lineHalfInterStr := strings.Repeat(inSpace, maxLen*int(math.Pow(2, float64(p))-1)+maxLen/2+maxLen/2)
		lineHalfSpaceStr := strings.ReplaceAll(lineHalfInterStr, inSpace, space)
		valueStr := preValue
		lineStr := preLine
		for j := 0; j < len(list[i]); j += 2 {
			leftNode := list[i][j]
			valueStr += leftNode
			valueStr += valueInterStr
			if j+1 >= len(list[i]) {
				break
			}
			rightNode := list[i][j+1]
			leftEmpty := strings.TrimSpace(leftNode) == ""
			rightEmpty := strings.TrimSpace(rightNode) == ""
			//处理值打印

			valueStr += rightNode
			valueStr += valueInterStr
			//处理字符打印
			if !leftEmpty && !rightEmpty {
				lineStr += preSpace
				lineStr += lineHalfInterStr
				lineStr += preSpace1
				lineStr += lineHalfInterStr
				lineStr += postSpace
			} else if !leftEmpty && rightEmpty {
				lineStr += preSpace
				lineStr += lineHalfInterStr
				lineStr += preSpace2
				lineStr += lineHalfSpaceStr
				lineStr += space
			} else if leftEmpty && !rightEmpty {
				lineStr += space
				lineStr += lineHalfSpaceStr
				lineStr += preSpace3
				lineStr += lineHalfInterStr
				lineStr += postSpace
			} else if leftEmpty && rightEmpty {
				lineStr += space
				lineStr += lineHalfSpaceStr
				lineStr += space
				lineStr += lineHalfSpaceStr
				lineStr += space
			}
			lineStr += lineInterStr
		}
		values[i] = valueStr
		lines[i] = lineStr
		valueStr = ""
		lineStr = ""
		p++
	}
	if len(compress) <= 0 || (len(compress) > 0 && !compress[0]) {
		t.printCompress(values, lines, 2)
	}
	for index, v := range lines {
		v = strings.ReplaceAll(v, preSpace, _preSpace)
		v = strings.ReplaceAll(v, preSpace1, _preSpace1)
		v = strings.ReplaceAll(v, preSpace2, _preSpace2)
		v = strings.ReplaceAll(v, preSpace3, _preSpace3)
		v = strings.ReplaceAll(v, inSpace, _inSpace)
		v = strings.ReplaceAll(v, postSpace, _postSpace)
		lines[index] = v
	}
	for index, v := range values {
		fmt.Println(v)
		if index+1 < len(values) {
			fmt.Println(lines[index+1])
		}
	}

}

func (t *BinaryTree[T]) printCompress(values, lines []string, compressCount int) {
	i := 0
	maxLen := len(values[len(values)-1])
	beginIndex := -1
	endIndex := -1
	for {
		var (
			canCompress = true
		)
		for j1 := 0; j1 < len(values); j1++ {
			v := values[j1]
			if i >= len(v) {
				continue
			}
			v1 := v[i : i+1]
			if v1 != space && v1 != inSpace {
				canCompress = false
			}
		}

		for j2 := 1; j2 < len(lines); j2++ {
			v := lines[j2]
			if i >= len(v) {
				continue
			}
			line1 := v[i : i+1]
			if line1 != space && line1 != inSpace {
				canCompress = false
			}
		}
		if canCompress {
			if beginIndex == -1 && endIndex == -1 {
				beginIndex = i
				endIndex = i
			} else {
				endIndex = i
			}
		} else if beginIndex != -1 && endIndex != -1 {
			if endIndex-beginIndex+1 > compressCount {
				for j1 := 0; j1 < len(values); j1++ {
					v := values[j1]
					if i >= len(v) {
						goto END
					}
					values[j1] = v[:beginIndex] + v[endIndex+1-compressCount:]
				}

				for j2 := 1; j2 < len(lines); j2++ {
					v := lines[j2]
					if i >= len(v) {
						goto END
					}
					lines[j2] = v[:beginIndex] + v[endIndex+1-compressCount:]
				}
				i = beginIndex
			}
			beginIndex = -1
			endIndex = -1
			continue
		}
	END:
		i++
		if i >= maxLen {
			break
		}
	}
}

func (t *BinaryTree[T]) getPrintCount(str string) (count int) {
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count += 2 //汉字加两个字符
		} else {
			count++
		}
	}
	return
}

func (t *BinaryTree[T]) Add(ele T) {
	if t.Root == nil {
		t.Root = NewBinaryTreeNode[T](ele, nil)
		t.size++
		return
	}
	node := t.Root
	var parent *BinaryTreeNode[T]
	isLeft := true
	for node != nil {
		parent = node
		if node.element < ele {
			node = node.right
			isLeft = false
		} else if node.element > ele {
			node = node.left
			isLeft = true
		} else {
			return
		}
	}
	newNode := NewBinaryTreeNode(ele, parent)
	if !isLeft {
		parent.right = newNode
	} else {
		parent.left = newNode
	}
	t.size++
}

func (t *BinaryTree[T]) PreOrder(cb func(node *BinaryTreeNode[T])) {
	t.preOrder(t.Root, cb)
}

func (t *BinaryTree[T]) preOrder(node *BinaryTreeNode[T], cb func(node *BinaryTreeNode[T])) {
	if node == nil {
		return
	}
	cb(node)
	t.preOrder(node.left, cb)
	t.preOrder(node.right, cb)
}
func (t *BinaryTree[T]) InOrder(cb func(node *BinaryTreeNode[T])) {
	t.inOrder(t.Root, cb)

}
func (t *BinaryTree[T]) inOrder(node *BinaryTreeNode[T], cb func(node *BinaryTreeNode[T])) {
	if node == nil {
		return
	}
	t.inOrder(node.left, cb)
	cb(node)
	t.inOrder(node.right, cb)

}
func (t *BinaryTree[T]) PostOrder(cb func(node *BinaryTreeNode[T])) {
	t.postOrder(t.Root, cb)
}
func (t *BinaryTree[T]) postOrder(node *BinaryTreeNode[T], cb func(node *BinaryTreeNode[T])) {
	if node == nil {
		return
	}

	t.postOrder(node.left, cb)
	t.postOrder(node.right, cb)

}

func (t *BinaryTree[T]) LevelOrder(cb func(node *BinaryTreeNode[T])) {
	t.levelOrder(t.Root, cb)
}

func (t *BinaryTree[T]) levelOrder(node *BinaryTreeNode[T], cb func(node *BinaryTreeNode[T])) {
	if node == nil {
		return
	}
	que := NewSimpleQueue()
	que.Offer(node)
	for que.Len() != 0 {
		node := que.Poll().(*BinaryTreeNode[T])
		cb(node)
		if node.left != nil {
			que.Offer(node.left)
		}
		if node.right != nil {
			que.Offer(node.right)
		}
	}
}

func (t *BinaryTree[T]) PreOrderByVisitor(cb func(node *BinaryTreeNode[T])) {
	t.preOrder(t.Root, cb)
}
func (t *BinaryTree[T]) InOrderByVisitor(cb func(node *BinaryTreeNode[T])) {
	t.inOrder(t.Root, cb)
}
func (t *BinaryTree[T]) PostOrderByVisitor(cb func(node *BinaryTreeNode[T])) {
	t.postOrder(t.Root, cb)
}

func (t *BinaryTree[T]) LevelOrderByVisitor(cb func(node *BinaryTreeNode[T])) {
	t.levelOrder(t.Root, cb)
}

// 找后继者
func (t *BinaryTree[T]) FindSuccessor(node *BinaryTreeNode[T]) *BinaryTreeNode[T] {
	if node == nil {
		return nil
	}
	if node.right != nil {
		node = node.right
		for node.left != nil {
			node = node.left
		}
		return node
	}

	for node.parent != nil {
		node = node.parent
		if node == node.parent.left {
			return node
		}
	}
	return node
}

// 根据元素找到节点
func (t *BinaryTree[T]) FindNode(ele T) (*BinaryTreeNode[T], bool) {
	node := t.Root
	for node != nil {
		if ele > node.element {
			node = node.right
		} else if ele < node.element {
			node = node.left
		} else {
			return node, true
		}
	}
	return nil, false
}

// 找前驱
func (t *BinaryTree[T]) Predecessor(node *BinaryTreeNode[T]) *BinaryTreeNode[T] {
	if node == nil {
		return nil
	}
	if node.left != nil {
		node = node.left
		for node.right != nil {
			node = node.right
		}
		return node
	}

	for node.parent != nil {
		node = node.parent
		if node == node.parent.right {
			return node
		}
	}
	return node
}

func (t *BinaryTree[T]) Remove(ele T) {
	node, ok := t.FindNode(ele)
	if !ok {
		return
	}
	if node.IsLeaf() { //是叶子节点
		t.removeLeaf(node)
	} else if node.HasTwoChildren() { //度为2的节点
		replaceNode := t.FindSuccessor(node)
		node.element = replaceNode.element
		if replaceNode.IsLeaf() { //是叶子节点
			t.removeLeaf(replaceNode)
		} else {
			t.remove1(replaceNode)
		}
	} else { //度为1的节点
		t.remove1(node)
	}
	t.size--
}

func (t *BinaryTree[T]) removeLeaf(node *BinaryTreeNode[T]) {
	if node.parent == nil {
		t.Root = nil
	} else if node == node.parent.left {
		node.parent.left = nil
	} else {
		node.parent.right = nil
	}
}

// 移除度为1的节点
func (t *BinaryTree[T]) remove1(node *BinaryTreeNode[T]) {
	if node.parent == nil { //移除根节点
		if node.left != nil {
			t.Root = node.left
			node.left.parent = nil
		} else {
			t.Root = node.right
			node.right.parent = nil
		}
		return
	}
	if node == node.parent.left {
		if node.left != nil {
			node.parent.left = node.left
			node.left.parent = node.parent
		} else {
			node.parent.left = node.right
			node.right.parent = node.parent
		}
	} else {
		if node.right != nil {
			node.parent.right = node.right
			node.right.parent = node.parent
		} else {
			node.parent.right = node.left
			node.left.parent = node.parent
		}
	}
}
