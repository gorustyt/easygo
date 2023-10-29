package base

import (
	"cmp"
	"math"
)

const (
	BinarySearchTreeType = iota
	AvlTreeType
	RbTreeType
)

func NewRbTree[T CmpT]() *BinaryTree[T] {
	tree := NewBinaryTree[T]()
	tree.TreeType = RbTreeType
	return tree
}

type CmpT interface {
	any
	cmp.Ordered
}

// 二叉树
type BinaryTree[T CmpT] struct {
	size     int
	Root     *BinaryTreeNode[T]
	TreeType int
}

type BinaryTreeNode[T CmpT] struct {
	parent  *BinaryTreeNode[T] //父节点
	left    *BinaryTreeNode[T] //左子节点
	right   *BinaryTreeNode[T] //右子节点
	element T                  //元素
	Index   int                //在二叉树中索引,打印辅助
	height  int                //节点的高度
	isBlack bool               //是否式黑色
}

func (n *BinaryTreeNode[T]) tallerChild() (node *BinaryTreeNode[T]) {
	if n == nil {
		return nil
	}
	node = n.left
	if n.LeftHeight() < n.RightHeight() {
		node = n.right
	}
	return
}
func (n *BinaryTreeNode[T]) Height() int {
	return n.height
}
func (n *BinaryTreeNode[T]) ColorRed() {
	n.isBlack = false
}
func (n *BinaryTreeNode[T]) ColorBlack() {
	n.isBlack = true
}
func (n *BinaryTreeNode[T]) IsBlack() bool {
	return n == nil || n.isBlack
}
func (n *BinaryTreeNode[T]) IsRed() bool {
	return !n.IsBlack()
}
func (n *BinaryTreeNode[T]) IsLeftChild() bool {
	if n == nil {
		return false
	}
	return n.parent != nil && n.parent.left == n
}
func (n *BinaryTreeNode[T]) IsRightChild() bool {
	if n == nil {
		return false
	}
	return n.parent != nil && n.parent.right == n
}

// avl树是否平衡
func (n *BinaryTreeNode[T]) IsAvlBalance() bool {
	return math.Abs(float64(n.LeftHeight()-n.RightHeight())) <= 1
}

