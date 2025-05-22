package model

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

var (
	ErrDataConsistency = errors.New("Data consistency violated (such data exists or cannot be made)")
	EmptyValue         = errors.New("Object is not found or value is empty")
)

// type Node struct {
// 	Id int `json:"id"`
// }

// type Link struct {
// 	Source Node `json:"source"`
// 	Target Node `json:"target"`
// }

type Component struct {
	Component int `json:"component"`
}

type IsEuler struct {
	IsEuler bool `json:"isEuler"`
}

type IsBipartition struct {
	IsBipartition bool `json:"isBipartition"`
}

type Node struct {
	Id     int
	Label  string
	Color  string
	Weight int
	X      float64
	Y      float64
}

type Edge struct {
	Id     string
	Source Node
	Target Node
	Label  string
	Color  string
	Weight int
}

// по умолчанию граф неориентированный
type Graph struct {
	Nodes    []Node
	Edges    []Edge
	oriented bool
}

func (g *Graph) MakeOriented() {
	g.oriented = true
}

func (g *Graph) IsOriented() bool {
	return g.oriented
}

func (g *Graph) AddNode(new_node Node) error {
	for _, node := range g.Nodes {
		if new_node.Id == node.Id || new_node.Label == node.Label {
			return ErrDataConsistency
		}
	}
	g.Nodes = append(g.Nodes, new_node)
	return nil
}

func (g *Graph) AddNodeByInfo(id int, label string, color string, weight int, x float64, y float64) error {
	new_node := Node{Id: id, Label: label, Color: color, Weight: weight, X: x, Y: y}
	return g.AddNode(new_node)
}

func (g *Graph) IsNodeById(id int) bool {
	for _, node := range g.Nodes {
		if node.Id == id {
			return true
		}
	}
	return false
}

func (g *Graph) IsNodeByLabel(label string) bool {
	for _, node := range g.Nodes {
		if node.Label == label {
			return true
		}
	}
	return false
}

func (g *Graph) IsEdgeById(id string) bool {
	for _, edge := range g.Edges {
		if edge.Id == id {
			return true
		}
	}
	return false
}

func (g *Graph) IsEdgeByLable(label string) bool {
	for _, edge := range g.Edges {
		if edge.Label == label {
			return true
		}
	}
	return false
}

func (g *Graph) FindEdge(source, target string) (bool, Edge) {
	for _, edge := range g.Edges {
		if g.oriented {
			if source == edge.Source.Label && target == edge.Target.Label {
				return true, edge
			}
		} else {
			if source == edge.Source.Label && target == edge.Target.Label || source == edge.Target.Label && target == edge.Source.Label {
				return true, edge
			}
		}
	}
	return false, Edge{}
}

func (g *Graph) AddEdge(new_edge Edge) error {
	// if g.IsEdgeById(new_edge.Id) {
	// 	return ErrDataConsistency
	// }
	if !g.IsNodeByLabel(new_edge.Source.Label) {
		g.AddNode(new_edge.Source)
	}
	if !g.IsNodeByLabel(new_edge.Target.Label) {
		g.AddNode(new_edge.Target)
	}
	g.Edges = append(g.Edges, new_edge)
	return nil
}

func (g *Graph) AddEdgeByInfo(source Node, target Node, id string, label string, color string, weight int) error {
	new_edge := Edge{Source: source, Target: target, Id: id, Label: label, Color: color, Weight: weight}
	return g.AddEdge(new_edge)
}

func (g *Graph) FindNodeByLabel(label string) (Node, error) {
	for _, node := range g.Nodes {
		if node.Label == label {
			return node, nil
		}
	}
	return Node{}, EmptyValue
}

func (g *Graph) FindNodeById(id int) (Node, error) {
	for _, node := range g.Nodes {
		if node.Id == id {
			return node, nil
		}
	}
	return Node{}, EmptyValue
}

// Матрица смежности вершин, но вместо индексов - Label
func (g *Graph) NodeLabelAdjacentMatrix() map[string]map[string]int {
	matrix := make(map[string]map[string]int)
	for _, node1 := range g.Nodes {
		matrix[node1.Label] = make(map[string]int)
		for _, node2 := range g.Nodes {
			matrix[node1.Label][node2.Label] = 0
		}
	}
	for _, edge := range g.Edges {
		matrix[edge.Source.Label][edge.Target.Label] = 1
		matrix[edge.Target.Label][edge.Source.Label] = 1
	}
	return matrix
}

