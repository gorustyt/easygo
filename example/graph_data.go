package example

import "github.com/lirongyangtao/mygo/base"

var (
	BFS_01 = [][]string{
		{"A", "B"}, {"A", "F"},
		{"B", "C"}, {"B", "I"}, {"B", "G"},
		{"C", "I"}, {"C", "D"},
		{"D", "I"}, {"D", "G"}, {"D", "E"}, {"D", "H"},
		{"E", "H"}, {"E", "F"},
		{"F", "G"},
		{"G", "H"},
	}
	BFS_02 = [][]int{
		{0, 1}, {0, 4},
		{1, 2},
		{2, 0}, {2, 4}, {2, 5},
		{3, 1},
		{4, 6}, {4, 7},
		{5, 3}, {5, 7},
		{6, 2}, {6, 7},
	}
	BFS_03 = [][]int{
		{0, 2}, {0, 3},
		{1, 2}, {1, 3}, {1, 6},
		{2, 4},
		{3, 7},
		{4, 6},
		{5, 6},
		{6, 7},
	}
	BFS_04 = [][]int{
		{1, 2}, {1, 3}, {1, 5},
		{2, 0},
		{3, 5},
		{5, 6}, {5, 7},
		{6, 2},
		{7, 6},
	}
	DFS_01 = [][]int{
		{0, 1},
		{1, 3}, {1, 5}, {1, 6}, {1, 2},
		{2, 4},
		{3, 7},
	}
	DFS_02 = [][]string{
		{"a", "b"}, {"a", "e"},
		{"b", "e"},
		{"c", "b"},
		{"d", "a"},
		{"e", "c"}, {"e", "f"},
		{"f", "c"},
	}
	TOPO = [][]int{
		{0, 2},
		{1, 0},
		{2, 5}, {2, 6},
		{3, 1}, {3, 5}, {3, 7},
		{5, 7},
		{6, 4},
		{7, 6},
	}
	NO_WEIGHT2 = [][]int{
		{0, 3},
		{1, 3}, {1, 6},
		{2, 1},
		{3, 5},
		{6, 2}, {6, 5},
		{4, 7},
	}

	NO_WEIGHT3 = [][]int{
		{0, 1}, {0, 2},
		{1, 2}, {1, 5},
		{2, 4}, {2, 5},
		{5, 6}, {7, 6},
		{3},
	}
	MST_01 = [][]int{
		{0, 2, 2}, {0, 4, 7},
		{1, 2, 3}, {1, 5, 1}, {1, 6, 7},
		{2, 4, 4}, {2, 5, 3}, {2, 6, 6},
		{3, 7, 9},
		{4, 6, 8},
		{5, 6, 4}, {5, 7, 5},
	}
	MST_02 = [][]base.EdgeInfo{
		{{"A", "B", 17}, {"A", "F", 1}, {"A", "E", 16},
			{"B", "C", 6}, {"B", "D", 5}, {"B", "F", 11},
			{"C", "D", 10},
			{"D", "E", 4}, {"D", "F", 14},
			{"E", "F", 33}},
	}
	WEIGHT3 = [][]base.EdgeInfo{
		{{"广州", "佛山", 100}, {"广州", "珠海", 200}, {"广州", "肇庆", 200},
			{"佛山", "珠海", 50}, {"佛山", "深圳", 150},
			{"肇庆", "珠海", 100}, {"肇庆", "南宁", 150},
			{"珠海", "南宁", 350}, {"珠海", "深圳", 100},
			{"南宁", "香港", 500}, {"南宁", "深圳", 400},
			{"深圳", "香港", 150}},
	}
	SP = [][]base.EdgeInfo{{
		{"A", "B", 10}, {"A", "D", 30}, {"A", "E", 100},
		{"B", "C", 50},
		{"C", "E", 10},
		{"D", "C", 20}, {"D", "E", 60}},
	}
	BF_SP = [][]base.EdgeInfo{{
		{"A", "B", 10}, {"A", "E", 8},
		{"B", "C", 8}, {"B", "E", -5},
		{"D", "C", 2}, {"D", "F", 6},
		{"E", "D", 7}, {"E", "F", 3},
	}}
	WEIGHT5 = [][]int{
		{0, 14, 1}, {0, 4, 8},
		{1, 2, 9},
		{2, 3, 6}, {2, 5, 9},
		{3, 17, 1}, {3, 10, 4},
		{4, 5, 2}, {4, 8, 2},
		{5, 6, 6}, {5, 8, 1}, {5, 9, 4},
		{6, 9, 8},
		{7, 11, 4},
		{8, 9, 1}, {8, 11, 2}, {8, 12, 7},
		{9, 10, 7}, {9, 13, 4},
		{10, 13, 2},
		{11, 12, 7}, {11, 15, 4},
		{12, 13, 2}, {12, 16, 2},
		{13, 16, 7},
		{15, 16, 7}, {15, 17, 7},
		{16, 17, 2},
	}
	NEGATIVE_WEIGHT1 = [][]base.EdgeInfo{{
		{"A", "B", -1}, {"A", "C", 4},
		{"B", "C", 3}, {"B", "D", 2}, {"B", "E", 2},
		{"D", "B", 1}, {"D", "C", 5},
		{"E", "D", -3},
	}}
	/**
	 * 有负权环
	 */
	NEGATIVE_WEIGHT2 = [][]int{
		{0, 1, 1},
		{1, 2, 7},
		{1, 0, -2},
	}
)

func UnDirectGraph(data interface{}) base.Graph {
	return initGraph(data, false)
}

func initGraph(data interface{}, direct bool) base.Graph {
	g := base.NewGraph()
	switch t := data.(type) {
	case [][]string:
		for _, v := range t {
			g.AddEdge(v[0], v[1])
			if !direct {
				g.AddEdge(v[1], v[0])
			}
		}
	case [][]int:
		for _, v := range t {
			if len(v) == 3 {
				g.AddEdgeWithWeight(v[0], v[1], v[2])
				if !direct {
					g.AddEdgeWithWeight(v[1], v[0], v[2])
				}
			} else if len(v) == 2 {
				g.AddEdge(v[0], v[1])
				if !direct {
					g.AddEdge(v[1], v[0])
				}
			}
		}
	case [][]base.EdgeInfo:
		for _, v := range t {
			for _, v1 := range v {
				g.AddEdgeWithWeight(v1.From, v1.To, v1.Weight)
				if !direct {
					g.AddEdgeWithWeight(v1.To, v1.From, v1.Weight)
				}
			}
		}
	}
	return g
}
func DirectGraph(data interface{}) base.Graph {
	return initGraph(data, true)
}
