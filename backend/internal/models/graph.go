package models

import (
	"errors"
	"fmt"
	"math"
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
	Id     int
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

func (g *Graph) IsEdgeById(id int) bool {
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

func (g *Graph) FindEdge(source, target string) bool {
	for _, edge := range g.Edges {
		if g.oriented {
			if source == edge.Source.Label && target == edge.Target.Label {
				return true
			}
		} else {
			if source == edge.Source.Label && target == edge.Target.Label || source == edge.Target.Label && target == edge.Source.Label {
				return true
			}
		}
	}
	return false
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

func (g *Graph) AddEdgeByInfo(source Node, target Node, id int, label string, color string, weight int) error {
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
			if (edge1.Source.Id == edge2.Source.Id || edge1.Source.Id == edge2.Target.Id || edge1.Target.Id == edge2.Source.Id || edge1.Target.Id == edge2.Target.Id) && (edge1.Id != edge2.Id) {
				matrix[edge1.Label][edge2.Label] = 1
				matrix[edge2.Label][edge1.Label] = 1
			}
		}
	}
	return matrix
}

// Матрица смежности ребер, но вместо индексов - Id
func (g *Graph) EdgeIdAdjacentMatrix() map[int]map[int]int {
	matrix := make(map[int]map[int]int)
	for _, edge1 := range g.Edges {
		matrix[edge1.Id] = make(map[int]int)
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
	edges_g := make(map[int]Edge)
	edges_final := make(map[int]Edge)
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
	for _, node := range nodes_g_set {
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
		edges_set[edge.Label] = edge
	}
	for _, edge := range graph.Edges {
		edges_set[edge.Label] = edge
	}
	answer := new(Graph)
	for _, node := range nodes_set {
		answer.AddNode(node)
	}
	for _, edge := range edges_set {
		answer.AddEdge(edge)
	}
	return answer
}

func (g *Graph) Join(graph *Graph) *Graph {
	unioned_graphs := g.Union(graph)
	nodes_g := make(map[string]struct{})
	flag := false
	for _, node := range g.Nodes {
		for _, node1 := range graph.Nodes {
			if node.Label == node1.Label {
				flag = true
			}
		}
		if !flag {
			nodes_g[node.Label] = struct{}{}
		}
	}
	nodes_graph := make(map[string]struct{})
	for _, node := range graph.Nodes {
		for _, node1 := range g.Nodes {
			if node.Label == node1.Label {
				flag = true
			}
		}
		if !flag {
			nodes_graph[node.Label] = struct{}{}
		}
	}
	for node_src_l := range nodes_g {
		for node_trg_l := range nodes_graph {
			node_src, _ := g.FindNodeByLabel(node_src_l)
			node_trg, _ := g.FindNodeByLabel(node_trg_l)
			unioned_graphs.AddEdgeByInfo(node_src, node_trg, 0, "", "", 0)
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
	/////////////////

	return answer
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
	for node_src_l, node_list := range matrix {
		for node_trg_l, val := range node_list {
			if val != 0 {
				if !new_graph.FindEdge(node_src_l, node_trg_l) {
					new_graph.AddEdgeByInfo(
						Node{Id: node_map[node_src_l], Label: node_src_l},
						Node{Id: node_map[node_trg_l], Label: node_trg_l},
						id, "", "", 0)
					id++
				}
			}
		}
	}
	return new_graph
}