// Матрица смежности вершин, но вместо индексов - Id
func (g *Graph) NodeIdAdjacentMatrix() map[int]map[int]int {
	matrix := make(map[int]map[int]int)
	for _, node1 := range g.Nodes {
		matrix[node1.Id] = make(map[int]int)
		for _, node2 := range g.Nodes {
			matrix[node1.Id][node2.Id] = 0
		}
	}
	for _, edge := range g.Edges {
		matrix[edge.Source.Id][edge.Target.Id] = 1
		matrix[edge.Target.Id][edge.Source.Id] = 1
	}
	return matrix
}

// Матрица смежности ребер, но вместо индексов - Label
func (g *Graph) EdgeLabelAdjacentMatrix() map[string]map[string]int {
	matrix := make(map[string]map[string]int)
	for _, edge1 := range g.Edges {
		matrix[edge1.Label] = make(map[string]int)
		for _, edge2 := range g.Edges {
			matrix[edge1.Label][edge2.Label] = 0
		}
	}
	for _, edge1 := range g.Edges {
		for _, edge2 := range g.Edges {
			if (edge1.Source.Label == edge2.Source.Label || edge1.Source.Label == edge2.Target.Label || edge1.Target.Label == edge2.Source.Label || edge1.Target.Label == edge2.Target.Label) && (edge1.Label != edge2.Label) {
				matrix[edge1.Label][edge2.Label] = 1
				matrix[edge2.Label][edge1.Label] = 1
			}
		}
	}
	return matrix
}

// Матрица смежности ребер, но вместо индексов - Id
func (g *Graph) EdgeIdAdjacentMatrix() map[string]map[string]int {
	matrix := make(map[string]map[string]int)
	for _, edge1 := range g.Edges {
		matrix[edge1.Id] = make(map[string]int)
		for _, edge2 := range g.Edges {
			matrix[edge1.Id][edge2.Id] = 0
		}
	}
	for _, edge1 := range g.Edges {
		for _, edge2 := range g.Edges {
			if (edge1.Source.Id == edge2.Source.Id || edge1.Source.Id == edge2.Target.Id || edge1.Target.Id == edge2.Source.Id || edge1.Target.Id == edge2.Target.Id) && (edge1.Id != edge2.Id) {
				matrix[edge1.Id][edge2.Id] = 1
				matrix[edge2.Id][edge1.Id] = 1
			}
		}
	}
	return matrix
}

// Реализация алгоритма Дейкстры, возвращает (минимальный путь, матрицу расстояний)
// Все вершины должны быть достижимы, расстояния не меньше 0
func (g *Graph) MinPath(source Node, target Node, edges_have_weights bool) (int, map[string]int) {
	visited := make(map[string]bool)
	weights_path := make(map[string]int)
	for _, node := range g.Nodes {
		visited[node.Label] = false
		weights_path[node.Label] = math.MaxInt
	}
	weights_path[source.Label] = 0
	weights_matrix := make(map[string]map[string]int)
	for _, node1 := range g.Nodes {
		weights_matrix[node1.Label] = make(map[string]int)
		for _, node2 := range g.Nodes {
			weights_matrix[node1.Label][node2.Label] = 0
		}
	}
	for _, edge := range g.Edges {
		if edges_have_weights {
			weights_matrix[edge.Source.Label][edge.Target.Label] = edge.Weight
			weights_matrix[edge.Target.Label][edge.Source.Label] = edge.Weight
		} else {
			weights_matrix[edge.Source.Label][edge.Target.Label] = 1
			weights_matrix[edge.Target.Label][edge.Source.Label] = 1
		}
	}
	min_node := "1"
	for min_node != "" {
		min_val := math.MaxInt
		min_node = ""
		for _, node := range g.Nodes {
			if !visited[node.Label] && weights_path[node.Label] < min_val {
				min_val = weights_path[node.Label]
				min_node = node.Label
			}
		}
		if min_node != "" {
			for _, node := range g.Nodes {
				if weights_matrix[min_node][node.Label] > 0 {
					temp := min_val + weights_matrix[min_node][node.Label]
					if temp < weights_path[node.Label] {
						weights_path[node.Label] = temp
					}
				}
			}
			visited[min_node] = true
		}
	}
	return weights_path[target.Label], weights_path
}