func (n *BinaryTreeNode[T]) LeftHeight() int {
	if n.left == nil {
		return 0
	}
	return n.left.height
}
func (n *BinaryTreeNode[T]) RightHeight() int {
	if n.right == nil {
		return 0
	}
	return n.right.height
}
func (n *BinaryTreeNode[T]) UpdateHeight() {
	if n == nil {
		return
	}
	n.height = int(math.Max(float64(n.LeftHeight()), float64(n.RightHeight()))) + 1
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

func (t *BinaryTree[T]) Add(ele T) {
	if t.Root == nil {
		t.Root = NewBinaryTreeNode[T](ele, nil)
		t.afterAdd(t.Root)
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
	t.afterAdd(newNode)
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
		isLeft := node.IsLeftChild()
		t.removeLeaf(node)
		t.afterRemove(node, isLeft)
	} else if node.HasTwoChildren() { //度为2的节点
		replaceNode := t.FindSuccessor(node)
		node.element = replaceNode.element
		isLeft := replaceNode.IsLeftChild()
		if replaceNode.IsLeaf() { //是叶子节点
			t.removeLeaf(replaceNode)
		} else {
			t.remove1(replaceNode)
		}
		t.afterRemove(replaceNode, isLeft)
	} else { //度为1的节点
		isLeft := node.IsLeftChild()
		node = t.remove1(node)
		t.afterRemove(node, isLeft)
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
func (t *BinaryTree[T]) remove1(node *BinaryTreeNode[T]) (replacement *BinaryTreeNode[T]) {
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
			replacement = node.left
			node.parent.left = node.left
			node.left.parent = node.parent
		} else {
			replacement = node.right
			node.parent.left = node.right
			node.right.parent = node.parent
		}
	} else {
		if node.right != nil {
			replacement = node.right
			node.parent.right = node.right
			node.right.parent = node.parent
		} else {
			replacement = node.left
			node.parent.right = node.left
			node.left.parent = node.parent
		}
	}
	return
}

func (t *BinaryTree[T]) updateHeight(node *BinaryTreeNode[T]) {
	leftHeight, rightHeight := 0, 0
	if node.left != nil {
		leftHeight = node.left.height
	}
	if node.right != nil {
		rightHeight = node.right.height
	}
	node.height = int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

func (t *BinaryTree[T]) isBalance(node *BinaryTreeNode[T]) bool {
	leftHeight, rightHeight := 0, 0
	if node.left != nil {
		leftHeight = node.left.height
	}
	if node.right != nil {
		rightHeight = node.right.height
	}
	return math.Abs(float64(leftHeight-rightHeight)) > 1
}

func (t *BinaryTree[T]) fixUpAvl(node *BinaryTreeNode[T]) {
	for node != nil {
		if node.IsAvlBalance() {
			node.UpdateHeight()
		} else {
			parent := node.tallerChild()
			child := parent.tallerChild()
			if parent.IsLeftChild() && child.IsLeftChild() {
				parent = node
				node = RotateRight(node)
			} else if parent.IsRightChild() && child.IsRightChild() {
				parent = node
				node = RotateLeft(node)
			} else if parent.IsLeftChild() && child.IsRightChild() {
				child = parent
				RotateLeft(parent)
				parent = node
				node = RotateRight(node)
			} else {
				child = parent
				RotateRight(parent)
				parent = node
				node = RotateLeft(node)
			}
			child.UpdateHeight()
			parent.UpdateHeight()
			node.UpdateHeight()
		}
		if node.parent == nil {
			t.Root = node
		}
		node = node.parent
	}
}

func (t *BinaryTree[T]) fixRbTreeAdd(node *BinaryTreeNode[T]) {
	parent := node.parent
	if node == t.Root || parent == nil {
		t.Root = node
		node.ColorBlack()
		return
	}
	p := t.Root
	if parent.IsBlack() {
		return
	}
	if parent.IsLeftChild() {
		if parent.parent.right.IsRed() { //第一种情况
			parent.parent.ColorRed()
			parent.ColorBlack()
			parent.parent.right.ColorBlack()
			t.afterAdd(parent.parent)
		} else if node.IsLeftChild() {
			parent.ColorBlack()
			parent.parent.ColorRed()
			p = RotateRight(parent.parent)
		} else if node.IsRightChild() {
			parent = RotateLeft(parent)
			parent.ColorBlack()
			parent.parent.ColorRed()
			p = RotateRight(parent.parent)
		}
	} else {
		if parent.parent.left.IsRed() { //第一种情况
			parent.parent.ColorRed()
			parent.ColorBlack()
			parent.parent.left.ColorBlack()
			t.afterAdd(parent.parent)
		} else if node.IsRightChild() {
			parent.ColorBlack()
			parent.parent.ColorRed()
			p = RotateLeft(parent.parent)
		} else if node.IsLeftChild() {
			parent = RotateRight(parent)
			parent.ColorBlack()
			parent.parent.ColorRed()
			p = RotateLeft(parent.parent)
		}
	}
	if p.parent == nil {
		p.ColorBlack()
		t.Root = p
	}
}

func (t *BinaryTree[T]) fixRbTreeRemove(x *BinaryTreeNode[T], isLeft bool) {
	var p = t.Root
	for x.parent != nil && x.IsBlack() {
		if isLeft {
			w := x.parent.right
			if w.IsRed() {
				x.parent.ColorRed()
				w.ColorBlack()
				p = RotateLeft(x.parent)
				w = x.parent.right
			}
			if w.left.IsBlack() && w.right.IsBlack() {
				w.ColorRed()
				x = x.parent
				isLeft = x.parent.IsLeftChild()
			} else {
				if w.left.IsRed() {
					w.left.ColorBlack()
					w.ColorRed()
					RotateRight(w)
					w = x.parent.right
				}
				if x.parent.IsBlack() {
					w.ColorBlack()
				} else {
					w.ColorRed()
				}
				w.right.ColorBlack()
				x.parent.ColorBlack()
				p = RotateLeft(x.parent)
			}
		} else {
			w := x.parent.left
			if w.IsRed() {
				x.parent.ColorRed()
				w.ColorBlack()
				p = RotateRight(x.parent)
				w = x.parent.left
			}
			if w.left.IsBlack() && w.right.IsBlack() {
				w.ColorRed()
				x = x.parent
				isLeft = x.parent.IsLeftChild()
			} else {
				if w.right.IsRed() {
					w.right.ColorBlack()
					w.ColorRed()
					RotateLeft(w)
					w = x.parent.left
				}
				if x.parent.IsBlack() {
					w.ColorBlack()
				} else {
					w.ColorRed()
				}
				w.left.ColorBlack()
				x.parent.ColorBlack()
				p = RotateRight(x.parent)
			}
		}
	}
	if p != nil && p.parent == nil {
		p.ColorBlack()
		t.Root = p
	}
	if x != nil && x.parent == nil {
		x.ColorBlack()
		t.Root = x
	}
	if x.IsRed() {
		x.ColorBlack()
		return
	}
}

func (t *BinaryTree[T]) afterAdd(node *BinaryTreeNode[T]) {
	switch t.TreeType {
	case BinarySearchTreeType:
	case AvlTreeType:
		t.fixUpAvl(node)
	case RbTreeType:
		t.fixRbTreeAdd(node)
	}

}

func (t *BinaryTree[T]) afterRemove(node *BinaryTreeNode[T], isLeft bool) {
	if node.parent == nil {
		return
	}
	switch t.TreeType {
	case AvlTreeType:
		t.fixUpAvl(node)
	case RbTreeType:
		t.fixRbTreeRemove(node, isLeft)
	}
}

func RotateLeft[T CmpT](node *BinaryTreeNode[T]) *BinaryTreeNode[T] {
	parent := node.right
	if node.parent != nil && node.IsLeftChild() {
		node.parent.left = parent
	} else if node.parent != nil && node.IsRightChild() {
		node.parent.right = parent
	}
	parent.parent = node.parent
	node.right = parent.left
	if node.right != nil {
		node.right.parent = node
	}
	parent.left = node
	node.parent = parent
	return parent
}

func RotateRight[T CmpT](node *BinaryTreeNode[T]) *BinaryTreeNode[T] {
	parent := node.left
	if node.parent != nil && node.IsLeftChild() {
		node.parent.left = parent
	} else if node.parent != nil && node.IsRightChild() {
		node.parent.right = parent
	}
	parent.parent = node.parent
	node.left = parent.right
	if node.left != nil {
		node.left.parent = node
	}
	parent.right = node
	node.parent = parent

	return parent
}
