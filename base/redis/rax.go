package redis

type RaxNode struct {
	IsCompress bool
	IsKey      bool
	Data       interface{}
	C          string //节点代表的字符
	Children   []*RaxNode
	Parent     *RaxNode //指向父节点
	childAtPos int      //指向父节点的位置
}

func NewRaxNode(c string, data interface{}) *RaxNode {
	return &RaxNode{
		C:          c,
		Data:       data,
		IsCompress: len(c) > 1,
		childAtPos: -1,
	}
}
func (n *RaxNode) LastChild() *RaxNode {
	if len(n.Children) == 0 {
		return nil
	}
	return n.Children[0]
}
func (n *RaxNode) FirstChild() *RaxNode {
	if len(n.Children) == 0 {
		return nil
	}
	return n.Children[len(n.Children)-1]
}

func (n *RaxNode) addChild(c string, remain bool, data interface{}, isLeaf bool) *RaxNode {
	pos := 0
	for index, v := range n.Children {
		if v.C > c {
			pos = index
			break
		}
	}
	child := NewRaxNode(c, data)
	child.Parent = n
	child.childAtPos = pos
	if isLeaf {
		child.IsKey = true
	}
	if len(n.Children) == 0 {
		n.Children = append(n.Children, child)
	} else {
		newChildren := make([]*RaxNode, len(n.Children)+1)
		copy(newChildren[:pos], n.Children[:pos])
		newChildren[pos] = child
		copy(newChildren[pos+1:], n.Children[pos:])
		for i := pos + 1; i < len(newChildren); i++ {
			newChildren[i].childAtPos++
		}
		n.Children = newChildren
	}
	if remain {
		//给child 添加空节点
		newChild := NewRaxNode("", nil)
		newChild.Parent = child
		newChild.childAtPos = 0
		child.Children = append(child.Children, newChild)
		return newChild
	}
	return child
}
func (n *RaxNode) CompressNode(c string) *RaxNode {
	n.IsCompress = true
	child := NewRaxNode("", nil)
	n.C = c
	child.Parent = n
	child.childAtPos = 0
	n.Children = append([]*RaxNode{}, child)
	return child
}

type Rax struct {
	Root        *RaxNode
	NumNode     int
	NumElements int
}

func NewRax() Rax {
	return Rax{
		Root: &RaxNode{},
	}
}
func (rax *Rax) Find(c string) (interface{}, bool) {
	node, i, j := rax.raxLowWalk(c)
	if i != len(c) || (!node.IsKey) || (node.IsCompress && j != 0) {
		return nil, false
	}
	return node.Data, true
}

// i字符串索引位置
// k子节点位置
func (rax *Rax) raxLowWalk(c string) (node *RaxNode, i, j int) {
	node = rax.Root
	for i < len(c) && (len(node.C) > 0 || len(node.Children) > 0) {
		if node.IsCompress {
			for j = 0; j < len(node.C) && i < len(c); {
				c1 := rax.getC(node.C, j)
				c2 := rax.getC(c, i)
				if c1 != c2 {
					break
				}
				j++
				i++
			}
			if j != len(node.C) {
				break
			}
		} else {
			for j = 0; j < len(node.Children); j++ {
				if rax.getC(node.Children[j].C, 0) == rax.getC(c, i) {
					break
				}
			}
			if j == len(node.Children) {
				break
			}
			i++
		}
		if node.IsCompress {
			j = 0
		}
		node = node.Children[j]
		if node.IsCompress {
			i--
		}
		j = 0
	}
	//if i == len(c) && node.Children != nil {
	//	if node.Children[0].IsKey {
	//		return node.Children[0], i, j
	//	}
	//}
	return node, i, j
}
func (rax *Rax) Remove() {

}

// 转换父子节点连接关系
func (rax *Rax) replaceNode(node1, node2 *RaxNode) {
	if node1.Parent == nil {
		root := NewRaxNode("", nil)
		root.Children = make([]*RaxNode, 1)
		node1.Parent = root
		node1.childAtPos = 0
		rax.Root = root
	}
	node2.Parent = node1.Parent
	node2.childAtPos = node1.childAtPos
	node1.Parent.Children[node1.childAtPos] = node2
}

func (rax *Rax) getC(c string, index int) string {
	return string(c[index])
}

