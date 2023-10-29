package base

import (
	"fmt"
	"math"
	"strings"
	"time"
	"unicode"
)

// 获取二叉树的每一层节点索引
func getPrintList[T CmpT](node *BinaryTreeNode[T], cb func(node *BinaryTreeNode[T]) string) (res [][]string, maxLen int) {
	if node == nil {
		return
	}
	node.Index = 0 //防止旋转操作导致的索引变化
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
		length := getPrintTreeCount(arr[index])
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
func PrintBinaryTree[T CmpT](node *BinaryTreeNode[T], cb func(node *BinaryTreeNode[T]) string) {
	if node == nil {
		return
	}
	list, maxLen := getPrintList(node, cb)
	if maxLen%2 == 0 {
		maxLen++
	}
	//填充长度
	for i, arr := range list {
		for j, v := range arr {
			length := getPrintTreeCount(v)
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
	printTree(list, maxLen)
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

func printTree(list [][]string, maxLen int, compress ...bool) {
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
		printTreeCompress(values, lines, 2)
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
	fmt.Println("===================================================================")
	time.Sleep(10 * time.Millisecond)
	for index, v := range values {
		fmt.Println(v)
		time.Sleep(10 * time.Millisecond)
		if index+1 < len(values) {
			time.Sleep(10 * time.Millisecond)
			fmt.Println(lines[index+1])
			time.Sleep(10 * time.Millisecond)
		}
	}

}

func printTreeCompress(values, lines []string, compressCount int) {
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

func getPrintTreeCount(str string) (count int) {
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count += 2 //汉字加两个字符
		} else {
			count++
		}
	}
	return
}