// Возвращает матрицу кратчайших расстояний между каждой парой вершин, на диагонали 0
func (g *Graph) DistanceMatrix(edges_have_weights bool) map[string]map[string]int {
	matrix := make(map[string]map[string]int)
	for _, node := range g.Nodes {
		_, matrix[node.Label] = g.MinPath(node, node, edges_have_weights)
		matrix[node.Label][node.Label] = 0
	}
	return matrix
}

func (g *Graph) IsEdgesAdjacent(edge1 Edge, edge2 Edge) bool {
	if edge1.Source.Id == edge2.Source.Id || edge1.Source.Id == edge2.Target.Id || edge1.Target.Id == edge2.Source.Id || edge1.Target.Id == edge2.Target.Id {
		return true
	}
	return false
}

func Max_(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func (g *Graph) Intersect(graph *Graph) *Graph {
	nodes_g_set := make(map[string]Node)
	for _, node_g := range g.Nodes {
		nodes_g_set[node_g.Label] = node_g
	}
	nodes_answer := make(map[string]Node)
	for _, node_graph := range graph.Nodes {
		_, exist := nodes_g_set[node_graph.Label]
		if exist {
			nodes_answer[node_graph.Label] = nodes_g_set[node_graph.Label]
		}
	}
	edges_g := make(map[string]Edge)
	edges_final := make(map[string]Edge)
	for _, edge := range g.Edges {
		edges_g[edge.Id] = edge
		edges_final[edge.Id] = edge
	}
	flag := false
	for id, edge_g := range edges_g {
		for _, edge_graph := range graph.Edges {
			if edge_graph.Source.Label == edge_g.Source.Label && edge_graph.Target.Label == edge_g.Target.Label || edge_graph.Source.Label == edge_g.Target.Label && edge_graph.Target.Label == edge_g.Source.Label {
				flag = true
			}
		}
		if !flag {
			delete(edges_final, id)
		}
	}
	answer := new(Graph)
	node_id := 0
	for _, node := range nodes_g_set {
		node.Id = node_id
		node_id++
		answer.AddNode(node)
	}
	for _, edge := range edges_final {
		answer.AddEdge(edge) // ребра соединяются по label, не по id
	}
	return answer
}

func (g *Graph) Union(graph *Graph) *Graph {
	nodes_set := make(map[string]Node)
	for _, node := range g.Nodes {
		nodes_set[node.Label] = node
	}
	for _, node := range graph.Nodes {
		nodes_set[node.Label] = node
	}
	edges_set := make(map[string]Edge)
	for _, edge := range g.Edges {
		edges_set[fmt.Sprintf("%s,%s", edge.Source.Label, edge.Target.Label)] = edge
	}
	for _, edge := range graph.Edges {
		key1 := fmt.Sprintf("%s,%s", edge.Source.Label, edge.Target.Label)
		key2 := fmt.Sprintf("%s,%s", edge.Target.Label, edge.Source.Label)
		_, key1_flg := edges_set[key1]
		_, key2_flg := edges_set[key2]
		if !key1_flg && !key2_flg {
			edges_set[key1] = edge
		}
	}
	answer := new(Graph)
	node_id := 0
	for _, node := range nodes_set {
		answer.AddNodeByInfo(node_id, node.Label, node.Color, node.Weight, node.X, node.Y)
		node_id++
	}
	edge_id := 0
	for _, edge := range edges_set {
		edge.Id = strconv.Itoa(edge_id)
		edge_id++
		answer.AddEdge(edge)
	}
	return answer
}

func (g *Graph) Join(graph *Graph) *Graph {
	unioned_graphs := g.Union(graph)
	nodes_g1 := make(map[string]struct{})
	nodes_g2 := make(map[string]struct{})
	for _, node := range g.Nodes {
		nodes_g1[node.Label] = struct{}{}
	}
	for _, node := range graph.Nodes {
		nodes_g2[node.Label] = struct{}{}
	}
	nodes_g1_wo_g2 := make(map[string]struct{})
	nodes_g2_wo_g1 := make(map[string]struct{})
	for node := range nodes_g1 {
		if _, exist := nodes_g2[node]; !exist {
			nodes_g1_wo_g2[node] = struct{}{}
		}
	}
	for node := range nodes_g2 {
		if _, exist := nodes_g1[node]; !exist {
			nodes_g2_wo_g1[node] = struct{}{}
		}
	}
	edge_id := len(unioned_graphs.Edges)
	for node1 := range nodes_g1_wo_g2 {
		for node2 := range nodes_g2_wo_g1 {
			src, _ := unioned_graphs.FindNodeByLabel(node1)
			trg, _ := unioned_graphs.FindNodeByLabel(node2)
			unioned_graphs.AddEdgeByInfo(src, trg, strconv.Itoa(edge_id), "", "", 0)
			edge_id++
		}
	}
	return unioned_graphs
}

func (g *Graph) CartesianProduct(graph *Graph) *Graph {
	answer := new(Graph)
	id := 1
	for _, node1 := range g.Nodes {
		for _, node2 := range graph.Nodes {
			new_node_label := fmt.Sprintf("(%s,%s)", node1.Label, node2.Label)
			answer.AddNode(Node{Id: id, Label: new_node_label})
			id++
		}
	}
	edge_id := 1
	for _, edge := range g.Edges {
		for _, node2 := range graph.Nodes {
			src, _ := answer.FindNodeByLabel(fmt.Sprintf("(%s,%s)", edge.Source.Label, node2.Label))
			trg, _ := answer.FindNodeByLabel(fmt.Sprintf("(%s,%s)", edge.Target.Label, node2.Label))
			answer.AddEdgeByInfo(src, trg, strconv.Itoa(edge_id), "", "", 0)
			edge_id++
		}
	}
	for _, edge := range graph.Edges {
		for _, node1 := range g.Nodes {
			src, _ := answer.FindNodeByLabel(fmt.Sprintf("(%s,%s)", node1.Label, edge.Source.Label))
			trg, _ := answer.FindNodeByLabel(fmt.Sprintf("(%s,%s)", node1.Label, edge.Target.Label))
			answer.AddEdgeByInfo(src, trg, strconv.Itoa(edge_id), "", "", 0)
			edge_id++
		}
	}
	return answer
}

func (g *Graph) TensorProduct(graph *Graph) *Graph {
	answer := new(Graph)
	id := 1
	for _, node1 := range g.Nodes {
		for _, node2 := range graph.Nodes {
			new_node_label := fmt.Sprintf("(%s,%s)", node1.Label, node2.Label)
			answer.AddNode(Node{Id: id, Label: new_node_label})
			id++
		}
	}
	edge_id := 1
	for _, edge1 := range g.Edges {
		for _, edge2 := range graph.Edges {
			node1, _ := answer.FindNodeByLabel(fmt.Sprintf("(%s,%s)", edge1.Source.Label, edge2.Source.Label))
			node2, _ := answer.FindNodeByLabel(fmt.Sprintf("(%s,%s)", edge1.Target.Label, edge2.Target.Label))
			answer.AddEdgeByInfo(node1, node2, strconv.Itoa(edge_id), "", "", 0)
			edge_id++
			node3, _ := answer.FindNodeByLabel(fmt.Sprintf("(%s,%s)", edge1.Source.Label, edge2.Target.Label))
			node4, _ := answer.FindNodeByLabel(fmt.Sprintf("(%s,%s)", edge1.Target.Label, edge2.Source.Label))
			answer.AddEdgeByInfo(node3, node4, strconv.Itoa(edge_id), "", "", 0)
			edge_id++
		}
	}
	return answer
}

func (g *Graph) LexicographicalProduct(graph *Graph) *Graph {
	answer := new(Graph)
	id := 1
	for _, node1 := range g.Nodes {
		for _, node2 := range graph.Nodes {
			new_node_label := fmt.Sprintf("(%s,%s)", node1.Label, node2.Label)
			answer.AddNode(Node{Id: id, Label: new_node_label})
			id++
		}
	}
	edge_id := 1
	for _, node1 := range g.Nodes {
		for _, edge2 := range graph.Edges {
			src1, _ := answer.FindNodeByLabel(fmt.Sprintf("(%s,%s)", node1.Label, edge2.Source.Label))
			trg1, _ := answer.FindNodeByLabel(fmt.Sprintf("(%s,%s)", node1.Label, edge2.Target.Label))
			answer.AddEdgeByInfo(src1, trg1, strconv.Itoa(edge_id), "", "", 0)
			edge_id++
			src2, _ := answer.FindNodeByLabel(fmt.Sprintf("(%s,%s)", node1.Label, edge2.Source.Label))
			trg2, _ := answer.FindNodeByLabel(fmt.Sprintf("(%s,%s)", node1.Label, edge2.Target.Label))
			answer.AddEdgeByInfo(src2, trg2, strconv.Itoa(edge_id), "", "", 0)
			edge_id++
		}
	}
	for _, edge1 := range g.Edges {
		for _, node1 := range graph.Nodes {
			for _, node2 := range graph.Nodes {
				src, _ := answer.FindNodeByLabel(fmt.Sprintf("(%s,%s)", edge1.Source.Label, node1.Label))
				trg, _ := answer.FindNodeByLabel(fmt.Sprintf("(%s,%s)", edge1.Target.Label, node2.Label))
				answer.AddEdgeByInfo(src, trg, strconv.Itoa(edge_id), "", "", 0)
				edge_id++
			}
		}
	}

	return answer
}

func (g *Graph) hamiltonianCycleAuxilaryCheck(index int, node string, path []string, matrix map[string]map[string]int) bool {
	if matrix[path[index-1]][node] == 0 {
		return false
	}
	for _, node1 := range path {
		if node == node1 {
			return false
		}
	}
	return true
}

func (g *Graph) hamiltonianCycleAuxilary(index int, path []string, matrix map[string]map[string]int) (bool, []string) {
	if index == len(g.Nodes) {
		if matrix[path[index-1]][path[0]] == 1 {
			return true, path
		} else {
			return false, path
		}
	}

	for _, node := range g.Nodes[1:] {
		if g.hamiltonianCycleAuxilaryCheck(index, node.Label, path, matrix) {
			path[index] = node.Label
		}
		if bool_ans, _ := g.hamiltonianCycleAuxilary(index+1, path, matrix); bool_ans {
			return true, path
		}
		path[index] = ""
	}
	return false, nil
}

func (g *Graph) HamiltonianCycle() (bool, []string) {
	matrix := g.NodeLabelAdjacentMatrix()
	nodes_count := len(g.Nodes)
	path := make([]string, nodes_count, nodes_count)
	path[0] = g.Nodes[0].Label
	ans_is_ham, path_ans := g.hamiltonianCycleAuxilary(1, path, matrix)
	return ans_is_ham, path_ans
}

// !!!!!!!
// !!! Точно ли здесь используется вес у ребра, узнать, как хранится вес на фронте .. map[string]int
func (g *Graph) MinimalSpanningTree() ([]Edge, int) {
	// 	inf := math.MaxInt
	// 	used := make(map[string]bool)
	// 	min_edge := make(map[string]int)
	// 	best_edge := make(map[string]string)
	// 	for _, node := range g.Nodes {
	// 		min_edge[node.Label] = inf
	// 		best_edge[node.Label] = ""
	// 		used[node.Label] = false
	// 	}
	// 	min_edge[g.Nodes[0].Label] = 0
	// 	//for _, node := range g.Nodes {
	// 	for i := 0; i < len(g.Nodes); i++ {
	// 		v := ""
	// 		for _, node1 := range g.Nodes {
	// 			if !used[node1.Label] && (v == "" || min_edge[node1.Label] < min_edge[v]) {
	// 				v = node1.Label
	// 			}
	// 		}
	// 		if min_edge[v] == inf {
	// 			return min_edge, false
	// 		}
	// 		used[v] = true

	// 		for _, node2 := range g.Nodes {
	// 			found, edge := g.FindEdge(v, node2.Label)
	// 			if found && edge.Weight < min_edge[node2.Label] {
	// 				min_edge[node2.Label] = edge.Weight
	// 				best_edge[node2.Label] = v
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(best_edge)
	// 	return min_edge, true
	// }

	// func (g *Graph) PrimMST() ([]Edge, int) {
	n := len(g.Nodes)
	if n == 0 {
		return nil, 0
	}

	// Множество посещённых вершин
	visited := make(map[string]bool)

	// MST ребра
	mstEdges := make([]Edge, 0, n-1)
	totalWeight := 0

	// Начинаем с первой вершины
	start := g.Nodes[0]
	visited[start.Label] = true

	// Мин-купа для ребер, где key — вес, value — ребро
	type EdgeWithWeight struct {
		Edge   Edge
		Weight int
	}
	edgeHeap := make([]EdgeWithWeight, 0)

	// Функция для добавления ребер, исходящих из вершины v, в кучу
	addEdges := func(v Node) {
		for _, e := range g.Edges {
			if visited[e.Source.Label] && !visited[e.Target.Label] {
				edgeHeap = append(edgeHeap, EdgeWithWeight{Edge: e, Weight: e.Weight})
			} else if visited[e.Target.Label] && !visited[e.Source.Label] {
				edgeHeap = append(edgeHeap, EdgeWithWeight{Edge: e, Weight: e.Weight})
			}
		}
	}

	addEdges(start)

	// Функция для вытаскивания ребра с минимальным весом из кучи
	popMin := func() (EdgeWithWeight, bool) {
		if len(edgeHeap) == 0 {
			return EdgeWithWeight{}, false
		}
		minIndex := 0
		for i := 1; i < len(edgeHeap); i++ {
			if edgeHeap[i].Weight < edgeHeap[minIndex].Weight {
				minIndex = i
			}
		}
		minEdge := edgeHeap[minIndex]
		edgeHeap = append(edgeHeap[:minIndex], edgeHeap[minIndex+1:]...)
		return minEdge, true
	}

	// Основной цикл
	for len(mstEdges) < n-1 {
		minEdge, ok := popMin()
		if !ok {
			// Граф не связный, MST не существует
			break
		}
		e := minEdge.Edge

		// Определяем новую вершину, которую добавим
		if visited[e.Source.Label] && !visited[e.Target.Label] {
			visited[e.Target.Label] = true
			mstEdges = append(mstEdges, e)
			totalWeight += e.Weight
			addEdges(e.Target)
		} else if visited[e.Target.Label] && !visited[e.Source.Label] {
			visited[e.Source.Label] = true
			mstEdges = append(mstEdges, e)
			totalWeight += e.Weight
			addEdges(e.Source)
		}
		// Если обе вершины уже посещены — пропускаем ребро
	}

	return mstEdges, totalWeight
}

func MakeGraphFromAdjLabelMatrix(matrix map[string]map[string]int) *Graph {
	new_graph := new(Graph)
	id := 1
	node_map := make(map[string]int)
	for node_label := range matrix {
		new_graph.AddNodeByInfo(id, node_label, "", 0, 0, 0)
		node_map[node_label] = id
		id++
	}
	edge_id := 0
	for node_src_l, node_list := range matrix {
		for node_trg_l, val := range node_list {
			if val != 0 {
				if ans, _ := new_graph.FindEdge(node_src_l, node_trg_l); !ans {
					new_graph.AddEdgeByInfo(
						Node{Id: node_map[node_src_l], Label: node_src_l},
						Node{Id: node_map[node_trg_l], Label: node_trg_l},
						strconv.Itoa(edge_id), "", "", 0)
					edge_id++
				}
			}
		}
	}
	return new_graph
}

func (g *Graph) FindDominatingSets() [][]string {
	n := len(g.Nodes)
	labelToIdx := make(map[string]int)
	for i, node := range g.Nodes {
		labelToIdx[node.Label] = i
	}

	// Матрица смежности для быстрого поиска соседей
	adjMatrix := make([][]bool, n)
	for i := range adjMatrix {
		adjMatrix[i] = make([]bool, n)
	}
	for _, edge := range g.Edges {
		u := labelToIdx[edge.Source.Label]
		v := labelToIdx[edge.Target.Label]
		adjMatrix[u][v] = true
		adjMatrix[v][u] = true
	}

	var results [][]string
	var currentSet []int
	covered := make([]bool, n) // покрытые вершины

	var backtrack func(start int)
	backtrack = func(start int) {
		// Проверяем, покрыты ли все вершины (внешняя устойчивость)
		allCovered := true
		for i := 0; i < n; i++ {
			if !covered[i] {
				allCovered = false
				break
			}
		}
		if allCovered {
			setLabels := make([]string, len(currentSet))
			for i, idx := range currentSet {
				setLabels[i] = g.Nodes[idx].Label
			}
			results = append(results, setLabels)
			return
		}

		for i := start; i < n; i++ {
			// Добавляем вершину i
			currentSet = append(currentSet, i)

			// Сохраняем старое состояние covered для отката
			oldCovered := make([]bool, n)
			copy(oldCovered, covered)

			// Помечаем i и соседей i как покрытые
			covered[i] = true
			for j := 0; j < n; j++ {
				if adjMatrix[i][j] {
					covered[j] = true
				}
			}

			backtrack(i + 1)

			// Откатываем изменения
			covered = oldCovered
			currentSet = currentSet[:len(currentSet)-1]
		}
	}

	backtrack(0)
	return results
}

func (g *Graph) FindIndependentSets() [][]string {
	n := len(g.Nodes)
	labelToIdx := make(map[string]int)
	for i, node := range g.Nodes {
		labelToIdx[node.Label] = i
	}

	// Матрица смежности для быстрого поиска соседей
	adjMatrix := make([][]bool, n)
	for i := range adjMatrix {
		adjMatrix[i] = make([]bool, n)
	}
	for _, edge := range g.Edges {
		u := labelToIdx[edge.Source.Label]
		v := labelToIdx[edge.Target.Label]
		adjMatrix[u][v] = true
		adjMatrix[v][u] = true
	}

	var results [][]string
	var currentSet []int

	var backtrack func(start int)
	backtrack = func(start int) {
		// Добавляем текущее множество в результаты
		setLabels := make([]string, len(currentSet))
		for i, idx := range currentSet {
			setLabels[i] = g.Nodes[idx].Label
		}
		results = append(results, setLabels)

		for i := start; i < n; i++ {
			conflict := false
			for _, selectedIdx := range currentSet {
				if adjMatrix[i][selectedIdx] {
					conflict = true
					break
				}
			}
			if !conflict {
				currentSet = append(currentSet, i)
				backtrack(i + 1)
				currentSet = currentSet[:len(currentSet)-1]
			}
		}
	}

	backtrack(0)
	return results
}

// func (g *Graph) MinCutUnweighted() (minCutWeight int, cutPartition map[string]bool) {
// 	n := len(g.Nodes)
// 	if n == 0 {
// 		return 0, nil
// 	}

// 	labelToIdx := make(map[string]int)
// 	idxToLabel := make([]string, n)
// 	for i, node := range g.Nodes {
// 		labelToIdx[node.Label] = i
// 		idxToLabel[i] = node.Label
// 	}

// 	// Строим матрицу смежности, вес 1 для существующего ребра, 0 - иначе
// 	weight := make([][]int, n)
// 	for i := range weight {
// 		weight[i] = make([]int, n)
// 	}
// 	for _, edge := range g.Edges {
// 		u := labelToIdx[edge.Source.Label]
// 		v := labelToIdx[edge.Target.Label]
// 		weight[u][v] = 1
// 		weight[v][u] = 1
// 	}

// 	vertices := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		vertices[i] = i
// 	}

// 	minCutWeight = int(^uint(0) >> 1) // max int
// 	var bestCut []bool

// 	for len(vertices) > 1 {
// 		nVertices := len(vertices)
// 		selected := make([]bool, nVertices)
// 		weights := make([]int, nVertices)
// 		selected[0] = true

// 		for i := 1; i < nVertices; i++ {
// 			weights[i] = weight[vertices[0]][vertices[i]]
// 		}

// 		var prev int
// 		var last int
// 		for i := 1; i < nVertices; i++ {
// 			maxW := -1
// 			next := -1
// 			for j := 1; j < nVertices; j++ {
// 				if !selected[j] && weights[j] > maxW {
// 					maxW = weights[j]
// 					next = j
// 				}
// 			}
// 			selected[next] = true
// 			last = next

// 			for j := 1; j < nVertices; j++ {
// 				if !selected[j] {
// 					weights[j] += weight[vertices[next]][vertices[j]]
// 				}
// 			}
// 			prev = next
// 		}

// 		if weights[last] < minCutWeight {
// 			minCutWeight = weights[last]
// 			bestCut = make([]bool, n)
// 			for i, v := range vertices {
// 				bestCut[v] = selected[i]
// 			}
// 		}

// 		vLast := vertices[last]
// 		vEnd := vertices[nVertices-1]

// 		for i := 0; i < nVertices; i++ {
// 			weight[vEnd][vertices[i]] += weight[vLast][vertices[i]]
// 			weight[vertices[i]][vEnd] = weight[vEnd][vertices[i]]
// 		}

// 		vertices = append(vertices[:last], vertices[last+1:]...)
// 	}

// 	cutPartition = make(map[string]bool)
// 	for i, inSet := range bestCut {
// 		cutPartition[idxToLabel[i]] = inSet
// 	}

// 	return minCutWeight, cutPartition
// }