func (rax *Rax) Add(c string, data interface{}) (old interface{}) {
	old = data
	node, i, j := rax.raxLowWalk(c)
	if i == len(c) && (!node.IsCompress || j == 0) {
		if !node.IsKey { //节点不是key
			node.addChild("", false, data, true)
		}
		if node.IsKey { //已经存在则覆盖
			old = node.Data
			node.Data = data
			return
		}
	}
	//* ============================= ALGO 2 ============================================
	//"ANNIBALE" =============>"ANNIBALE" -> "SCO" -> []
	//5) Inserting "ANNI"
	//     *
	//     *     "ANNI" -> "BALE" -> "SCO" -> []
	if node.IsCompress && i == len(c) { //是压缩节点，匹配到最后一个
		postfixNode := NewRaxNode("", data)
		trimNode := NewRaxNode("", nil)
		next := node.Children[0] //存储之前的节点

		next.Parent = postfixNode
		next.childAtPos = 0
		postfixNode.Children = append(postfixNode.Children, next)

		postfixLen := len(node.C) - j
		postfixNode.IsKey = true
		postfixNode.C = node.C[j:]
		postfixNode.IsCompress = postfixLen > 1

		postfixNode.childAtPos = 0
		postfixNode.Parent = trimNode
		trimNode.Children = append(trimNode.Children, postfixNode)

		trimNode.IsCompress = j > 1
		trimNode.C = node.C[:j]         //复制前半段字符串
		rax.replaceNode(node, trimNode) //用trimNode 覆盖原来节点

		return

		//* ============================= ALGO 1 ===========================================
		//"ANNIBALE" =============>"ANNIBALE" -> "SCO" -> []
		//           * 1) Inserting "ANNIENTARE"
		//     *
		//     *               |B| -> "ALE" -> "SCO" -> []
		//     *     "ANNI" -> |-|
		//     *               |E| -> (... continue algo ...) "NTARE" -> []
		//     *
		//     * 2) Inserting "ANNIBALI"
		//     *
		//     *                  |E| -> "SCO" -> []
		//     *     "ANNIBAL" -> |-|
		//     *                  |I| -> (... continue algo ...) []
		//     *
		//     * 3) Inserting "AGO" (Like case 1, but set iscompr = 0 into original node)
		//     *
		//     *            |N| -> "NIBALE" -> "SCO" -> []
		//     *     |A| -> |-|
		//     *            |G| -> (... continue algo ...) |O| -> []
		//     *
		//     * 4) Inserting "CIAO"
		//     *
		//     *     |A| -> "NNIBALE" -> "SCO" -> []
		//     *     |-|
		//     *     |C| -> (... continue algo ...) "IAO" -> []
	} else if node.IsCompress && i != len(c) { //是压缩节点匹配到压缩节点中间
		spiltNode := NewRaxNode(rax.getC(node.C, j), nil)
		postfix := NewRaxNode("", nil)
		next := node.Children[0] //存储之前的节点
		postfixLen := len(node.C) - i - 1
		if j == 0 {
			if node.IsKey {
				spiltNode.IsKey = node.IsKey
				spiltNode.Data = node.Data
			}
			rax.replaceNode(node, spiltNode)
		} else {
			trimNode := NewRaxNode(node.C, node.Data)
			trimNode.IsKey = node.IsKey
			trimNode.C = node.C[:j]
			rax.replaceNode(node, trimNode)
			//这里用空节点替代，保证压缩节点只有一个子节点
			emptyNode := NewRaxNode("", nil)
			emptyNode.Parent = trimNode
			emptyNode.childAtPos = 0
			trimNode.Children = append(trimNode.Children, emptyNode)
			trimNode = emptyNode

			spiltNode.Parent = trimNode
			spiltNode.childAtPos = 0
			trimNode.Children = append(trimNode.Children, spiltNode)

		}
		if postfixLen > 0 { //
			postfix.IsCompress = postfixLen > 1

			next.Parent = postfix
			next.childAtPos = 0
			postfix.Children = append(postfix.Children, next)
			postfix.C = node.C[j+1:]
		} else { //case 2
			postfix = next
		}

		postfix.Parent = spiltNode
		postfix.childAtPos = 0
		spiltNode.Children = append(spiltNode.Children, postfix)

		node = spiltNode.Parent
	}

	for i < len(c) {
		child := NewRaxNode("", nil)
		if (len(node.C) == 0 && len(node.Children) == 0) && len(c)-i > 1 {
			compressSize := len(c) - i
			node.Children = append(node.Children, child)
			child = node.CompressNode(c[i:])
			i += compressSize
		} else { //不是压缩节点
			child = node.addChild(rax.getC(c, i), len(c)-i-1 != 1, nil, false)
			i++
		}
		node = child
	}
	node.IsKey = true
	node.Data = data
	return old
}

func (rax *Rax) raxRemoveChild(parent *RaxNode, child *RaxNode) {
	if parent.IsCompress {
		parent.IsCompress = false
		parent.Data = nil
		parent.IsKey = false
		return
	}
	for index, v := range parent.Children {
		if v == child {
			parent.Children = append(parent.Children[:index], parent.Children[:index]...)
			break
		}
	}
}
