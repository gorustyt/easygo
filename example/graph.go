package example

import (
	"fmt"
	"github.com/lirongyangtao/mygo/base"
)

type Visitor struct {
}

func (vi *Visitor) Visitor(v *base.Vertex) {
	fmt.Println(v.Value)
}

// 深度有效搜索
func Dfs() {
	graph := base.NewGraph()
	graph.AddEdgeWithWeight(3, 7, 9)
	graph.AddEdgeWithWeight(7, 3, 9)
	graph.AddEdgeWithWeight(1, 5, 9)
	graph.AddEdgeWithWeight(5, 1, 9)
	graph.AddEdgeWithWeight(1, 0, 9)
	graph.AddEdgeWithWeight(0, 1, 9)
	graph.AddEdgeWithWeight(1, 2, 9)
	graph.AddEdgeWithWeight(2, 1, 9)
	graph.AddEdgeWithWeight(2, 4, 9)
	graph.AddEdgeWithWeight(4, 2, 9)
	graph.AddEdgeWithWeight(1, 6, 9)
	graph.AddEdgeWithWeight(6, 1, 9)
	graph.Dfs(0, &Visitor{})
}

// 深度有效搜索
func Dfs1() {
	graph := base.NewGraph()
	graph.AddEdgeWithWeight("a", "b", 9)
	graph.AddEdgeWithWeight("d", "a", 9)
	graph.AddEdgeWithWeight("a", "e", 9)
	graph.AddEdgeWithWeight("b", "e", 9)
	graph.AddEdgeWithWeight("e", "c", 9)
	graph.AddEdgeWithWeight("c", "b", 9)
	graph.AddEdgeWithWeight("e", "f", 9)
	graph.AddEdgeWithWeight("f", "c", 9)
	graph.Dfs("a", &Visitor{}) //a,b,e,f,c
}

// 拓扑排序
func TopLogicSort() {
	graph := base.NewGraph()
	graph.AddEdgeWithWeight(0, 2, 9)
	graph.AddEdgeWithWeight(1, 0, 3)

	graph.AddEdgeWithWeight(2, 5, 9)
	graph.AddEdgeWithWeight(2, 6, 3)

	graph.AddEdgeWithWeight(3, 1, 9)
	graph.AddEdgeWithWeight(3, 5, 3)
	graph.AddEdgeWithWeight(3, 7, 3)

	graph.AddEdgeWithWeight(5, 7, 3)
	graph.AddEdgeWithWeight(6, 4, 3)
	graph.AddEdgeWithWeight(7, 6, 3)
	fmt.Println(graph.TopLogicSort()) //result: [3 1 0 2 5 7 6 4]
}

// 宽度有效搜索
func Bfs() {
	graph := base.NewGraph()
	graph.AddEdgeWithWeight("A", "B", 9)
	graph.AddEdgeWithWeight("A", "F", 3)
	graph.AddEdgeWithWeight("B", "C", 2)
	graph.AddEdgeWithWeight("B", "I", 5)
	graph.AddEdgeWithWeight("B", "G", 1)
	graph.AddEdgeWithWeight("C", "I", 6)
	graph.AddEdgeWithWeight("C", "D", 6)
	graph.AddEdgeWithWeight("D", "I", 6)
	graph.AddEdgeWithWeight("D", "G", 6)
	graph.AddEdgeWithWeight("D", "E", 6)
	graph.AddEdgeWithWeight("D", "H", 6)
	graph.AddEdgeWithWeight("E", "H", 6)
	graph.AddEdgeWithWeight("E", "F", 6)
	graph.AddEdgeWithWeight("F", "G", 6)
	graph.AddEdgeWithWeight("G", "H", 6)

	graph.AddEdgeWithWeight("B", "A", 9)
	graph.AddEdgeWithWeight("F", "A", 3)
	graph.AddEdgeWithWeight("C", "B", 2)
	graph.AddEdgeWithWeight("I", "B", 5)
	graph.AddEdgeWithWeight("G", "B", 1)
	graph.AddEdgeWithWeight("I", "C", 6)
	graph.AddEdgeWithWeight("D", "C", 6)
	graph.AddEdgeWithWeight("I", "D", 6)
	graph.AddEdgeWithWeight("G", "D", 6)
	graph.AddEdgeWithWeight("E", "D", 6)
	graph.AddEdgeWithWeight("H", "D", 6)
	graph.AddEdgeWithWeight("H", "E", 6)
	graph.AddEdgeWithWeight("F", "E", 6)
	graph.AddEdgeWithWeight("G", "F", 6)
	graph.AddEdgeWithWeight("H", "G", 6)
	graph.Bfs("A", &Visitor{})
	graph.Print()
}

func Kruskal() {
	graph := UnDirectGraph(MST_01)
	edges := graph.Kruskal()
	for _, v := range edges {
		fmt.Println(&v)
	}
}

func Dijkstra() {
	graph := UnDirectGraph(SP)
	edges := graph.Dijkstra("A")
	for k, v := range edges {
		fmt.Printf("A==========>%v,path=%v,weight:%v\n", k, v.MinPath, v.Weight)
	}
}

func BellFord() {
	graph := DirectGraph(NEGATIVE_WEIGHT1)
	edges := graph.BellFord("A")
	for k, v := range edges {
		fmt.Printf("A==========>%v,path=%v,weight:%v\n", k, v.MinPath, v.Weight)
	}
}

func Floyd() {
	graph := DirectGraph(NEGATIVE_WEIGHT1)
	edges := graph.Floyd()
	for from, toPathInfos := range edges {
		for to, path := range toPathInfos {
			fmt.Printf("%v==========>%v,path=%v,weight:%v\n", from, to, path, path.Weight)
		}
	}
}

func Floyd1() {
	graph := DirectGraph(SP)
	edges := graph.Floyd()
	for from, toPathInfos := range edges {
		for to, path := range toPathInfos {
			fmt.Printf("%v==========>%v,path=%v,weight:%v\n", from, to, path, path.Weight)
		}
	}
}
